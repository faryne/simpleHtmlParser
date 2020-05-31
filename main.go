package main

import (
	"./services"
	"flag"
	"golang.org/x/net/html/charset"

	"fmt"
	"os"
)

var (
	uri string
	file string
)

func main () {
	// 解析參數
	flag.StringVar(&uri, "uri", "", "要爬取的網址")
	flag.StringVar(&file, "file", "", "爬取的設定檔")
	flag.Parse()

	enc, label := charset.Lookup("utf-8")
	fmt.Println(enc)
	fmt.Println(label)
	//return

	// 檢查參數
	if len(uri) <= 0 {
		services.GenerateError(001, "uri should not be empty")
		return
	}
	if len(file) <= 0 {
		services.GenerateError(002, "file path should not be empty")
		return
	}
	if _, errFile := os.Stat(file); os.IsNotExist(errFile) {
		services.GenerateError(003, "file path is not valid")
		return
	}

	// 讀取 html
	reader, errHTML := services.GetHTMLResponse(uri)
	if errHTML != nil {
		services.GenerateError(004, errHTML.Error())
		return
	}

	// 解析 json 檔內容
	req, errReq := services.InitRequest(file)
	if errReq != nil {
		services.GenerateError(005, errReq.Error())
		return
	}

	// 開始進入解析步驟
	query := services.Parse(reader, req)
	//fmt.Printf("%+v", resp)
	fmt.Println(query)
}