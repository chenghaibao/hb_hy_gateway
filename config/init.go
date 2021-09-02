package config

import (
	"github.com/spf13/viper"
	"strings"
)

var Viper *viper.Viper

func init() {
	Viper = viper.New()
	viperDefault()
	Viper.SetConfigType("env")
	Viper.AutomaticEnv()
	Viper.SetEnvPrefix("gateway")
	Viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
}

func viperDefault() {
	Viper.SetDefault("gateway.http.listen_addr", "0.0.0.0:9086")
	Viper.SetDefault("gateway.http.timeout.read", 30)
	Viper.SetDefault("gateway.http.timeout.header_read", 5)
	Viper.SetDefault("gateway.http.timeout.write", 65)
	Viper.SetDefault("gateway.http.max_header_size", 8*1024)            // 8k
}
