package jobmanager

import (
	"Demo1/common"
	"context"
	"encoding/json"
	"fmt"
	"github.com/coreos/etcd/clientv3"
)

// 监听任务
func (jobM *JobManager)WatchJobs(job *Job) (err error){

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

