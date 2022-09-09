package utils

import "encoding/json"

func JsonToString(data interface{}) string {

	marshal, _ := json.Marshal(data)
	return string(marshal)
}
