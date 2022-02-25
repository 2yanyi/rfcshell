# RFC Shell
HTTP remote function call development framework.

<br>

设计调用路由

```go
func router(sh *rfcshell.ServerHandler) *rfcshell.ServerHandler {
    user := sh.Route("/user", handle1)
    user.Route("/info", handle2, handle3)
    return sh
}
```
