# RFC Shell
HTTP remote function call development framework.

<br>

**1. 定义全局请求 HOOK**

```go
func requestHookFunction(w http.ResponseWriter, r *http.Request) error {
    fmt.Println("call requestHookFunction")
    return rfcshell.Next
}
```

**2. 设计调用路由**

```go
func router(sh *rfcshell.ServerHandler) *rfcshell.ServerHandler {
    user := sh.Route("/user", handle1)
    user.Route("/info", handle2, handle3) // -/user/info
    return sh
}
```

**3. 监听服务**

```go
func main() {
    rfcshell.SetRequestHookFunction(requestHookFunction)
    if err := router(rfcshell.New("127.0.0.1:2000")).Server.ListenAndServe(); err != nil {
        println(err.Error())
    }
}
```

<br>

**test handles ...**

```go
func handle1(w http.ResponseWriter, r *http.Request) error {
    fmt.Println("call handle1")
    return rfcshell.Next
}

func handle2(w http.ResponseWriter, r *http.Request) error {
    fmt.Println("call handle2")
    return rfcshell.Next
}

func handle3(w http.ResponseWriter, r *http.Request) error {
    fmt.Println("call handle3")
    w.Write([]byte("Hello World"))
    return nil
}
```
