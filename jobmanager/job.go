package jobmanager

import (
	"Demo1/config"
	"github.com/coreos/etcd/clientv3"
	"time"
)

type JobManager struct {
	client *clientv3.Client
	kv clientv3.KV
	lease clientv3.Lease
}

var (
	// 单例
	GJobManager *JobManager
)

type Job struct {
	Name string
	Command string
	CronExpr string
}

func Init() (err error) {

	config := clientv3.Config{
		Endpoints: config.GConfig.EtcdEndpoints,		// 集群地址
		DialTimeout : 5000 * time.Millisecond,			// 连接超时时间
	}

	// 建立etcd连接
	client, err := clientv3.New(config)
	if err != nil {
		return
	}

	// 得到KV和Lease的API
	kv := clientv3.NewKV(client)
	lease := clientv3.NewLease(client)

	// 赋值为单例
	GJobManager = &JobManager{
		client: client,
		kv:     kv,
		lease:  lease,
	}
	return
}
