package server

import(."../../server"
	"net/http"
	"../../db/user"
)

func init() {
	Methods["/friend/get_my_friends"] = get_my_friends
}
func get_my_friends(w http.ResponseWriter,r *http.Request){
	r.ParseForm()
	uuid := r.Form["uuid"]
	if len(uuid) == 0 || len(uuid[0]) != 16{
		WriteError(w,ParameterError)
		return
	}
	friends := user_db.GetAllFriends(uuid[0])
	WriteJson(w,friends)
}