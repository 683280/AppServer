package main

import (
	"net/http"
	"io/ioutil"
)

func handler(rw http.ResponseWriter,r *http.Request){
	r.ParseForm()
	data,_ := ioutil.ReadAll(r.Body)
	println(string(data))
	println("--------------")
}
func main() {
	http.HandleFunc("/log",handler)
	http.ListenAndServe(":8081",nil)
}
