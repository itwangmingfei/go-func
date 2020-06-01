package main

import (
	"io"
	"net/http"
)
//简单访问
func main(){
	http.HandleFunc("/", ControllerD)
	http.ListenAndServe(":8080",nil)
}
//页面helow输出
func ControllerD(write http.ResponseWriter,request *http.Request){
	io.WriteString(write,"hellow world")
}
