package task

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/coreos/etcd/clientv3"
)

type Job struct {
	Name 		string 		`json:"name"`				// 任务名
	Command 	string 		`json:"command"`			// 任务命令
	Expr 		string 		`json:"expr"`				// 任务cron
}

// 增
func (j *Job) Save() (oldJob *Job, err error){

	jobKey := GC.JOB_SAVE_DIR + j.Name
	jobVal, err := json.Marshal(j)
	if(err != nil){
		fmt.Println("marshal err：",err)
		return
	}

	putRes, err := GM.KV.Put(context.TODO(), jobKey, string(jobVal), clientv3.WithPrevKV())
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
func (j *Job) Delete()  (oldJob *Job, err error) {

	jobKey := GC.JOB_SAVE_DIR + j.Name
	delRes, err := GM.KV.Delete(context.TODO(), jobKey)
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
func (j *Job) Update() (oldJob *Job, err error){
	return j.Save()
}

// 查：列表
func (j *Job) List()  (js []*Job, err error) {

	jobKey := GC.JOB_SAVE_DIR + j.Name
	getRes, err := GM.KV.Get(context.TODO(), jobKey, clientv3.WithPrefix())
	if(err != nil){
		fmt.Println("etcd get err:", err)
		return
	}

	js = make([]*Job, 0)

	for _, kvPair := range getRes.Kvs {
		job := &Job{}
		_ = json.Unmarshal(kvPair.Value, job)
		js = append(js, job)
	}
	return
}

// 过期
func (j *Job) Kill()  (err error) {

	killkEY := GC.JOB_KILL_DIR + j.Name

	// 业务代码：设置一个租约，通知 worker 杀死任务
	grantRes, err := GM.Lease.Grant(context.TODO(), 1)
	if(err != nil){
		fmt.Println("etcd grant err:", err)
		return
	}

	_, err =  GM.KV.Put(context.TODO(), killkEY, "", clientv3.WithLease( grantRes.ID))
	if(err != nil){
		fmt.Println("etcd put err:", err)
	}

	return
}
