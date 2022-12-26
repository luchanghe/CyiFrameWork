package main

import (
	"Cyi/frame"
)

func main() {
	//加载需要的配置
	frame.InitConfig()
	//启用监听
	frame.StartListen()

}
