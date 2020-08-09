package services

import (
	"encoding/csv"
	"io"
	"strconv"
)

func CsvParse (reader io.Reader, req *JsonRequest) []map[string]interface{} {
	var output = make([]map[string]interface{}, 0)

	csvReader := csv.NewReader(reader)
	for {
		c, _ := csvReader.Read()
		if len(c) == 0 {
			break
		}
		var tmp = make(map[string]interface{})

		for i := range req.Selectors {
			selectorPath, _ := strconv.Atoi(req.Selectors[i].Selector)
			ele := &ElementOutput{
				Converted: 	convData(c[selectorPath], req.Selectors[i].Output.Type),
				Raw:       	c[selectorPath],
				RegExp:    	regexpConverted(c[selectorPath], CollRegexp[req.Selectors[i].Output.Regexp]),
			}
			tmp[req.Selectors[i].Identifer]= ele
		}
		output = append(output, tmp)
	}

	return output
}
