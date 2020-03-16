# kitgen

记录时间: 2020-03-16

地址: [https://github.com/go-kit/kit/tree/master/cmd/kitgen](https://github.com/go-kit/kit/tree/master/cmd/kitgen)

kitgen用来生成一些样板代码, 该工具还处于实验阶段

## 安装

把go-kit/kit项目clone到本地, `cd $GOPATH/src/github.com/go-kit/kit/cmd/kitgen`, 执行安装命令: `go install`

## 常用命令

查看帮助信息: `kitgen -h` 

```
USAGE
  kitgen [flags] path/to/service.go

FLAGS
  -h	打印帮助信息
  -repo-layout string
    	default, flat... (default "default")
  -target-dir string
    	base directory to emit into (default ".")
```

## 生成代码

创建`service/service.go`文件, 并添加下面的代码: 

```
// service.go
package profilesvc // don't forget to name your package
import "context"

type Service interface {
    PostProfile(ctx context.Context, p Profile) error
    // ...
}
type Profile struct {
    ID        string    `json:"id"`
    Name      string    `json:"name,omitempty"`
    // ...
}
```

执行`kitgen ./service.go`, 则会在当前目录下生成如下的目录结构: 

```
.
├── endpoints
│   └── endpoints.go
├── http
│   └── http.go
├── service
│   └── service.go
└── service.go
```

执行`kitgen -repo-layout flat ./service.go`, 则会把上面生成的多个文件的代码合并到一个文件`go-kit.go`中: 

```
.
├── gokit.go
└── service.go
```
