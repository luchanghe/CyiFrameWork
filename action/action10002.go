package action

import (
	"fmt"
	"net/http"
)

func (*Action) Action10002(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprint(writer, "成功进入Action10002")
}
