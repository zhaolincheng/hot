package controller

import (
	"fmt"
	"net/http"
)

func login(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method:", r.Method) // 获取请求的方法
	fmt.Println(r.Header.Get("Content-Type"))
	// 请求的是登录数据，那么执行登录的逻辑判断
	fmt.Println("username:", r.Form["username"])
	fmt.Println("password:", r.Form["password"])
}
