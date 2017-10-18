package main

import (
	"os"
	"fmt"
	"strings"
	"strconv"
	"io"
)

func main() {
	args := os.Args //获取用户输入的所有参数
	if args == nil || len(args) <2{
		fmt.Println("请填写正确的写入读取文件")
		return
	}
	w_file := args[1] //获取输入的第一个参数
	r_file := args[2]  //获取输入的第二个参数
	offset := args[3]  //获取输入的第二个参数

	if len(w_file) == 0 && len(r_file) == 0{

		return
	}
	//_file(&w_file)
	//_file(&r_file)
	write, error := os.OpenFile(w_file, os.O_RDWR, 0666)  //打开文件
	if error != nil {
		fmt.Println("写出文件")
		fmt.Println(error)
	}
	read,error := os.Open(r_file)
	if error != nil {
		fmt.Println("读取文件")
		fmt.Println(error)
	}
	a,_ := strconv.ParseInt(offset,10,16)
	write.Seek(a * 512,0)
	_,error = io.Copy(write,read)
	if error != nil {
		fmt.Println(error)
	}
	write.Close()
}
func _file(file* string){
	f := *file
	one := string(f[1])
	if strings.Compare(one,".") != -1 && strings.Compare(one,"/") != -1 {
		a := "./" + f
		file = &a
	}
}