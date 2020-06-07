package main

import (
	"./services"
	"flag"
	"fmt"
	"os"
	"time"
)

var (
	uri string
	file string
)

type a struct {
	Msg string
	Code int
}

func main () {
	// 解析參數
	flag.StringVar(&uri, "uri", "", "要爬取的網址")
	flag.StringVar(&file, "file", "", "爬取的設定檔")
	flag.Parse()

	// 最頂層負責錯誤輸出的地方
	defer func () {
		if e := recover(); e != nil {
			if e.(services.CrawlerError).ErrorMessage == "" {
				services.SetError(000, fmt.Sprintf("%+v", e))
			}
			services.GenerateErrorOutput(e.(services.CrawlerError))
		}
	}()

	startTime := time.Now()
	// 檢查參數
	if len(uri) <= 0 {
		services.SetError(001, "uri should not be empty")
	}
	if len(file) <= 0 {
		services.SetError(002, "file path should not be empty")
	}
	if _, errFile := os.Stat(file); os.IsNotExist(errFile) {
		services.SetError(003, "file path is not valid")
	}

	// 讀取 html
	reader, errHTML := services.GetHTMLResponse(uri)
	if errHTML != nil {
		services.SetError(004, errHTML.Error())
	}
	// 不要在 function 內下 defer close
	defer reader.Close()

	// 解析 json 檔內容
	req, errReq := services.InitRequest(file)
	if errReq != nil {
		services.SetError(005, errReq.Error())
	}

	// 開始進入解析步驟
	convReader, errReader := services.InitGoquery(reader)
	if errReader != nil {
		services.SetError(006, errReader.Error())
	}
	output := services.ParsePage(convReader, req)
	endTime := time.Now()

	// 將爬蟲結果輸出為 json
	services.GenerateOutput(output, endTime.Sub(startTime).Seconds())
}
