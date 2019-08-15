package jobmanager

import (
	"Demo1/common"
	"context"
	"encoding/json"
	"fmt"
	"github.com/coreos/etcd/clientv3"
)

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
