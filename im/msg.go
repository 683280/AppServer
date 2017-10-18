package im

import (
	"fmt"
	"encoding/json"
	"../db"
)

func init() {
	Method[SEND_MSG] = sendMsg
}

const (
	SEND_MSG = 100
)
type SendMsg struct {
	To int	`json:"msg_to"`
	From int	`json:"msg_from"`
	Data string 	`json:"msg_data"`
	Time int 	`json:"time"`
}
func sendMsg(data *Pack,conn *ClientConn)  {
	//conn.w
	var send = SendMsg{}
	err := json.Unmarshal(data.data,&send)
	if err != nil{
		return
	}
	user := Clients[send.To]
	if user != nil {
		user.sendPack(data)
	}else{
		_ = conn.user.User_id
		db.GetInstance().InsertSql(db.SQL_InsertMessage,conn.user.User_id,send.To,send.Data,data.time)
	}
	//conn.sendPack(data)
	fmt.Println(string(data.data) + "-----------")
}
