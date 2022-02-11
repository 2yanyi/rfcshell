package main

import (
	"context"
	"net/http"
	"r/console"
	"r/rfcshell"
	"time"
)

func main() {
	core := rfcshell.New(":2000")
	core.FrontHook = func(w http.ResponseWriter, r *http.Request) error {
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization, x-requested-with")
		w.Header().Set("Access-Control-Allow-Origin", "*") // 允许跨域
		w.Header().Set("Access-Control-Max-Age", "3600")   // 缓存时间(OPTIONS)
		return nil
	}
	user := core.Route("/user", test1)
	user.Route("/info", test2, test3)
	bind := core.Server.ListenAndServe()
	println(bind.Error())

	core.Server.Shutdown(context.Background())
}

func test1(w http.ResponseWriter, r *http.Request) error {
	//fmt.Println("call /user", r.Header.Get("User"))
	return nil
}

func test2(w http.ResponseWriter, r *http.Request) error {
	//fmt.Println("call /user/info")
	return rfcshell.Next
}

func test3(w http.ResponseWriter, r *http.Request) error {
	//fmt.Println("call /user/info2")
	time.Sleep(time.Millisecond * 200)
	return write(w, writeJson{
		Code:    0,
		Message: "success",
		Data:    map[string]string{"result": "234234"},
	})
}

func write(w http.ResponseWriter, data interface{}) error {
	jss, _ := console.JsonMarshal(data, "")
	_, _ = w.Write(jss)
	return nil
}

type writeJson struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    interface{}
}
