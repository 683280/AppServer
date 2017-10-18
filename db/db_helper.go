package db

import (
	"reflect"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"database/sql"
	"strings"
)

var mInstance* dbHelper
type dbHelper struct {
	username string
	password string
	address string
	port string
	dbname string
	Db *sql.DB
}
//var database_address = "rm-wz9v77h2a76d9l47c.mysql.rds.aliyuncs.com" //内网
var database_address = "rm-wz9v77h2a76d9l47co.mysql.rds.aliyuncs.com"//外网

func GetInstance()(*dbHelper)  {
	if mInstance == nil {
		var db1 = dbHelper{"root","683280loveC",database_address,"3306","myapp",nil}
		db1.connect()
		mInstance = &db1
	}
	return mInstance
}
func (self* dbHelper)connect(){

	addr := self.username + ":" + self.password + "@tcp(" + self.address+":" + self.port + ")/" + self.dbname + "?charset=utf8"
	fmt.Println(addr)
	db,err := sql.Open("mysql",addr)
	if err != nil || db == nil{
		fmt.Println(err)
		return
	}
	if db.Ping() != nil {
		fmt.Println("连接失败")
		return
	}
	fmt.Println("连接成功")
	self.Db = db

}

func (self *dbHelper)SelectDb(in interface{},table ,where string){

	value := reflect.Indirect(reflect.ValueOf(in))
	length := value.NumField()
	var s string
	var dest = make([]interface{}, length)
	for i := 0; i < length ; i++ {
		filed := value.Type().Field(i)
		name := filed.Name
		s += ","+name
		dest[i] = value.Field(i).Addr().Interface()
	}
	s = "SELECT " + strings.ToLower(s[1:]) +" FROM " + table
	if len(where) > 0 {
		s += " " + where
	}
	fmt.Println(s)
	rows ,err := self.Db.Query(s)
	if err != nil{
		fmt.Print(err)
		return
	}

	for rows.Next() {
		err = rows.Scan(dest...)
		if err != nil{
			fmt.Println(err)
			continue
		}
	}
}

func (self* dbHelper)InsertSql(sql string,in...interface{})int64{
	stmt,err := self.Db.Prepare(sql)
	if err != nil {
		println(err.Error())
		return 0
	}
	res, err := stmt.Exec(in...)
	if err != nil {
		println(err.Error())
		return 0
	}
	id,err := res.LastInsertId()
	if err != nil {
		println(err.Error())
		return 0
	}
	return id
	//id, err := res.LastInsertId()
}

func (self *dbHelper)SelectSql(sql string,in interface{},data...interface{})([]interface{}){
	rows ,err := self.Db.Query(sql,data...)
	if err != nil{
		fmt.Print(err)
		return nil
	}
	defer rows.Close()

	var datas []interface{}
	if rows.Next(){
		datas = make([]interface{},0)
	}else {
		return nil
	}
	if reflect.TypeOf(in).Kind() == reflect.String {

	  loop:	var s string
		err = rows.Scan(&s)
		if err != nil{
			fmt.Println(err)
			return datas
		}
		datas = append(datas,s)
		if rows.Next(){
			goto loop
		}
		return datas
	}else if reflect.TypeOf(in).Kind() == reflect.Int{
	 loopi:	var s int
		err = rows.Scan(&s)
		if err != nil{
			fmt.Println(err)
			return datas
		}
		datas = append(datas,s)
		if rows.Next(){
			goto loopi
		}
		return datas
	}
	value := reflect.Indirect(reflect.ValueOf(in))
	//types := reflect.TypeOf(in)
	s,_ := rows.Columns()
	length := len(s)
	var dest = make([]interface{}, length)
	for i := 0; i < length ; i++ {
		v := value.Field(i)
		dest[i] = v.Addr().Interface()
	}
	for {
		err = rows.Scan(dest...)
		if err != nil{
			fmt.Println(err)
			return datas
		}
		v := copy(value.Interface())
		datas = append(datas,v)
		if !rows.Next() {
			return datas
		}
	}
	return datas
}
func (self *dbHelper)DelSql(sql string,in []interface{})int64{
	stmt,err := self.Db.Prepare(sql)
	if err != nil {
		println(err.Error())
		return 0
	}
	res, err := stmt.Exec(in...)
	if err != nil {
		println(err.Error())
		return 0
	}
	id,err := res.LastInsertId()
	if err != nil {
		println(err.Error())
		return 0
	}
	return id
	//id, err := res.LastInsertId()
}
func copy(in interface{})(interface{})  {
	return in
}