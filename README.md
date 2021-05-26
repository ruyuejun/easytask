# README

## 说明

多机配置同步 crontab 任务较为繁琐，查询任务执行列表不方便，且宕机会造成任务丢失，Quartz 太复杂了，只需要一个并发同步任务的调度器即可。

master：在 etcd 中，对任务进行简单的增删改查，提供 web 操作界面。

worker：从 etcd 同步任务到内存，并发调度、执行任务。

## 项目启动

### 创建环境

创建 docker 本地环境：

```txt
# 下载镜像
docker pull bitnami/etcd:latest

# 运行服务
docker run -d --name myetcd --publish 2379:2379 \
--env ALLOW_NONE_AUTHENTICATION=yes \
--env ETCD_ADVERTISE_CLIENT_URLS=http://myetcd:2379 \
bitnami/etcd:latest
```

贴士 ：Go1.13/1.14 项目在启动时，由于 go mod 包版本原因会报如下错误：`undefined: balancer.PickOptions`，可以在 go.mod 中添加：`replace google.golang.org/grpc => google.golang.org/grpc v1.26.0`。

### 启动

```txt
# 启动 master，这是默认启动方式，可以忽略 -s 参数
go run main.go -s master

# 启动 worker
go run main.go -s worker

# 启动 frontend开发环境
cd frontend
npm run serve
```

任务保存数据为：

```txt
    [
        {
            "name": "/job/demo1",
            "command": "echo 111",
            "expr": "*/1 * * * * * *"
        },
        {
            "name": "/job/demo2",
            "command": "echo 222",
            "expr": "*/2 * * * * * *"
        },
        {
            "name": "/job/demo3",
            "command": "echo 333",
            "expr": "*/5 * * * * * *"
        }
    ]
```
