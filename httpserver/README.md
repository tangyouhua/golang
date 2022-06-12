# Homework

## 练习8-1：编写 httpserver 部署脚本

### 执行部署

- 创建 namespace

```shell
kubectl create -f httpserver-namespace.yaml
kubectl get namespaces

root@k8smaster:~/httpserver# kubectl get namespaces
NAME               STATUS   AGE
calico-apiserver   Active   25h
calico-system      Active   25h
default            Active   25h
httpserver         Active   9m50s
kube-node-lease    Active   25h
kube-public        Active   25h
kube-system        Active   25h
nginx              Active   37m
tigera-operator    Active   25h
```

- 创建 Deployment

```shell
kubectl create -f httpserver-deployment.yaml 
kubectl get pods -n httpserver

NAME                                      READY   STATUS    RESTARTS   AGE
httpserver-deployment1-66d67bb6bf-ncmhh   1/1     Running   0          9m8s
```

- 创建 service

```shell
kubectl create -f httpserver-service.yaml 
kubectl get services -n httpserver

NAME                     TYPE       CLUSTER-IP     EXTERNAL-IP   PORT(S)          AGE
httpserver-deployment1   NodePort   10.107.179.2   <none>        9001:31091/TCP   8m5s
```

查看 Service 信息

```shell
kubectl describe service httpserver-deployment1 -n httpserver

Name:                     httpserver-deployment1
Namespace:                httpserver
Labels:                   app=httpserver
Annotations:              <none>
Selector:                 app=httpserver
Type:                     NodePort
IP Family Policy:         SingleStack
IP Families:              IPv4
IP:                       10.107.179.2
IPs:                      10.107.179.2
Port:                     httpserver-service80  9001/TCP
TargetPort:               80/TCP
NodePort:                 httpserver-service80  31091/TCP
Endpoints:                192.168.249.44:80
Session Affinity:         None
External Traffic Policy:  Cluster
Events:                   <none>
```

### 验证结果

```shell
curl http://192.168.38.128:31091/healthz
HTTP server is working.
```

**思考点**

- 优雅启动：对启动有依赖项或者启动条件要求时，可采用 [PostStart Container Hook][1] 或者配置 [Init Container][2] 提供支持。httpserver 服务暂时无此需求；
- 优雅终止
- 资源需求和 QoS 保证
- 探活：通过 [liveness HTTP request][x] 对 httpServer `/healthz` 接口探活；
- 日常运维需求，日志等级
- 配置和代码分离

[1]: https://kubernetes.io/docs/concepts/containers/container-lifecycle-hooks/
[2]: https://kubernetes.io/docs/tasks/configure-pod-container/configure-pod-initialization/
[x]: https://kubernetes.io/docs/tasks/configure-pod-container/configure-liveness-readiness-startup-probes/

## 练习1, 2：编写 HTTP 服务器，制作镜像

### 要求

1. 接收客户端 request，并将 request 中带的 header 写入 response header
2. 读取当前系统的环境变量中的 VERSION 配置，并写入 response header
3. Server 端记录访问日志包括客户端 IP，HTTP 返回码，输出到 server 端的标准输出
4. 当访问 localhost/healthz 时，应返回 200
5. 构建本地镜像
6. 推送镜像到 Docker Hub
7. 通过 docker 本地命令启动 httpserver
8. 通过 nsenter 进入容器查看 IP 地址

### 运行

```shell
# 宿主机
$ docker run -p 8080:80 -d tangyouhua/httpserver:v1.0 # 映射 httpserver 到本地 8080 端口
$ curl localhost:8080/healthz #查看服务
HTTP server is working.
$ docker inspect <container id> -f '{{.State.Pid}}'
9337
$ sudo nsenter -n -t9937
root# ip addr
1: lo: <LOOPBACK,UP,LOWER_UP> mtu 65536 qdisc noqueue state UNKNOWN group default qlen 1000
    link/loopback 00:00:00:00:00:00 brd 00:00:00:00:00:00
    inet 127.0.0.1/8 scope host lo
       valid_lft forever preferred_lft forever
46: eth0@if47: <BROADCAST,MULTICAST,UP,LOWER_UP> mtu 1500 qdisc noqueue state UP group default
    link/ether 02:42:ac:11:00:02 brd ff:ff:ff:ff:ff:ff link-netnsid 0
    inet 172.17.0.2/16 brd 172.17.255.255 scope global eth0
       valid_lft forever preferred_lft forever
```

## 镜像

- golang 1.17 镜像：https://hub.docker.com/_/golang/
- busybox 1.34 镜像：<https://hub.docker.com/_/busybox>
- tangyouhua/httpserver:v1.0 镜像：<https://hub.docker.com/repository/docker/tangyouhua/httpserver>
