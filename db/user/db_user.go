package user_db

import ("../../db"
	"strconv"
	"fmt"
	"../../protocol"
)
var dbInstance = db.GetInstance()

type User struct {
	User_id		int `json:"user_id"`
	User_uuid	string	`json:"user_uuid"`
	User_name	string	`json:"user_name"`
	User_password	string	`json:"user_password"`
	User_mail	string	`json:"user_mail"`
	User_desc	string	`json:"user_desc"`
	User_autograph	string	`json:"user_autograph"`
	User_gender	string	`json:"user_gender"`
	User_location	string	`json:"user_location"`
	User_head	string	`json:"user_head"`
}
type Message struct {
	Msg_from int `json:"msg_from"`
	Msg_to int `json:"msg_to"`
	Msg_data string `json:"msg_data"`
	Time int `json:"time"`

}
/**
CREATE TABLE `user` (
	`user_id` int(11) NOT NULL,
	`user_uuid` char(36) NOT NULL,
	`user_name` char NOT NULL,
	`user_password` char NOT NULL,
	`user_mail` char NOT NULL,
	`user_desc` char NOT NULL,
	`user_autograph` char NOT NULL,
	`user_gender` char NOT NULL,
	`user_location` char NOT NULL,
	PRIMARY KEY (`user_id`)
) ENGINE=InnoDB
DEFAULT CHARACTER SET=utf8 COLLATE=utf8_general_ci;
*/
func Login(username,password string)(*User){
	var users = db.GetInstance().SelectSql(db.SQL_LoginUser,&User{},username,password)
	if users != nil && len(users) > 0 {
		u := users[0].(User)
		return &u
	}
	return nil
}

func Register(username,password string)(*User){
	uuid := protocol.GetGuid()
	var result = db.GetInstance().InsertSql(db.SQL_RegisterUser,uuid,username,password)
	if result == 0 {
		return nil
	}
	user_id := strconv.FormatInt(result,10)
	fmt.Println(user_id)
	user := LoginById(user_id,password)
	return user
}
func LoginById(id,password string)(*User){
	var users = db.GetInstance().SelectSql(db.SQL_LoginUserById,&User{},id,password)
	if users != nil && len(users) > 0 {
		u := users[0].(User)
		return &u
	}
	return nil
}
func NameIsExis(name string)bool{
	var users = db.GetInstance().SelectSql(db.SQL_NameIsExis,1,name)
	if users != nil && len(users) > 0 {
		u := users[0]
		return u != 0
	}
	return false
}

func GetUUIDById(id int)(string){
	uuid := db.GetInstance().SelectSql(db.SQL_GetUUIDById,"",id)
	if uuid != nil && len(uuid) > 0 {
		u := uuid[0].(string)
		return u
	}
	return ""
}
func GetUserById(id int)*User{
	var users = db.GetInstance().SelectSql(db.SQL_GetUserById,&User{},id)
	if users != nil && len(users) > 0 {
		u := users[0].(User)
		return &u
	}
	return nil
}
func GetAllFriends(uuid string)interface{}{
	friends := db.GetInstance().SelectSql(db.SQL_GetAllFriends,&User{},uuid)
	if friends != nil && len(friends) > 0 {
		return friends
	}
	return nil
}
func GetUserAllMsg(id int)*[]interface{}{
	datas := db.GetInstance().SelectSql(db.SQL_GetUserAllMsg,&Message{},id)
	var in interface{}

	if datas != nil && len(datas) > 0 {
		in = datas
		var msgs = in.([]interface{})
		return &msgs
	}
	return nil
}
func DelUserMsgByTime(id int)*[]interface{}{
	datas := db.GetInstance().SelectSql(db.SQL_DelUserMsgByTime,&Message{},id)
	var in interface{}

	if datas != nil && len(datas) > 0 {
		in = datas
		var msgs = in.([]interface{})
		return &msgs
	}
	return nil
}



























