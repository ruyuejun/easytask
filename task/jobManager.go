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
}


// 监视
func (manager *Manager) Watch() {

	var watchChan clientv3.WatchChan
	var watchRes clientv3.WatchResponse
	var watchEvent *clientv3.Event
	var job *Job
	var jobEvent *JobEvent

	// 首先：系统启动时要推送所有任务给调度协程
	getRes, err := GM.KV.Get(context.TODO(), GC.JOB_SAVE_DIR, clientv3.WithPrefix())
	for _, kvpair := range getRes.Kvs {
		job := &Job{}
		err = json.Unmarshal(kvpair.Value, job)
		if err != nil {
			fmt.Println("Unmarshal warn:", err)
		}

		// 构建任务
		jobEvent = NewJobEvent(GC.JOB_EVENT_SAVE, job)
		fmt.Println("*jobEvent:", *jobEvent)

		// 推送任务给调度协程
		GScheduler.PushJobEvent(jobEvent)
	}

	// 其次：系统启动后要从下一次 revision 开始向后监听变化事件
	go func(){

		// 启动监听，监听 任务目录后续变化
		watchStartRevision := getRes.Header.Revision + 1
		watchChan = GM.Watcher.Watch(context.TODO(), GC.JOB_SAVE_DIR, clientv3.WithRev(watchStartRevision), clientv3.WithPrefix())
		
		for watchRes = range watchChan {
			for _, watchEvent = range watchRes.Events {

				// 构建任务
				switch watchEvent.Type {
				case mvccpb.PUT:		// 构建更新任务事件
					job = &Job{}
					err = json.Unmarshal(watchEvent.Kv.Value, job)
					if err != nil {
						fmt.Println("Unmarshal warn:", err)
						continue
					}
					jobEvent = NewJobEvent(GC.JOB_EVENT_SAVE, job)

				case mvccpb.DELETE:	 	// 构建删除任务事件
					jobName := strings.TrimPrefix(string(watchEvent.Kv.Key), GC.JOB_SAVE_DIR)
					job = &Job{
						Name: jobName,
					}
					jobEvent = NewJobEvent(GC.JOB_EVENT_DELETE, job)
				}
				fmt.Println("推送的事件是：", *jobEvent)

				// 推送任务事件给调度协程 Scheduler
				GScheduler.PushJobEvent(jobEvent)
			}
		}
	}()

	return
}
