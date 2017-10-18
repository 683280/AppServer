package server

import ("net/http"
	."../../server"
	"../../db"
)

func init() {
	Methods["/topic/topicForId"] = GetTopicForId
	Methods["/topic/topicById"] = GetTopicById
}
type Topic struct {
	T_id 		int	`json:"t_id"`
	T_imgs 		string	`json:"t_imgs"`
	T_big_imgs	string	`json:"t_big_imgs"`
	T_video		string	`json:"t_video"`
	T_content	string	`json:"t_content"`
	T_v_img 	string	`json:"t_v_img"`
	T_v_width 	int	`json:"t_v_width"`
	T_v_height 	int	`json:"t_v_height"`
	T_urls 		string	`json:"t_urls"`
}

func GetTopicById(w http.ResponseWriter,r *http.Request){
	r.ParseForm()
	topic_id := r.FormValue("topic_id")
	if len(topic_id) <= 0 {
		WriteError(w,ParameterError)
		return
	}
	topics := db.GetInstance().SelectSql(db.SQL_GetTopicForId,&Topic{},topic_id)
	if topics == nil{
		WriteError(w,"没有找到这个帖子")
		return
	}
	WriteJson(w,topics)
}
func GetTopicForId(w http.ResponseWriter,r *http.Request){
	r.ParseForm()
	topic_id := r.FormValue("topic_id")
	if len(topic_id) <= 0 {
		WriteError(w,ParameterError)
		return
	}
	topics := db.GetInstance().SelectSql(db.SQL_GetTopicForId,&Topic{},topic_id)
	if topics == nil{
		WriteError(w,"没有找到这个帖子")
		return
	}
	WriteJson(w,topics)
}