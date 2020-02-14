package main

import (
	"net/http"
	"io"
	"fmt"
	"os"
	"io/ioutil"
	"flag"
)

var html_file_path string

func webcopyindex(response http.ResponseWriter, request *http.Request) {
	file, err := os.Open(html_file_path)
	defer file.Close()
	if err != nil {
		io.WriteString(response, err.Error())
		return
	}
	content, err := ioutil.ReadAll(file)
	if err != nil {
		io.WriteString(response, err.Error())
		return
	}
	io.WriteString(response, string(content))
}


func main() {
	filepath := flag.String("f", "mytest/webcopy/index.html", "指定渲染的HTML 文件")
	port := flag.String("p", "10000", "服务运行的指定端口")
	flag.Parse()
	
	html_file_path = *filepath
	http.HandleFunc("/", webcopyindex)
	fmt.Println("图片 拷贝-粘贴-展示 服务运行中 http://localhost:" + *port + "/ ")
	http.ListenAndServe(":" + *port, nil)
}