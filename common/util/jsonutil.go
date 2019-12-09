package util

import (
	"encoding/json"
)

func InterfaceToJsonString(v interface{}) string {
	str, err := json.Marshal(v)
	if err != nil {
		Error.Fatalln(err)
	}
	return string(str)
}

func JsonStringToInterface(str string, v interface{}) {
	err := json.Unmarshal([]byte(str), v)
	if err != nil {
		Error.Fatalln(err)
	}
}
