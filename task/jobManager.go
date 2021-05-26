package task

import (
	"context"
	"easycron/common"
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/coreos/etcd/clientv3"
	"github.com/coreos/etcd/mvcc/mvccpb"
)

// 任务管理器
type Manager struct {
	Client 	*clientv3.Client
	KV 		clientv3.KV
	Lease 	clientv3.Lease
	Watcher clientv3.Watcher
}

var GM *Manager


func InitManager() {

	etcdConfig := clientv3.Config{
		Endpoints:   common.GConfig.WorkerEndPonts,
		DialTimeout: time.Duration(common.GConfig.WorkerDialTimeout) * time.Millisecond,
	}

	etcdClient, err := clientv3.New(etcdConfig)
	if err != nil {
		fmt.Println("etcd客户端连接异常")
		panic(err)
	}

	kv := clientv3.NewKV(etcdClient)
	lease := clientv3.NewLease(etcdClient)
	watcher := clientv3.NewWatcher(etcdClient)

	GM = &Manager{
		etcdClient,
		kv,
		lease,
		watcher,
	}

	// 派发任务
	GM.Dispatch()
}

// 派发：系统首次启动要派发所有任务给调度协程
func (manager *Manager) Dispatch() {

	var jobEvent *JobEvent

	getRes, err := GM.KV.Get(context.TODO(), GC.JOB_SAVE_DIR, clientv3.WithPrefix())
	if err != nil {
		fmt.Println("watch err:", err)
		return
	}
	for _, kvpair := range getRes.Kvs {
		dispatchJob := &Job{}
		err = json.Unmarshal(kvpair.Value, dispatchJob)
		if err != nil {
			fmt.Println("Unmarshal warn:", err)
			return
		}
		// 把该任务同步给调度协程执行调度
		jobEvent = NewJobEvent(GC.JOB_EVENT_DELETE, dispatchJob)
		fmt.Println("派发任务：", jobEvent.EventJob.Name)
		GScheduler.PushJobEvent(jobEvent)
	}

	// 监视后续任务变化
	manager.watch(getRes)
}

// 监视：系统启动后要从下一次 revision 开始向后监听变化事件
func (manager *Manager) watch(getRes *clientv3.GetResponse) {
	var jobEvent *JobEvent
	go func(){
		// 启动监听，监听任务目录后续变化
		watchStartRevision := getRes.Header.Revision + 1
		watchChan := GM.Watcher.Watch(context.TODO(), GC.JOB_SAVE_DIR, clientv3.WithRev(watchStartRevision), clientv3.WithPrefix())
		
		for watchRes := range watchChan {
			for _, watchEvent := range watchRes.Events {

				// 构建任务
				switch watchEvent.Type {

				case mvccpb.PUT:		// 构建更新任务事件：让调度协程更新该任务信息，下次按照新信息执行
					putJob := &Job{}
					err := json.Unmarshal(watchEvent.Kv.Value, putJob)
					if err != nil {
						fmt.Println("Unmarshal warn:", err)
						continue
					}
					jobEvent = NewJobEvent(GC.JOB_EVENT_SAVE, putJob)

				case mvccpb.DELETE:	 	// 构建删除任务事件：让调度协程终止该任务
					delJob := &Job{
						Name: strings.TrimPrefix(string(watchEvent.Kv.Key), GC.JOB_SAVE_DIR),
					}
					jobEvent = NewJobEvent(GC.JOB_EVENT_DELETE, delJob)
				}

				// 推送任务事件给调度协程 Scheduler
				fmt.Println("监听任务：", jobEvent.EventJob.Name)
				GScheduler.PushJobEvent(jobEvent)
			}
		}
	}()
}