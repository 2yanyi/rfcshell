# RFC Shell
HTTP remote function call development framework.

<br>

1. 定义全局请求钩子

```go
func requestHookFunction(w http.ResponseWriter, r *http.Request) error {
    fmt.Println("call requestHookFunction")
    return rfcshell.Next
}
```

2. 设计调用路由

```go
func router(sh *rfcshell.ServerHandler) *rfcshell.ServerHandler {
    user := sh.Route("/user", handle1)
    user.Route("/info", handle2, handle3)
    return sh
}
```

3. 启动服务，导入请求钩子和路由方法。

```go
func main() {
    rfcshell.SetRequestHookFunction(requestHookFunction)
    println(
        router(rfcshell.New("127.0.0.1:2000")).Server.ListenAndServe(),
	)
}
```
