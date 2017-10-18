package main

import (
	."./server"
	"./im"
	_"./server/http"
)
func main() {

	go im.StartImServer()
	StartHttpServer()

}
