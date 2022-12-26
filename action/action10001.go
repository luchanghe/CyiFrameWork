package action

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Action10001 struct {
	Account  string `json:"Account"`
	Password string `json:"Password"`
}

func (*Action) Action10001(writer http.ResponseWriter, request *http.Request) {
	reqData, _ := io.ReadAll(request.Body)
	reqJsonStruct := Action10001{}
	if err := json.Unmarshal(reqData, &reqJsonStruct); err != nil {
		writer.WriteHeader(400)
		return
	}
	fmt.Fprint(writer, "账号为"+reqJsonStruct.Account+"密码为"+reqJsonStruct.Password)
}
