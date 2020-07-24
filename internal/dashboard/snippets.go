package dashboard

import (
	"fmt"
	// "io/ioutil"
	"log"
	// "os"
	// "regexp"
	// "path/filepath"
	// "github.com/mijies/dashboard_builder/pkg/utils"
	"github.com/360EntSecGroup-Skylar/excelize"
)

type snippets struct {
	snippets		[]snippet
	// finalized		[]snippet
	// rows			[][]string // made by intoRows()
	// styles			[][]string // cell styles
}

type snippet struct {
	snipMap		map[string]string
}

// func(t *snippets) iterable() iterator {
// 	i := iter{
// 		index:	0,
// 		length:	t.getLength(),
// 		values:	&t.rows,
// 		styles:	&t.styles,
// 	}
// 	return iterator(&i)
// }

func(t *snippets) len() int {
	return len(t.snippets)
}

func(s *snippets) load(book *excelize.File) {
	// find label row
	rowi, err := findRow(book, MACRO_SHEET_NAME, SNIPPETS_LABEL, [2]int{1, 100}, [2]int{1, 5})
	if err != nil {
		log.Fatal(err)
	}

	// iterate unless No. column is empty
	var key string
	empty_row := 0
	for {
		rowi++
		axis := fmt.Sprintf("%s%d", "C", rowi)
		value, err := book.GetCellValue(MACRO_SHEET_NAME, axis)
		if err != nil {
			log.Fatal(err)
		}

		if value == "" {
			if empty_row > 1 { // break if 3 continuous rows are empty
				break
			}
			empty_row++
			continue
		}

		if value[0] == '[' {
			key = value
			continue
		}

		snip := snippet{snipMap: map[string]string{key: value}}
		s.snippets = append(s.snippets, snip)
		empty_row = 0
	}
	fmt.Printf("%d\n", len(s.snippets))

// 	snippetsFromDir(&t.snippets, base_dir)

// 	if _, err := os.Stat(user_dir); os.IsNotExist(err) {
// 		return
// 	}
// 	snippetsFromDir(&t.user_snippets, user_dir)
}

// func(t *snippets) finalize() {
// 	// title(file name) duplication not allowed
// 	for _, s := range t.snippets {
// 		for _, u := range t.user_snippets {
// 			for sk, _ := range s.snipMap {
// 				for uk, _ := range u.snipMap {
// 					if sk == uk {
// 						log.Fatal("code name duplication with your custom code: " + sk)
// 					}
// 				}
// 			}
// 		}
// 	}

// 	t.finalized = append(t.snippets, t.user_snippets...)
// 	t.snippets 		= nil
// 	t.user_snippets = nil
// 	t.intoRows()
// }

// func(t *snippets) intoRows() {
// 	for _, s := range t.finalized {
// 		for k, v := range s.snipMap {
// 			t.rows = append(t.rows, []string{"", "", "[" + k + "]"}) // 1st and 2nd columns are empty
// 			t.rows = append(t.rows, []string{"", "", string(v)})
// 			t.rows = append(t.rows, []string{"", "", ""})
// 			t.styles = append(t.styles, []string{"", "", STYLE_TITLE})
// 			t.styles = append(t.styles, []string{"", "", ""})
// 			t.styles = append(t.styles, []string{"", "", ""})
// 		}
// 	}
// 	t.finalized	= nil
// }

// func snippetsFromDir(snips *[]snippet, dir string) {
// 	file_names := utils.DirWalk(dir, onlyTextFile)
// 	for _, name := range file_names {
// 		bs, err := ioutil.ReadFile(filepath.Join(dir, name))
// 		if err != nil {
// 			log.Fatal(err)
// 		}
// 		snip := snippet{
// 			snipMap: map[string][]byte{name[:len(name)-4]: bs}, // remove .txt from name
// 		}
// 		*snips = append(*snips, snip)
// 	}
// }

// // used for utils.DirWalk
// func onlyTextFile(dir string, file os.FileInfo) string {
// 	r := regexp.MustCompile(`txt$`)
// 	if file.IsDir() || !r.MatchString(file.Name()) {
// 		return ""
// 	}
// 	return file.Name()
// }

// const (
// 	STYLE_TITLE = `{"border":[{"type":"left","style":1},{"type":"right","style":1},{"type":"top","style":1},{"type":"bottom","style":1}],
// 					"fill":{"type":"gradient","color":["#FFFFFF","#FFE6E6"],"shading":5}}`
// )