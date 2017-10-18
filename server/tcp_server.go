package server

import (
	"net"
	"fmt"
	"encoding/binary"
	"bytes"
)
func GetAppServer(port string)(*appServer){

	add,err := net.ResolveTCPAddr("tcp4",":" + port)
	if err != nil{
		 //return new error()
	}
	return &appServer{add,make(chan []byte,16),make(chan []byte,16)}
}
type appServer struct {
	addr *net.TCPAddr
	sendQuq,recvQue chan []byte
}

func StartServer(server appServer)(int){
	listen,err := net.ListenTCP("tcp",server.addr)
	if err != nil{
		return -1
	}
	for  {
		conn,err := listen.Accept()
		if err != nil{
			fmt.Print(err)
			continue
		}
		server.handlerConn(conn)
	}
}

func (self appServer)handlerConn(conn net.Conn){
	go self.sendThread(conn)
	go self.recvThread(conn)
}
func (self appServer)sendThread(conn net.Conn){
	for true{
		data := <-self.sendQuq
		conn.Write(data)
	}
}
func (self appServer)handerRecv(){
	for true  {
		self.sendQuq <- <-self.recvQue
	}
}
func (self appServer)recvThread(conn net.Conn){
	defer conn.Close()
	defer close(self.sendQuq)
	lenBuffer := make([]byte,4)
	buffer := make([]byte,2048)
	for true  {
		len,err := conn.Read(lenBuffer)
		if err != nil || len == 0{
			return
		}
		if len != 4{
			fmt.Println("length != 4")
			continue
		}
		len = bytesToInt(lenBuffer)
		tem := buffer[:len]
		lens,err := conn.Read(tem)
		if  err != nil{
			return
		}
		if lens != len {

		}
		data := make([]byte,lens)
		copy(data,tem)
		self.recvQue <- data
	}
}
func bytesToInt(b []byte)(int){
	buffer := bytes.NewBuffer(b)
	var i int32
	binary.Read(buffer,binary.BigEndian,&i)
	return int(i)
}