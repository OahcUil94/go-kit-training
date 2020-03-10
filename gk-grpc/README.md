# go-kit中使用grpc

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
