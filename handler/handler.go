package handler

import (
	"github.com/micro/go-micro/util/log"
	"net/http"
)

func Login(w http.ResponseWriter, r *http.Request) {
	// 处理推送的入口
	// 只接受POST请求
	if r.Method != "POST" {
		log.Logf("非法请求")
		http.Error(w, "非法请求", 400)
		return
	}
	r.ParseForm()

}
