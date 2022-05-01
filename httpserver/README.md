# 练习：编写 HTTP 服务器，制作镜像

## 要求

1. 接收客户端 request，并将 request 中带的 header 写入 response header
2. 读取当前系统的环境变量中的 VERSION 配置，并写入 response header
3. Server 端记录访问日志包括客户端 IP，HTTP 返回码，输出到 server 端的标准输出
4. 当访问 localhost/healthz 时，应返回200
5. 构建本地镜像
6. 推送镜像到 Docker Hub
7. 通过 docker 本地命令启动 httpserver
8. 通过 nsenter 进入容器查看 IP 地址

## 运行

```shell
# 宿主机
docker run -p 80:80 -d tangyouhua/httpserver:v1.0 # 映射 httpserver 到本地 80 端口
curl localhost:/healthz #查看服务
docker cp /usr/sbin/ifconfig <contaienrid>:/bin/ifconfig # 拷贝 ifconfig 到容器
PID=$(docker inspect --format "{{.State.Pid}}" <containerid>)
sudo nsenter --target $PID --mount --uts --ipc --net --pid
# 容器
# https://hub.docker.com/repository/docker/tangyouhua/httpserver
$ifconfig #查看 IP
```