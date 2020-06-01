package services

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"golang.org/x/net/html/charset"
	"io"
	"strconv"

	//"github.com/PuerkitoBio/goquery"
)


func InitGoquery (reader io.Reader) (*goquery.Document, error) {
	utfReader, err0 := charset.NewReader(reader, "text/html")
	if err0 != nil {
		return nil, err0
	}
	query, err1 := goquery.NewDocumentFromReader(utfReader)

	if err1 != nil {
		return query, err1
	}

	return query, nil
}

func ParsePage (query *goquery.Document, req *JsonRequest) map[string]interface{} {
	var output = make(map[string]interface{})

	for _, e := range req.Selectors {
		elements := getElements(query, e)
		output[e.Identifer] = elements
	}
	return output
}

func clearData (selection *goquery.Selection, req Selector) interface{} {
	var html string
	switch req.Output.Target {
	case "html":
		html, _ = selection.Html()
	case "text":
		html = selection.Text()
	case "attr":
		if req.Output.Property == "" {
			panic(fmt.Sprintf("Try to retrive the attr of %s, but no attr target specified.", req.Selector))
		}
		html, _ = selection.Attr(req.Output.Property)
	}
	return convData(html, req.Type)
}
func convData (data string, dataType string) interface{} {
	switch dataType {
	case "string":
		return data
	case "integer":
		i, _ := strconv.Atoi(data)
		return i
	case "boolean":
		b, _ := strconv.ParseBool(data)
		return b
	}
	return nil
}
func getElements (query *goquery.Document, req Selector) interface{} {
	// 取出元素內容
	var output []interface{}
	query.Find(req.Selector).Each(func (idx int, selection *goquery.Selection) {
		if len(req.Children) > 0 {
			output = append(output, getChildrenElements(selection, req.Children))
		} else {
			output = append(output, clearData(selection, req))
		}
	})

	return output
}
func getChildrenElements (query *goquery.Selection, req []Selector) map[string]interface{} {
	var output = make(map[string]interface{})
	for _, r := range req {
		query.Find(r.Selector).Each(func(i int, selector *goquery.Selection){
			if len(r.Children) > 0 {
				output[r.Identifer] = getChildrenElements(selector, r.Children)
			} else {
				output[r.Identifer] = clearData(selector, r)
			}
		})
	}

	return output
}
//func parseSingle (data []string, req Selector) interface{} {
//	if data[0] == "" || len(data) <= 0 {
//		return ""
//	}
//	switch req.Type {
//	case "string":
//		return data[0]
//	case "integer":
//		output, err := strconv.Atoi(data[0])
//		if err == nil {
//			return output
//		}
//		return -1
//	case "boolean":
//		output, err := strconv.ParseBool(data[0])
//		if err == nil {
//			return output
//		}
//		return false
//	}
//	return ""
//}
//
//func parseMultiple (data []string, req Selector) []interface{} {
//	if len(data) <= 0 {
//		return []interface{}{}
//	}
//	var output = make([]interface{}, len(data))
//	for i, d := range data {
//		switch req.Type {
//		case "string":
//			output[i] = d
//			break
//		case "integer":
//			c, _ := strconv.Atoi(d)
//			output[i] = c
//			break
//		case "boolean":
//			c, _ := strconv.ParseBool(d)
//			output[i] = c
//			break
//		}
//	}
//	return output
//}