FROM golang:1.17 AS build
WORKDIR /httpserver/
COPY . .
ENV CGO_ENABLED=0
ENV GO111MODULE=on
ENV GOPROXY=https://goproxy.cn,direct
RUN GOOS=linux go build -installsuffix cgo -o httpserver main.go

FROM busybox:1.34
COPY --from=build /httpserver/httpserver /httpserver/httpserver
EXPOSE 80
WORKDIR /httpserver/
ENTRYPOINT ["./httpserver"]
