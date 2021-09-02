package http

import (
	"context"
	"hb_hy_gateway/config"
	"net/http"
	"strings"
	"time"
)


type Gateway struct {
	server *http.Server
}

func NewGateway(ctx context.Context) *Gateway {
	// 初始化网关
	g := &Gateway{
		server: &http.Server{
			Addr:              config.Viper.GetString("gateway.http.listen_addr"),
			ReadTimeout:       config.Viper.GetDuration("gateway.http.timeout.read") * time.Second,
			ReadHeaderTimeout: config.Viper.GetDuration("gateway.http.timeout.header_read") * time.Second,
			WriteTimeout:      config.Viper.GetDuration("gateway.http.timeout.write") * time.Second,
			MaxHeaderBytes:    config.Viper.GetInt("gateway.http.max_header_size"),
		},
	}

	//返回处理
	g.server.Handler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// 入口网关符合
		if isEnterIngress := strings.Contains("127.0.0.1", r.URL.Host); isEnterIngress {
			// 获取token 不验证鉴权
			g.GetToken(w, r)
			// 不会获取token 验证redis token 验证leveldb(模拟  正常使用etcd consul做注册中心)
			g.isToken(w, r)
			// 验证注册中心是否含有当前服务，接口地址是否正确 等等
			g.isService(w, r)
		}

		// 入门网关错误

	})
	return g
}

func (g *Gateway) ListenAndServe() error {
	// 监听端口
	return g.server.ListenAndServe()
}
