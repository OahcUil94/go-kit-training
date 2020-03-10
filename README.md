# go-kit微服务套件使用

## ServerOption

```
ServerOption为Serve设置可选的函数调用, 有以下几种: 

1. ServerBefore: 在调用decode函数之前执行，在HTTP请求对象上执行ServerBefore函数
2. ServerAfter: 在调用endpoint之后, encode函数之前执行
3. ServerErrorHandler: 收集decode, endpoint, encode中返回的第二个参数的错误对象信息, 简单的收集log 
4. ServerErrorEncoder: 收集decode, endpoint, encode中返回的第二个参数的错误对象信息, 并可以写入到http.ResponseWriter返回客户端
5. ServerFinalizer: 在每个HTTP请求结束时执行，在encode或者ServerErrorEncoder之后执行的函数

正常的请求流程: ServerBefore -> decode -> endpoint -> service -> ServerAfter -> encode -> ServerFinalizer
出现错误的请求流程: ServerBefore -> 出现错误(decode -> endpoint -> encode) -> ServerErrHandler -> ServerErrorEncoder(可写httpResponse) -> ServerFinalizer
```
