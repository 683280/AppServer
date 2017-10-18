package im

import (
	"net"
	"fmt"
	."../db/user"
	"../protocol"
	"../cryption"
	"log"
	"time"
	"encoding/json"
)
//var publicKey = []byte(`MIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQC519uh7Yp4AkzB2hQGgJ7Mvgcz
//tPRvdjHQkpA+UMwA1n1HzYA06SHXF021gXUYilBxgfpzbqQaEkvaLrwqlxslDfK7
//Al2mA3eM0EjusoFQF+v6VT65dC2TzpHoQeblC2b9xCwlyUXoH0uVIEcuAKSKZoMZ
//Qfxr3ohFb3TL5zTyPwIDAQAB`)
//var privateKey = []byte(`-----BEGIN RSA PRIVATE KEY-----
//MIICXAIBAAKBgQC519uh7Yp4AkzB2hQGgJ7MvgcztPRvdjHQkpA+UMwA1n1HzYA0
//6SHXF021gXUYilBxgfpzbqQaEkvaLrwqlxslDfK7Al2mA3eM0EjusoFQF+v6VT65
//dC2TzpHoQeblC2b9xCwlyUXoH0uVIEcuAKSKZoMZQfxr3ohFb3TL5zTyPwIDAQAB
//AoGAYLSxxp5sWqyfspQ/rW6Ks/ICn2Z/d+ziWS2bP8IdliYHBTErkNzrzhiDSHr4
//Ku/2kkpXwG+Hl0WEESIWqnb9GTPTmDXWhui9ZdoJ3wSf/B8JDfsMVR6R2uw2h7ig
//Z/LSm0ItCBEej6dhmkxC8XJE29/+Q+DQVdKmyuQPZSG1DmkCQQDyhBN3DngZ1+m8
//A5uQbiBOrH+u6CeMRWl2ZGHvjNbDAMJiGi+5jJeqNZpHf7WJ5Ou+eBxFqmUV1/jj
//Xj+7ACarAkEAxC0dCvg/GdJVKBsCW4vGlmV+JagNuL4f1afr5K2pKKJret658Hn1
//rfwIHp6Tj7S2I23UW+LgwVevqpm9faEyvQJANGUCi5NNsU+riNpCrsaMJlMwVsqD
//WNPaQCDZ49ZKw+CTHnzH2M+eKMDh7xaRUxRpNkJe4VI5+qkpdX30SON0dwJAYQVe
//w7oamw6nBvq0o8nxIRh41u7SOnftDqHJzIMGkg4h0datZv0qQC3RZjNPD1d0bPk4
//eWkvdu+C9YCrcqJykQJBAI8fq3oDk121FCP0YjDj/FglDuEX9O+Xinv0FOBE1cbE
//PL4vDFtaG5JRDdgTsIr87mKsssABd1DR/qZMxwCmDsc=
//-----END RSA PRIVATE KEY-----
//`)
type ClientConn struct {
	conn net.Conn
	msg chan *Pack
	send chan *Pack
	user *User
	status bool
	uuid []byte
}
type Pack struct {
	data []byte
	time int
	t int
	length int
}
var Method = make(map[int]func(*Pack,* ClientConn))
var Clients = make(map[int]*ClientConn)
func (s *ClientConn)write(p *Pack){
	s.conn.Write([]byte{1,2})
	data,_ := cryption.AesEncrypt(p.data,s.uuid)
	s.conn.Write(protocol.IntToBytes(p.t))
	s.conn.Write(protocol.IntToBytes(p.time))
	s.conn.Write(protocol.IntToBytes(len(data)))
	s.conn.Write(data)
	s.conn.Write([]byte{10,4})
}
func (s *ClientConn)writeString(data string,t int){
	d := []byte(data)
	s.write(&Pack{d,int(time.Now().Unix()),t,len(d)})
}
func (s *ClientConn)sendData(data []byte,t int){
	s.send <- &Pack{data,int(time.Now().Unix()),t,len(data)}
}
func (s *ClientConn)sendPack(data *Pack){
	s.send <- data
}
func (s *ClientConn)read()(*Pack)  {
	b := make([]byte,4)
	for true {
		i,err := s.conn.Read(b[:1])
		if err != nil || i == 0{
			log.Println(err)
			return nil
		}
		if b[0] != byte(1){
			log.Println("读取头错误")
			continue
		}
		s.conn.Read(b[1:2])
		if err != nil || i == 0{
			log.Println(err)
			return nil
		}
		if b[1] != byte(2){
			log.Println("读取头错误")
			continue
		}
		break
	}
	s.conn.Read(b[:4])
	t := protocol.BytesToInt(b[:4])
	s.conn.Read(b[:4])
	time := protocol.BytesToInt(b[:4])
	s.conn.Read(b[:4])
	l := protocol.BytesToInt(b[:4])
	if l == 0 {
		return nil
	}
	d := make([]byte,l)
	n,err := s.conn.Read(d)
	if err != nil || n != l{
		return nil
	}
	i,err := s.conn.Read(b[:2])
	if err != nil || i == 0{
		log.Println(err)
		return nil
	}
	if b[0] != byte(10) && b[1] != byte(4) {
		return nil
	}
	d,err = cryption.AesDecrypt(d,s.uuid)
	if err != nil {
		log.Panicln(err)
		return nil
	}
	p := Pack{d,time,t,len(d)}
	return &p
}
func StartImServer(){
	addr,err := net.ResolveTCPAddr("tcp4",":9999")
	if err != nil {
		fmt.Println(err)
		return
	}
	list,err := net.ListenTCP("tcp",addr)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("开启tcp成功")
	for {
		conn,err := list.Accept()
		if err != nil {
			fmt.Println(err)
			return
		}
		client := ClientConn{conn,make(chan *Pack),make(chan *Pack),nil,true,nil}
		go handlerConn(&client)
	}
}

