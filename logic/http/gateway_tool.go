package http

import (
	"fmt"
	"hb_hy_gateway/utils"
	"net/http"
)

var TraceId string

func (g *Gateway) GetToken(w http.ResponseWriter, r *http.Request) {
	TraceId = utils.GetTraceId()
	// 获取token
	h := w.Header()
	if r.URL.Path == "/get-token"{
		h.Set("x-stream-id",TraceId)
		// 参数传过去  去获取token 返回前端
		fmt.Println("12312")
	}
}



func (g *Gateway) isToken(w http.ResponseWriter, r *http.Request) {
	// 验证redis  是否存在当前token
	h := w.Header()
	if r.URL.Path == "/get-token"{
		h.Set("x-stream-id",TraceId)
		// 参数传过去  去获取token 返回前端
		fmt.Println("12312")
	}
}
func (g *Gateway) isService(w http.ResponseWriter, r *http.Request) {
	// 验证etcd  是否存在当前服务
	h := w.Header()
	if r.URL.Path == "/get-token"{
		h.Set("x-stream-id",TraceId)
		// 参数传过去  去获取token 返回前端
		fmt.Println("12312")
	}
}