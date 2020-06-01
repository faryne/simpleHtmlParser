package services

import (
	"io"
	"net/http"
)

func GetHTMLResponse (uri string) (io.ReadCloser, error) {
	req, e := http.Get(uri)

	if e != nil {
		return req.Body, e
	}

	return req.Body, nil
}
