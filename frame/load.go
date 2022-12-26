package frame

import (
	"fmt"
)

func InitConfig() {
	CyiManege.loadPlatform()
	CyiManege.loadChannel()
	CyiManege.loadlanguage()
	CyiManege.loadFilterWord()
	CyiManege.loadwhiteList()
	CyiManege.loadHttpRoute()

}

func (o *cyiManageInfo) loadPlatform() {
	o.Platform = "platform"
	fmt.Println("平台加载完成")
}

func (o *cyiManageInfo) loadChannel() {
	o.Channel = "channel"
	fmt.Println("渠道加载完成")
}

func (o *cyiManageInfo) loadlanguage() {
	o.Language = "language"
	fmt.Println("语言加载完成")
}

func (o *cyiManageInfo) loadFilterWord() {
	o.FilterWord = map[string]string{
		"敏感词": "敏感词",
	}
	fmt.Println("敏感词加载完成")
}

func (o *cyiManageInfo) loadwhiteList() {
	o.WhiteList = map[string]string{
		"白名单": "白名单",
	}
	fmt.Println("白名单加载完成")
}

func (o *cyiManageInfo) loadHttpRoute() {
	o.HttpRoute = map[string]string{
		"/Action10001": "Action10001",
		"/Action10002": "Action10002",
		"/Banshichu1":  "Banshichu1",
		"/Banshichu2":  "Banshichu2",
	}
	fmt.Println("http监听路由加载完成")
}
