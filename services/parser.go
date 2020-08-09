package services

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"golang.org/x/net/html/charset"
	"io"
	"regexp"
	"strconv"

	//"github.com/PuerkitoBio/goquery"
)

var (
	CollRegexp 	map[string]string
)

func InitGoquery (reader io.Reader, enc string) (*goquery.Document, error) {
	utfReader, err0 := charset.NewReaderLabel(enc, reader)
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
		output[e.Identifer] = getElementContent(query.Find(e.Selector), e)
	}
	return output
}

func clearData (selection *goquery.Selection, req Selector) *ElementOutput {
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

	output := &ElementOutput{
		Converted: 	convData(html, req.Output.Type),
		Raw:       	html,
		RegExp:    	regexpConverted(html, CollRegexp[req.Output.Regexp]),
	}
	return output
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
func regexpConverted (data string, pattern string) interface{} {
	if pattern == "" {
		return false
	}
	re, e := regexp.Compile(`` + pattern + ``)
	if e != nil {
		panic("RegExp Pattern Error: " + e.Error())
	}
	if !re.MatchString(data) {
		return nil
	}

	return re.FindStringSubmatch(data)
}

func getChildrenElements (query *goquery.Selection, req []Selector) map[string]interface{} {
	var output = make(map[string]interface{})
	for _, r := range req {
		query.Find(r.Selector).Each(func(i int, selector *goquery.Selection) {
			if len(r.Children) > 0 {
				output[r.Identifer] = getChildrenElements(selector, r.Children)
			} else {
				output[r.Identifer] = clearData(selector, r)
			}
		})
	}

	return output
}

func getElementContent (selector *goquery.Selection, req Selector) interface{} {
	// 如果是單一元素時
	if req.Repeated == false {
		if len(req.Children) > 0 {
			return getChildrenElements(selector, req.Children)
		}
		return clearData(selector, req)
	}
	var output []interface{}
	selector.Each(func (idx int, selection *goquery.Selection) {
		if len(req.Children) > 0 {
			output = append(output, getChildrenElements(selection, req.Children))
		} else {
			output = append(output, clearData(selection, req))
		}
	})

	return output
}