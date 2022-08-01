package rfcshell

import (
	"fmt"
	"net/http"
	"testing"
)

// 定义全局请求 HOOK
func requestHookFunction(w http.ResponseWriter, r *http.Request) error {
	fmt.Println("call requestHookFunction")

	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization, x-requested-with")
	w.Header().Set("Access-Control-Allow-Origin", "*") // 允许跨域
	w.Header().Set("Access-Control-Max-Age", "3600")   // 缓存时间(OPTIONS)

	return Next
}

// 设计调用路由
func router(sh *ServerHandler) *ServerHandler {
	user := sh.Route("/user", handle1)
	user.Route("/info", handle2, handle3)
	return sh
}

// Test main
func Test(t *testing.T) {
	SetRequestHookFunction(requestHookFunction)
	if err := router(New("127.0.0.1:2000")).Server.ListenAndServe(); err != nil {
		println(err.Error())
	}
}

// test handles ...

func handle1(w http.ResponseWriter, r *http.Request) error {
	fmt.Println("call handle1")
	return Next
}

func handle2(w http.ResponseWriter, r *http.Request) error {
	fmt.Println("call handle2")
	return Next
}

func handle3(w http.ResponseWriter, r *http.Request) error {
	fmt.Println("call handle3")
	w.Write([]byte("Hello World"))
	return nil
}
