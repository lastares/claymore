package jsonutil

import (
	"encoding/json"
)

func JsonEncode(s any) ([]byte, error) {
	return json.Marshal(s)
}

func JsonDecode(data []byte, v any) error {
	return json.Unmarshal(data, &v)
}
