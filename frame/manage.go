package frame

var CyiManege cyiManageInfo

type cyiManageInfo struct {
	Platform   string
	Channel    string
	Language   string
	FilterWord map[string]string
	WhiteList  map[string]string
	HttpRoute  map[string]string
}
