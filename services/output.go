package services

import (
	"encoding/json"
	"fmt"
)

type CrawlerError struct {
	ErrorCode 		int 						`json:"code"`
	ErrorMessage	string						`json:"message"`
}

type CrawlerOutput struct {
	Code 			int							`json:"code"`
	SpentTime 		float64						`json:"spent_time"`
	Data 			interface{}					`json:"data"`
}

func GenerateError (code int, message string) {
	content := CrawlerError{
		ErrorCode:    code,
		ErrorMessage: message,
	}
	j, e := json.Marshal(content)
	if e != nil {
		content.ErrorCode = 0
		content.ErrorMessage = "Generate JSON Failed"
	}
	fmt.Println(string(j))
}

func GenerateOutput (output interface{}, spentTime float64) {
	content := &CrawlerOutput{
		Code:      200,
		SpentTime: spentTime,
		Data:      output,
	}

	j, e := json.Marshal(content)
	if e != nil {
		GenerateError(0, e.Error())
		return
	}
	fmt.Println(string(j))
}
