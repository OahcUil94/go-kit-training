# go-kit中使用grpc

## grpc安装

### 安装编译程序protoc

网址：[https://github.com/protocolbuffers/protobuf/releases](https://github.com/protocolbuffers/protobuf/releases)

根据操作系统情况选择下载，例如：`protoc-3.11.3-osx-x86_64.zip`，下载之后解压，文件目录大致：

```
.
├── bin
│   └── protoc
├── include
│   └── google
│       └── protobuf
│           ├── any.proto
│           ├── api.proto
│           ├── compiler
│           │   └── plugin.proto
│           ├── descriptor.proto
│           ├── duration.proto
│           ├── empty.proto
│           ├── field_mask.proto
│           ├── source_context.proto
│           ├── struct.proto
│           ├── timestamp.proto
│           ├── type.proto
│           └── wrappers.proto
└── readme.txt
```

找到其中的`protoc`文件，放到`go/bin`目录下，或者放到其他的已加入环境变量的目录下

### 安装编译插件: protoc-gen-go

该插件会将protobuf文件转换成go代码的编译插件

`go get github.com/golang/protobuf/protoc-gen-go`

插件`protoc-gen-go`会以二进制形式，直接安装到`go/bin`目录下

## 常见问题

### proto3的rpc语法中是否允许空请求或响应

官方开发人员是这样评论的：作为开发人员，我们真的很难猜测未来的需求。因此，我建议始终为每个方法定义参数和结果，即使为空也是安全的。

```
// A generic empty message that you can re-use to avoid defining duplicated
// empty messages in your APIs. A typical example is to use it as the request
// or the response type of an API method. For instance:
//
//     service Foo {
//       rpc Bar(google.protobuf.Empty) returns (google.protobuf.Empty);
//     }
//
```

> 参考：[https://stackoverflow.com/questions/31768665/can-i-define-a-grpc-call-with-a-null-request-or-response](https://stackoverflow.com/questions/31768665/can-i-define-a-grpc-call-with-a-null-request-or-response)
