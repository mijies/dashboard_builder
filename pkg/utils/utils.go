package utils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"time"
)

func GetTimeStr(format string) string {
	return time.Now().Format(format)
}

func AddTimestampToFilename(path string, format string, extension string) string {
	timestamp := GetTimeStr(format)
	length := len(path) - len(extension) - 1 // -1 for a dot
	return fmt.Sprintf("%s_%s.%s", path[:length], timestamp, extension)
}

func LoadJSON(path string, obj interface{}) {
	bs, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}
	json.Unmarshal(bs, &obj)
}
