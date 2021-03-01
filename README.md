# README

## 项目原因

配置多个机器的任务比较繁琐，且容易因为机器故障导致任务丢失，任务执行状态查询不便，市面上已有的系统如 quartz 非常复杂，不必要。

项目核心：

- 调度器：高可用
- 执行器：可扩展，以支持大量任务的并发执行

满足的 CAP 特性（一致性--最终一致性、可用性、分区容错性）：P 必须满足

应用架构一般不采用 BASE 理论，而是采用 BASE 理论(基本可用--保证整体可用、软状态--延迟同步、最终一致性)

## 项目思路

市面上的调度器大多依赖 rpc 通信，让不同的 worker 节点执行不同的节点。本项目每个 worker 都执行全量任务，这里执行任务较小，且单机下能支持 10 万并发的任务，所以低于该数目的任务可以使用，且减少了 rpc 步骤，提升了性能。

由于每个 workker 都是全量任务，就需要抢占 etcd 的分布式锁，以解决并发调度问题。

## 第一步 环境创建

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

一些错误：

```txt
错误一：Go1.13与1.14项目在启动时，由于go mod包版本原因会报如下错误：
          undefined: balancer.PickOptions
      解决：go.mod 添加：replace google.golang.org/grpc => google.golang.org/grpc v1.26.0

```

启动:

```txt
# 启动 master，这是默认启动方式，可以忽略 -s 参数
go run main.go -s master

# 启动 worker
go run main.go -s worker

# 启动 frontend，这是前端开发环境。之所以将前后端代码混合，是因为该项目极小
cd frontend
npm run serve
```
