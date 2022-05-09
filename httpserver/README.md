# 练习：编写 HTTP 服务器，制作镜像

## 要求

1. 接收客户端 request，并将 request 中带的 header 写入 response header
2. 读取当前系统的环境变量中的 VERSION 配置，并写入 response header
3. Server 端记录访问日志包括客户端 IP，HTTP 返回码，输出到 server 端的标准输出
4. 当访问 localhost/healthz 时，应返回 200
5. 构建本地镜像
6. 推送镜像到 Docker Hub
7. 通过 docker 本地命令启动 httpserver
8. 通过 nsenter 进入容器查看 IP 地址

## 运行

```shell
# 宿主机
$ docker run -p 8080:80 -d tangyouhua/httpserver:v1.0 # 映射 httpserver 到本地 8080 端口
$ curl localhost:8080/healthz #查看服务
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

# 镜像

- golang 1.17 镜像：https://hub.docker.com/_/golang/
- busybox 1.34 镜像：<https://hub.docker.com/_/busybox>
- tangyouhua/httpserver:v1.0 镜像：<https://hub.docker.com/repository/docker/tangyouhua/httpserver>
