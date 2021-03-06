package services

import (
	"encoding/json"
	"os"
)

type JsonRequest struct {
	Encoding 		string 				`json:"encoding"`
	Regexp 			map[string]string	`json:"regexp"`
	Selectors 		[]Selector			`json:"selectors"`
}
type Selector struct {
	Identifer 		string				`json:"identifier"`
	Selector 		string 				`json:"selector"`
	Repeated 		bool 				`json:"repeated"`
	Children 		[]Selector 			`json:"children"`
	Output 			Element 			`json:"output"`
}
type Element struct {
	Property 		string 				`json:"property"`
	Target 			string 				`json:"target"`
	Type 			string 				`json:"type"`
	Regexp 			string 				`json:"regexp"`
}
type RegExp struct {
	Pattern 		string 				`json:"pattern"`
}

func getJsonContent (filename string) ([]byte, error){
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}

	stat, err1 := file.Stat()
	if err1 != nil {
		return nil, err1
	}

	content := make([]byte, stat.Size())
	file.Read(content)

	return content, nil
}

func InitRequest (filename string) (*JsonRequest, error) {
	content, err := getJsonContent(filename)

	var req = &JsonRequest{}

	if err != nil {
		return req, err
	}

	e := json.Unmarshal(content, req)
	if e != nil {
		panic("Cannot unmarshall json file: " + e.Error())
	}

	return req, nil
}
