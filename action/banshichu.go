package action

import (
	"fmt"
	"net/http"
)

func (*Action) Banshichu1(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprint(writer, "成功进入办事处方法1")
}

func (*Action) Banshichu2(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprint(writer, "成功进入办事处方法2")
}
