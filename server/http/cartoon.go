package server

import (
	."../../server"
	"net/http"
	"io/ioutil"
	"regexp"
	"fmt"
	"strconv"
	"strings"
)
type CartoonBean struct {
	Title string `json:"title"`
	Img string `json:"img"`
	Url string `json:"url"`
}
type CartoonTypeBean struct {
	Title string `json:"title"`
	Type string `json:"type"`
}
func init() {
	Methods["/cartoon/get_cartoon_type"] = get_cartoon_type
	Methods["/cartoon/get_cartoon_list"] = get_cartoon_list
	Methods["/cartoon/get_cartoon_detail"] = get_cartoon_detail
}
func get_cartoon_type(w http.ResponseWriter,r *http.Request){
	data := []CartoonTypeBean{	{"无翼鸟","wuyiniao"},
		{"肉肉","rourou"},
		{"王者荣耀","wzry"},
		{"不知火舞","bzhw"},
		{"少女","shaonv"},
		{"本子","benzi"},
		{"十九禁","19jin"},
		{"触手","chushou"},
		{"初音","chuyin"},
		{"本子","tag/本子"},
		{"色系","sexi"},
		{"二次元美女","tag/二次元"},
		{"动漫","dm"}}
	WriteJson(w,data)
}
func get_cartoon_list(w http.ResponseWriter,r *http.Request){
	r.ParseForm()
	typee := r.Form["type"]
	index := r.Form["index"]
	if len(typee) == 0 || len(index) == 0{
		WriteError(w,ParameterError)
		return
	}
	i,err := strconv.Atoi(index[0])
	if err != nil {
		WriteError(w,ParameterError)
		return
	}
	datas := get_cartoon_type_list(typee[0],i)
	WriteJson(w,datas)
}
func get_cartoon_detail(w http.ResponseWriter,r *http.Request){
	r.ParseForm()
	url := r.Form["url"]
	if len(url) == 0 {
		WriteError(w,ParameterError)
		return
	}
	WriteJson(w,get_list(url[0]))
}
func get_cartoon_type_list(typee string,index int)[]CartoonBean{
	url := "http://iktupian.com/" + typee +"/"
	if index > 1 {url = url + ("index_" + string(index) +".html")}
	resp,_ := http.Get(url)
	src,_ := ioutil.ReadAll(resp.Body)
	reg,_ := regexp.Compile("(show_box).*?(</div>)")
	reg2,_ := regexp.Compile("(=\").*?(\")")
	items := reg.FindAllString(string(src),-1)

	d := make([]CartoonBean,len(items))
	for x,i := range items  {
		datas := reg2.FindAllString(string(i),5)
		u := datas[1][2:len(datas[1]) - 1]
		title := datas[3][2:len(datas[3]) - 1]
		img := datas[4][2:len(datas[4]) - 1]

		item := CartoonBean{title,img,u}
		d[x] = item
	}
	return d
}
func get_list(url string)[]string{
	url = "http://iktupian.com" + url
	i := 1
	data := []string{}
	for i == 1  {
		img,next := get_img(url)
		data = append(data,img)
		//data,err := http.Get(img)
		//if err == nil {
		//	has := md5.Sum([]byte(img))
		//	md5str1 := fmt.Sprintf("%x", has)
			//filename := md5str1 + ".png"//img[strings.Index(img,"m/") + 2:]
			//f,_ := os.Create("H:\\Test\\" + filename)
			//io.Copy(f,data.Body)
		//}
		if strings.Compare(next,"") != 1 {
			break
		}
		url = next
	}
	return data
}
func get_img(url string)(string,string){
	resp,_ := http.Get(url)
	bytes,_ := ioutil.ReadAll(resp.Body)
	html := string(bytes)

	img := "http://p.iktupian.com" + find_string(html,"src=\"http://p.iktupian.com","\"")
	//rrr,_ := regexp.Compile("[0-9]{1,3}/[0-9]{1,3}")
	ss := strings.Split(find_string(html,"page_cat\"><span>","</span>"),"/")
	i,_:=strconv.Atoi(ss[0])
	ii,_:=strconv.Atoi(ss[1])
	next := ""
	if i < ii {
		next = "http://iktupian.com/"  + find_string(html,"id=\"123\" href=\"","\"")
	}
	return img,next
}
func find_string(src string,start string,end string)string{
	s := strings.Index(src,start) + len(start)
	ss := src[s:]
	e := s + strings.Index(ss,end)
	if s < e {
		ss = src[s:e]
		return ss
	}
	return ""
	r1,err :=regexp.Compile(start + ".*?" + end)
	if err != nil{
		fmt.Println(err)
		return ""
	}
	data := r1.FindString(src)
	data = strings.Replace(data,start,"",1)
	data = strings.Replace(data,end,"",1)
	return data
}