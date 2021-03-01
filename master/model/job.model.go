package model

import (
	"context"
	"easycron/task"
	"encoding/json"
	"fmt"

	"github.com/coreos/etcd/clientv3"
)

type JobModel task.Job

// 增
func (model *JobModel) Save() (oldJob *JobModel, err error){

	jobKey := task.GC.JOB_SAVE_DIR + model.Name
	jobVal, err := json.Marshal(model)
	if(err != nil){
		fmt.Println("marshal err：",err)
		return
	}

	putRes, err := task.GM.KV.Put(context.TODO(), jobKey, string(jobVal), clientv3.WithPrevKV())
	if(err != nil){
		fmt.Println("etcd put err:", err)
		return
	}

	// 检查put是更新还是全新插入
	if putRes.PrevKv != nil {
		_ = json.Unmarshal(putRes.PrevKv.Value, &oldJob)
	}
	return
}

// 删
func (model *JobModel) Delete()  (oldJob *JobModel, err error) {

	jobKey := task.GC.JOB_SAVE_DIR + model.Name
	delRes, err := task.GM.KV.Delete(context.TODO(), jobKey)
	if(err != nil){
		fmt.Println("etcd del err:", err)
		return
	}

	if len(delRes.PrevKvs) != 0 {
		_ = json.Unmarshal(delRes.PrevKvs[0].Value, &oldJob)
	}
	return
}

// 改
func (model *JobModel) Update() (oldJob *JobModel, err error){
	return model.Save()
}

// 查：列表
func (model *JobModel) List()  (models []*JobModel, err error) {

	jobKey := task.GC.JOB_SAVE_DIR + model.Name
	getRes, err := task.GM.KV.Get(context.TODO(), jobKey, clientv3.WithPrefix())
	if(err != nil){
		fmt.Println("etcd get err:", err)
		return
	}

	models = make([]*JobModel, 0)

	for _, kvPair := range getRes.Kvs {
		job := &JobModel{}
		_ = json.Unmarshal(kvPair.Value, job)
		models = append(models, job)
	}
	return
}

// 过期
func (model *JobModel) Kill()  (err error) {

	killkEY := task.GC.JOB_KILL_DIR + model.Name

	// 业务代码：设置一个租约，通知 worker 杀死任务
	grantRes, err := task.GM.Lease.Grant(context.TODO(), 1)
	if(err != nil){
		fmt.Println("etcd grant err:", err)
		return
	}

	_, err =  task.GM.KV.Put(context.TODO(), killkEY, "", clientv3.WithLease( grantRes.ID))
	if(err != nil){
		fmt.Println("etcd put err:", err)
	}

	return
}