func handlerConn(client *ClientConn){
	if !tcpLogin(client){return}

	go loopWrite(client)
	go loopRead(client)
	go client.handlerMsg()
	//发送离线消息
	msg := GetUserAllMsg(client.user.User_id)
	if msg != nil && len(*msg) > 0 {
		msgs := *msg
		for i:= 0;i < len(msgs) ; i++{
			m := msgs[i].(Message)
			a := SendMsg{}
			a.From = m.Msg_from
			a.To = m.Msg_to
			a.Data = m.Msg_data
			a.Time = m.Time
			bytes,_ := json.Marshal(a)
			client.sendData(bytes,100)
			DelUserMsgByTime(a.Time)
		}
	}
}
func tcpLogin(client *ClientConn)(bool)  {
	b := make([]byte,4)
	client.conn.Read(b)
	id := protocol.BytesToInt(b)
	//client.user.
	fmt.Println(id)
	uuid := GetUUIDById(id)
	fmt.Println(uuid)
	if len(uuid) == 0{
		return false
	}
	client.uuid = []byte(uuid)
	client.read()
	//userid := protocol.BytesToInt(b)
	c := Clients[id]
	if c != nil {

	}
	Clients[id] = client
	client.writeString("成功",1)
	client.user = GetUserById(id)
	return true
}

func loopRead(client *ClientConn)  {
	for client.status{
		data := client.read()
		if data == nil {
			client.conn.Close()
			Clients[client.user.User_id] = nil
			return
		}
		client.msg <- data
	}

	client.status = false
	client.send <- nil
}
func loopWrite(client *ClientConn)  {
	for client.status  {
		d :=<-client.send
		client.write(d)
	}
}
func (c *ClientConn)handlerMsg(){
	for c.status{
		d := <- c.msg
		method := Method[d.t]
		if method == nil {
			continue
		}
		method(d,c)
	}
}
