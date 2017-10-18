package server

import ("net/http"
	"../../db/user"
	."../../server"
)

func init() {
	Methods["/user/login"] = login
	Methods["/user/register"] = register
}

func login(w http.ResponseWriter,r *http.Request){
	r.ParseForm()
	username := r.Form["username"]
	password := r.Form["password"]
	id := r.Form["id"]
	var user *user_db.User
	if len(password) == 0 {
		WriteError(w,ParameterError)
		return
	}
	if len(username) > 0 {
		user = user_db.Login(username[0],password[0])
	}else if len(id) > 0 {
		user = user_db.LoginById(id[0],password[0])
	}else{
		WriteError(w,ParameterError)
		return
	}

	if user == nil {
		WriteError(w,"登陆失败")
		return
	}
	WriteJson(w,user)
}
func register(w http.ResponseWriter,r *http.Request){
	r.ParseForm()
	username := r.Form["username"]
	password := r.Form["password"]
	var user *user_db.User
	if len(password) == 0 || len(username) == 0{
		WriteError(w,ParameterError)
		return
	}
	if len(username[0]) > 0 || len(password[0]) > 0 {
		if user_db.NameIsExis(username[0]) {
			WriteError(w,"注册失败! 用户名已存在")
			return
		}
		user = user_db.Register(username[0],password[0])
	}else{
		WriteError(w,ParameterError)
		return
	}
	if user == nil {
		WriteError(w,"注册失败")
		return
	}
	WriteJson(w,user)
}

func name_is_exis(w http.ResponseWriter,r *http.Request){
	//r.ParseForm()
	//
	//NameIsExis
}