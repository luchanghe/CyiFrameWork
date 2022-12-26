package frame

import (
	"Cyi/action"
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"reflect"
)

func StartListen() {
	listenHttp()
	listenWebSocket()
	listenTcp()
	//阻止主程序关闭
	select {}
}

type ActionJson struct {
	Action    string `json:"Action"`
	SecretKey string `json:"SecretKey"`
}

func listenHttp() {

	go func() {
		var validateFunc action.Action
		http.HandleFunc("/request", func(writer http.ResponseWriter, request *http.Request) {
			if request.Method != "POST" {
				writer.WriteHeader(405)
				return
			}
			reqData, err := io.ReadAll(request.Body)
			if err != nil {
				writer.WriteHeader(400)
				return
			}
			reqJsonStruct := ActionJson{}
			if err := json.Unmarshal(reqData, &reqJsonStruct); err != nil {
				writer.WriteHeader(400)
				return
			}
			//验证令牌
			request.Body = io.NopCloser(bytes.NewBuffer(reqData))
			reflect.ValueOf(&validateFunc).MethodByName(reqJsonStruct.Action).Call(valOf(writer, request))
		})
		err := http.ListenAndServe(":8080", nil)
		if err != nil {
			return
		}
	}()
}

func listenWebSocket() {

}

func listenTcp() {

}

func valOf(i ...interface{}) []reflect.Value {
	var rt []reflect.Value
	for _, i2 := range i {
		rt = append(rt, reflect.ValueOf(i2))
	}
	return rt
}
