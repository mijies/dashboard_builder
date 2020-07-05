package inventory

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"time"
)

func getTimeStr(format string) string {
	return time.Now().Format(format)
}

func loadJSON(path string, obj interface{}) {
	bs, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}
	json.Unmarshal(bs, &obj)
}