package services

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"golang.org/x/net/html/charset"
	"io"
	"strconv"

	//"github.com/PuerkitoBio/goquery"
)


func initGoquery (reader io.Reader) (*goquery.Document, error) {
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

func ParsePage (reader io.Reader, req *JsonRequest) map[string]interface{} {
	query, err := initGoquery(reader)

	var output = make(map[string]interface{})

	if err != nil {
		return output
	}

	for _, e := range req.Selectors {
		elements := getElements(query, e)

		if e.Repeated == false {
			output[e.Identifer] = parseSingle(elements, e)
		} else {
			output[e.Identifer] = parseMultiple(elements, e)
		}

	}
	return output
}

func getElements (query *goquery.Document, req Selector) []string {
	// 取出元素內容
	output := make([]string, 0)
	query.Find(req.Selector).Each(func (idx int, selection *goquery.Selection) {
		var html string
		switch req.Property {
		case "html":
			html, _ = selection.Html()
		case "text":
			html = selection.Text()
		case "attr":
			if req.Target == "" {
				panic(fmt.Sprintf("Try to retrive the attr of %s, but no attr target specified.", req.Selector))
			}
			html, _ = selection.Attr(req.Target)
		}
		output = append(output, html)
	})

	return output
}
func parseSingle (data []string, req Selector) interface{} {
	if data[0] == "" || len(data) <= 0 {
		return ""
	}
	switch req.Type {
	case "string":
		return data[0]
	case "integer":
		output, err := strconv.Atoi(data[0])
		if err == nil {
			return output
		}
		return -1
	case "boolean":
		output, err := strconv.ParseBool(data[0])
		if err == nil {
			return output
		}
		return false
	}
	return ""
}

func parseMultiple (data []string, req Selector) []interface{} {
	if len(data) <= 0 {
		return []interface{}{}
	}
	var output = make([]interface{}, len(data))
	for i, d := range data {
		switch req.Type {
		case "string":
			output[i] = d
			break
		case "integer":
			c, _ := strconv.Atoi(d)
			output[i] = c
			break
		case "boolean":
			c, _ := strconv.ParseBool(d)
			output[i] = c
			break
		}
	}
	return output
}