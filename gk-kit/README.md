# kit

记录时间：2020-03-14

## 简介

`kit`是用来生成`go-kit`微服务代码的工具, 后面用来快速生成多个服务, 方便测试服务调用, 链路追踪等功能

## 常用命令

查看帮助文档：`kit help`

## 使用方法

### 创建一个hello的服务

新建hello服务：`kit new service hello`，会产生以下目录结构：

```html
.
└── hello
    └── pkg
        └── service
            └── service.go
```

此时`service.go`中只是定义了一个空的接口，并没有定义任何方法：

```go
// HelloService describes the service.
type HelloService interface {
	// Add your methods here
	// e.x: Foo(ctx context.Context,s string)(rs string, err error)
}
```

如果要继续创建`transport`、`endpoint`、`middleware`这些内容，则需要定义好接口里的方法，定义好之后，才能继续执行之后的命令，例如：

```go
type HelloService interface {
	Foo(ctx context.Context,s string)(rs string, err error)
}
```

### 根据定义好的接口方法，生成代码

生成其他代码: `kit g s hello`, 会产生以下目录结构(默认生成的transport是http):

```html
.
└── hello
    ├── cmd
    │   ├── main.go
    │   └── service
    │       ├── service.go
    │       └── service_gen.go
    └── pkg
        ├── endpoint
        │   ├── endpoint.go
        │   └── endpoint_gen.go
        ├── http
        │   ├── handler.go
        │   └── handler_gen.go
        └── service
            ├── middleware.go
            └── service.go
```

生成grpc: `kit g s hello -t grpc`: 

```html

```
