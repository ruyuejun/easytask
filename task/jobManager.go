package task

import (
	"context"
	"dcs-gocron/common"
	"dcs-gocron/config"
	"encoding/json"
	"fmt"
	"github.com/coreos/etcd/clientv3"
	"github.com/coreos/etcd/mvcc/mvccpb"
	"time"
)

type JobManager struct {
	client *clientv3.Client
	kv clientv3.KV
	lease clientv3.Lease
	watcher clientv3.Watcher
}

// 单例job管理对象
var GJobManager *JobManager

// 创建job管理对象
func NewJobManager() (err error) {

	conf := clientv3.Config{
		Endpoints: config.GConfig.EtcdEndpoints,		// 集群地址
		DialTimeout : 5000 * time.Millisecond,			// 连接超时时间
	}

	// 建立etcd连接
	client, err := clientv3.New(conf)
	if err != nil {
		return
	}

	// 得到KV和Lease的API
	kv := clientv3.NewKV(client)
	lease := clientv3.NewLease(client)
	watcher := clientv3.NewWatcher(client)

	// 赋值为单例
	GJobManager = &JobManager{
		client: client,
		kv:     kv,
		lease:  lease,
		watcher: watcher,
	}
	return
}

// 保存任务
func (jobM *JobManager)SaveJob(job *Job) (oldJob *Job, err error){

	key := common.JOBDIR + job.Name
	value, err  := json.Marshal(job)
	if err != nil {
		return
	}

	fmt.Println("存储任务：", key)

	putResp, err := jobM.kv.Put(context.TODO(), key, string(value), clientv3.WithPrevKV())
	if err != nil {
		return
	}

	fmt.Println("存储的prev：", putResp)

	// 如果是更新，返回旧值
	if putResp.PrevKv != nil {

		err = json.Unmarshal(putResp.PrevKv.Value, &oldJob)
		if err != nil {
			fmt.Println("反序列化出错：", err)
			err = nil
		}

	}

	return

}

// 删除任务
func (jobM *JobManager)DeleteJob(job *Job) (oldJob *Job, err error){

	key := common.JOBDIR + job.Name

	fmt.Println("删除任务：", key)

	delResp, err := jobM.kv.Delete(context.TODO(), key, clientv3.WithPrevKV())
	if err != nil {
		return
	}

	// 返回被删除的任务信息
	if len(delResp.PrevKvs) != 0 {
		err = json.Unmarshal(delResp.PrevKvs[0].Value, &oldJob)
		if err != nil {
			fmt.Println("反序列化出错：", err)
			err = nil
		}
	}

	return

}

// 杀死任务：修改任务名，则worker就能收到通知，停止任务
func (jobM *JobManager)KillJob(job *Job) (oldJob *Job, err error){

	key := common.JOBDIR + job.Name
	killKey := common.KILLDIR + job.Name

	fmt.Println("修改任务：", key)

	// 让worker监听到put操作，kill自动过期
	leaseGrantResp, err := jobM.lease.Grant(context.TODO(), 1)
	if err != nil {
		return
	}
	_, err = jobM.kv.Put(context.TODO(), killKey, "", clientv3.WithLease(leaseGrantResp.ID))
	if err != nil {
		return
	}

	return

}

// 查询任务
func (jobM *JobManager)ListJobs() (jobList []*Job, err error){

	getResp, err := jobM.kv.Get(context.TODO(), common.JOBDIR, clientv3.WithPrefix())
	if err != nil {
		return
	}

	jobList = make([]*Job, 0)

	for _, kvPair := range getResp.Kvs {
		job := &Job{}
		err = json.Unmarshal(kvPair.Value,  job)
		if err != nil {
			fmt.Println("反序列化失败", err)
			err = nil
		}
		jobList = append(jobList, job)
	}

	return

}

// 监听任务
func (jobM *JobManager)WatchJobs() (err error){

	// 获取/cron/jobs/目录下所有任务，以及当前集群的revision
	getResp, err := jobM.kv.Get(context.TODO(), common.JOBDIR, clientv3.WithPrefix())
	if err != nil {
		return
	}

	// 遍历出当前有哪些任务
	for _, kvPair := range getResp.Kvs {

		var job *Job

		err = json.Unmarshal(kvPair.Value, &job)
		if err != nil {
			break
		}
		// 生成任务事件
		jobEvent := NewJobEvent(JOB_EVENT_SAVE, job)
		// 将该job推送给scheduler调度协程
		GScheduler.Push(jobEvent)

	}

	// 监听协程：从该revision向后监听变化事件
	go func() {

		watchStartRevision := getResp.Header.Revision + 1
		watchChan := jobM.watcher.Watch(context.TODO(), common.JOBDIR, clientv3.WithRev(watchStartRevision), clientv3.WithPrefix())

		// 处理监听事件
		for watchResp := range watchChan {

			for _,watchEvent := range watchResp.Events {

				var job *Job

				switch watchEvent.Type {

				case mvccpb.PUT:											// 生成更新事件
					err := json.Unmarshal(watchEvent.Kv.Value, &job)
					if err != nil {
						fmt.Println("序列化失败：", err)
						err = nil
						continue
					}
					jobEvent := NewJobEvent(JOB_EVENT_SAVE, job)
					// 推送给调度协程scheduler
					GScheduler.Push(jobEvent)

				case mvccpb.DELETE:											// 生成删除事件
					jobName := job.ExtractRealName()
					job = &Job{Name: jobName}
					jobEvent := NewJobEvent(JOB_EVENT_DELET, job)
					// 推送给调度协程scheduler
					GScheduler.Push(jobEvent)
				}

			}

		}

	}()

	return
}

