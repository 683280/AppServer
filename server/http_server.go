package server

import (
	"net/http"
	"os"
	"encoding/json"
	"fmt"
	"sync"
	"log"
)
var ParameterError = "参数错误"
type Result struct {
	Success bool	`json:"success"`
	Message string	`json:"message"`
	Data interface{} `json:"data"`
}
type HttpsServer int
type HttpServer int
func PathExists(path string) (bool) {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	}
	return false
}
func (self *HttpServer)ServeHTTP(r http.ResponseWriter,rr *http.Request)()  {
	//uir := rr.Header["Upgrade-Insecure-Requests"]
	//if len(uir) > 0 && strings.EqualFold(uir[0], "1") {
	//	rr.URL.Scheme = "https"
	//	rr.URL.Path = "carljay.top:9090" + rr.URL.Path
	//	fmt.Println(rr.URL.String())
	//	http.Redirect(r,rr, rr.URL.String(), http.StatusFound)
	//	return
	//}
	f := Methods[rr.URL.Path]
	if f == nil {
		path := "." + rr.URL.Path
		if PathExists(path) {
			http.ServeFile(r,rr,"." + rr.URL.Path)
			return
		}
		http.Redirect(r,rr,"/error/index.html",http.StatusFound)
		return
	}
	f(r,rr)
}
func (self *HttpsServer)ServeHTTP(r http.ResponseWriter,rr *http.Request)()  {
	f := Methods[rr.URL.Path]
	if f == nil {
		path := "." + rr.URL.Path
		if PathExists(path){
			http.ServeFile(r,rr,"." + rr.URL.Path)
			return
		}
		http.Redirect(r,rr,"/error/index.html",http.StatusFound)
		return
	}
	f(r,rr)
}

var Methods = map[string]func(http.ResponseWriter,*http.Request){}
var wg  = sync.WaitGroup{}
func StartHttpServer(){
	fmt.Println("开启http中")
	wg.Add(2)
	go startHttpServer()
	go startHttpsServer()
	fmt.Println("开启http")
	wg.Wait()

}
func startHttpsServer(){
	var server HttpsServer
	defer wg.Done()
	err := http.ListenAndServeTLS(":9090","./key/carljay.top.crt","./key/carljay.top.key",&server)
	if err != nil{
		fmt.Println(err)
		return
	}
	log.Println("https 开启成功")
}
func startHttpServer(){
	var server HttpServer
	wg.Add(1)
	defer wg.Done()
	err := http.ListenAndServe(":8080",&server)
	if err != nil{
		fmt.Println(err)
		return
	}
	log.Println("http 开启成功")
}
func WriteError(w http.ResponseWriter,error string){
	result := Result{false,error,nil}
	j,_ := json.Marshal(result)
	w.Write(j)
}
func WriteJson(w http.ResponseWriter,jsondata interface{}){
	result := Result{true,"成功",jsondata}
	j,_ := json.Marshal(result)
	w.Write(j)
}