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
			isChecker := g.checker(w,r)
			if isChecker == true {
				g.sendRequest(w,r)
			}
		}
	})
	return g
}

func (g *Gateway) ListenAndServe() error {
	// 监听端口
	return g.server.ListenAndServe()
}
