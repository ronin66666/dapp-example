package util

import (
	"encoding/json"
)

var JsonPretty = true

func ParseToFile(data interface{}) []byte {
	var bytes []byte
	var e error
	if JsonPretty {
		bytes, e = json.MarshalIndent(data, "", "    ")
	} else {
		bytes, e = json.Marshal(data)

	}
	if e != nil {
		panic(e)
	}

	return bytes
}
