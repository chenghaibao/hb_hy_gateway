package http

import (
	"fmt"
	db2 "hb_hy_gateway/db"
	"hb_hy_gateway/utils"
	"io"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"
)

var TraceId string
var db *db2.Db

func init() {
	db = db2.GetDbInstance()
}

func (g *Gateway) checker(w http.ResponseWriter, r *http.Request) bool {
	// 入门网关错误
	TraceId = utils.GetTraceId()
	// 获取token 不验证鉴权
	token := g.GetToken(w, r)
	if token == false {

		// 不会获取token 验证redis token 验证leveldb(模拟  正常使用etcd consul做注册中心)
		isToken := g.isToken(w, r)
		if isToken == false {
			return false
		}

		// 验证注册中心是否含有当前服务，接口地址是否正确 等等
		isService := g.isService(w, r)
		if isService == false {
			return false
		}
	}
	return true
}

func (g *Gateway) GetToken(w http.ResponseWriter, r *http.Request) bool {
	// 获取token
	h := w.Header()
	if r.URL.Path == "/get-token" {
		h.Set("x-stream-id", TraceId)
		g.jsonSimpleTextSuccess(w, "success")
		return true
	}
	return false
}

func (g *Gateway) isToken(w http.ResponseWriter, r *http.Request) bool {
	// 验证redis  是否存在当前token
	token := r.Header.Get("token")
	dbToken, _ := db.GetStringKey("token")
	// 参数传过去  去获取token 返回前端
	if token != string(dbToken) {
		//报错
		g.jsonErrorWithTraceId(w, TraceId, 404, "token not exist")
		return false
	}
	return true
}

func (g *Gateway) isService(w http.ResponseWriter, r *http.Request) bool {
	// 验证redis  是否存在当前token
	token := r.Header.Get("service")
	dbToken, _ := db.HasStringKey(token)
	// 参数传过去  去获取token 返回前端
	fmt.Println(dbToken)
	if dbToken == false {
		//报错
		g.jsonErrorWithTraceId(w, TraceId, 404, "service not exist")
		return false
	}
	return true
}

func (g *Gateway) jsonSimpleTextSuccess(w http.ResponseWriter, msg string) {
	w.Header().Set("Content-Type", "application/json;charset=utf-8")
	_, _ = w.Write([]byte(`{"success":true,"msg":"` + msg + `"}`))
}

func (g *Gateway) jsonStreamSuccess(w http.ResponseWriter, bodyReader io.Reader) {
	w.Header().Set("Content-Type", "application/json;charset=utf-8")

	if bodyReader != nil {
		_, err := w.Write([]byte(`{"success":true,"msg":`))
		if err != nil {
			return
		}

		_, err = io.Copy(w, bodyReader)
		if err != nil {
			return
		}

		_, err = w.Write([]byte(`}`))
		if err != nil {
			return
		}
	} else {
		_, _ = w.Write([]byte(`{"success":true,"msg":null}`))
	}
}

func (g *Gateway) jsonError(w http.ResponseWriter, errCode uint64, errMsg string) {
	w.Header().Set("Content-Type", "application/json;charset=utf-8")

	if errCode < 600 && errCode >= 400 {
		w.WriteHeader(int(errCode))
	} else {
		w.WriteHeader(500)
	}
	_, _ = w.Write([]byte(`{"success":false,"error_code":` + strconv.FormatUint(errCode, 10) + `,"error_msg":"` + errMsg + `"}`))
}

func (g *Gateway) jsonErrorWithTraceId(w http.ResponseWriter, traceId string, errCode uint64, errMsg string) {
	w.Header().Set("Content-Type", "application/json;charset=utf-8")

	if traceId == "" {
		g.jsonError(w, errCode, errMsg)
	} else {
		if errCode < 600 && errCode >= 400 {
			w.WriteHeader(int(errCode))
		} else {
			w.WriteHeader(500)
		}
		_, _ = w.Write([]byte(`{"success":false,"error_code":` + strconv.FormatUint(errCode, 10) + `,"error_msg":"` + errMsg + `","trace_id":"` + traceId + `"}`))
	}
}

func SignalListen() {
	sigs := make(chan os.Signal)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM, syscall.SIGTSTP, syscall.SIGSTOP, syscall.SIGHUP, syscall.SIGUSR1)
	for {
		<-sigs
		fmt.Println("Bye Bye")
		os.Exit(1)
	}
}
