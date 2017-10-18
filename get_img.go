package main

import (
	"net/http"
	"io/ioutil"
	"fmt"
	"regexp"
	"strings"
	"strconv"
)

func main() {
	//url := "http://iktupian.com/wuyiniao/"
	//resp,_ := http.Get(url)
	//src,_ := ioutil.ReadAll(resp.Body)
	//reg,_ := regexp.Compile("(show_box).*?(</div>)")
	//reg2,_ := regexp.Compile("(=\").*?(\")")
	//items := reg.FindAllString(string(src),-1)
	//for _,i := range items  {
	//	datas := reg2.FindAllString(string(i),5)
	//	fmt.Println(datas[1][2:])
	//	fmt.Println(datas[3])
	//	fmt.Println(datas[4])
	//	fmt.Println()
	//}
	get_list("")
}
func get_list(url string){
	url = "http://iktupian.com/shaonv/16646p"
	i := 1

	for i == 1  {
		img,next := get_img(url)
		fmt.Println(img)
		//data,err := http.Get(img)
		//if err == nil {
		//	has := md5.Sum([]byte(img))
		//	md5str1 := fmt.Sprintf("%x", has)
		//	filename := md5str1 + ".png"//img[strings.Index(img,"m/") + 2:]
		//	f,_ := os.Create("H:\\Test\\" + filename)
		//	io.Copy(f,data.Body)
		//}
		if strings.Compare(next,"") != 1 {
			break
		}
		url = next
	}
}
func get_img(url string)(string,string){
	resp,_ := http.Get(url)
	bytes,_ := ioutil.ReadAll(resp.Body)
	html := string(bytes)

	img := find_string(html,"target=\"_blank\"><img src=\"","\"")
	//rrr,_ := regexp.Compile("[0-9]{1,3}/[0-9]{1,3}")
	ss := strings.Split(find_string(html,"page_cat\"><span>","</span>"),"/")
	fmt.Println(ss,len(ss))
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