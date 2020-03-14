# Truss

记录时间：2020-03-14

## 简介

`truss`是一款`go-kit`代码生成工具, 可以快速编写`go-kit`中繁杂的代码, 并生成支持`http`和`grpc`两种调用方式的服务

地址：[https://github.com/metaverse/truss.git](https://github.com/metaverse/truss.git) 

## 安装

1. 首先安装protoc，参考`gk-grpc`里的`README.md`文件
2. 下载truss源码，`go get -u -d github.com/metaverse/truss`，注意不要在go module接管的目录下执行(`export GO111MODULE=off`)，要确保`truss`安装在`$GOPATH/src下`，而不是`$GOPATH/pkg/mod`下
3. 进入`$GOPATH/src/github.com/metaverse/truss`目录下执行：`make dependencies && make`

```
-u参数: 利用网络来更新已有的代码包及其依赖包 
-d参数: 表示只执行下载动作，不执行安装动作
```

## 自动生成代码

```html
.
├── echo-service
│   ├── cmd
│   │   └── echo
│   │       └── main.go
│   ├── handlers
│   │   ├── handlers.go
│   │   ├── hooks.go
│   │   └── middlewares.go
│   └── svc
│       ├── client
│       │   ├── grpc
│       │   │   └── client.go
│       │   └── http
│       │       └── client.go
│       ├── endpoints.go
│       ├── server
│       │   └── run.go
│       ├── transport_grpc.go
│       └── transport_http.go
├── echo.pb.go
└── echo.proto
```
