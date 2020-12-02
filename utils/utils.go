package utils

import (
	// "encoding/json"
	"fmt"
	"bufio"
	"log"
	"os"
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

func MaxInt(a int, b int) int {
	if a < b {
		return b
	}
	return a
}

func ReadLineSlice(path string) []string {
	fp, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer fp.Close()

	var lines []string
	scanner := bufio.NewScanner(fp)
	buf := make([]byte, 0, 64*1024) // MaxScanTokenSize 64 * 1024 bytes
	scanner.Buffer(buf, 1024*1024)  // upto 1MB
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if err = scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return lines
}

// func LoadJSON(path string, obj interface{}) {
// 	bs, err := ioutil.ReadFile(path)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	json.Unmarshal(bs, &obj)
// }

// func DirWalk(dir string, fn func(string, os.FileInfo) string) []string {
//     files, err := ioutil.ReadDir(dir)
//     if err != nil {
//         log.Fatal(err)
//     }

//     var paths []string
//     for _, file := range files {
// 		path := fn(dir, file)
// 		if path != "" {
// 			paths = append(paths, path)
// 		}
//     }
//     return paths
// }
