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

生成grpc: `kit g s hello -t grpc`, 新增的内容有以下这些:  

```html
.
└── hello
    ├── cmd
    │   └── service
    │       ├── service.go
    │       └── service_gen.go
    └── pkg       
        ├── grpc
        │   ├── handler.go
        │   ├── handler_gen.go
        │   └── pb
        │       ├── compile.sh
        │       ├── hello.pb.go
        │       └── hello.proto       
        └── service
            ├── middleware.go
            └── service.go
```

> 当重新执行`kit g s hello`命令和添加endpoints后，以`_gen`结尾的文件会重新生成 

### 执行cmd/main.go

安装依赖: `go mod tidy`, 若出现下面的错误: 

```
go: finding module for package github.com/openzipkin/zipkin-go-opentracing
go: found github.com/openzipkin/zipkin-go-opentracing in github.com/openzipkin/zipkin-go-opentracing v0.4.5
go: gk-kit/hello/cmd/service imports
	github.com/openzipkin/zipkin-go-opentracing: github.com/openzipkin/zipkin-go-opentracing@v0.4.5: parsing go.mod:
	module declares its path as: github.com/openzipkin-contrib/zipkin-go-opentracing
	        but was required as: github.com/openzipkin/zipkin-go-opentracing
```

用`github.com/openzipkin-contrib/zipkin-go-opentracing`来替换`github.com/openzipkin/zipkin-go-opentracing`即可.

由于`zipkin-go-opentracing`包升级的原因, `cm/service/service.go`文件的内容需要做一些调整：

```go
import (
    opentracinggo "github.com/opentracing/opentracing-go"
	"github.com/openzipkin/zipkin-go"
	zipkinhttp "github.com/openzipkin/zipkin-go/reporter/http"
	zipkingoopentracing "github.com/openzipkin-contrib/zipkin-go-opentracing"
)

func Run() {
	if *zipkinURL != "" {
		logger.Log("tracer", "Zipkin", "URL", *zipkinURL)
		reporter := zipkinhttp.NewReporter(*zipkinURL)
		defer reporter.Close()
		// create our local service endpoint
		endpoint, err := zipkin.NewEndpoint("hello", "localhost:80")
		if err != nil {
			logger.Log("err", err)
			os.Exit(1)
		}
		nativeTracer, err := zipkin.NewTracer(reporter, zipkin.WithLocalEndpoint(endpoint))
		if err != nil {
			logger.Log("err", err)
			os.Exit(1)
		}
		// use zipkin-go-opentracing to wrap our tracer
		tracer = zipkingoopentracing.Wrap(nativeTracer)
		// optionally set as Global OpenTracing tracer instance
		// opentracinggo.SetGlobalTracer(tracer)
    }
}
```

执行`go run cmd/service/main.go`, 测试地址: `http://localhost:8080/metrics`

> 关于`zipkin`包升级代码不兼容的问题, 可使用这个kit库：[https://github.com/GrantZheng/kit](https://github.com/GrantZheng/kit)
