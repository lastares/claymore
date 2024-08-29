package generalutil

import (
	"bytes"
	"encoding/json"
	"log"
)

func PrettyPrintStruct(obj interface{}) {
	jsonStr, _ := json.Marshal(obj)
	var buf bytes.Buffer
	json.Indent(&buf, jsonStr, "", "\t")
	log.Printf("pretty struct output=%v\n", buf.String())
}
