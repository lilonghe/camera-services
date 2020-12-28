package utils

import "encoding/json"

func ToJSON(obj interface{}) string {
	b, _ := json.Marshal(obj)
	return string(b)
}
