package utils

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"time"
)

func GetTimeStr(format string) string {
	return time.Now().Format(format)
}

func LoadJSON(path string, obj interface{}) {
	bs, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}
	json.Unmarshal(bs, &obj)
}