package services

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"golang.org/x/net/html/charset"
	"io"
	//"github.com/PuerkitoBio/goquery"
)


func initGoquery (reader io.Reader) (*goquery.Document, error) {
	utfReader, _ := charset.NewReader(reader, "text/html")
	query, err := goquery.NewDocumentFromReader(utfReader)

	if err != nil {
		return query, err
	}

	return query, nil
}

func Parse (reader io.Reader, req *JsonRequest) string {
	_, err := initGoquery(reader)

	if err != nil {
		return ""
	}

	//var output map[string]interface{}

	for i, e := range req.Selectors {
		fmt.Println(i, ":::", e.Selector)
	}
	return ""
}