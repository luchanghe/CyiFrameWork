package frame

import (
	"Cyi/action"
	"fmt"
	"net/http"
	"reflect"
)

func StartListen() {
	listenHttp()
	listenWebSocket()
	listenTcp()
}

func listenHttp() {

	var validateFunc action.Action
	for key, name := range CyiManege.HttpRoute {
		routeFunc := name
		fmt.Println(key, routeFunc)
		http.HandleFunc(key, func(writer http.ResponseWriter, request *http.Request) {
			reflect.ValueOf(&validateFunc).MethodByName(routeFunc).Call(valOf(writer, request))
		})
	}

	err := http.ListenAndServe("127.0.0.1:8080", nil)
	if err != nil {
		return
	}
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
