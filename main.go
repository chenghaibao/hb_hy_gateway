package main

import (
	"hb_hy_gateway/cmd"
	"log"
	"runtime"
)

func main() {
	// 全部核心运行程序
	runtime.GOMAXPROCS(runtime.NumCPU())
	// 日志等级
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)

	cmd.Execute()
}