package dashboard

import (
	"fmt"
	// "io/ioutil"
	"log"
	// "os"
	// "regexp"
	// "path/filepath"
	// "github.com/mijies/dashboard_builder/utils"
	"github.com/360EntSecGroup-Skylar/excelize"
)

type snippets struct {
	snippets		[]snippet
}

type snippet struct {
	snipMap		map[string]string
}

func(s *snippets) len() int {
	return len(s.snippets)
}

func(s *snippets) into_iter() iterator {
	i := snippetsIterable{
		iterable: iterable{
			index:	0,
			length:	s.len() * 3, // a snippet takes 3 rows
		},
		component: s,
	}
	return iterator(&i)
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
	// fmt.Printf("%d\n", len(s.snippets))

// 	snippetsFromDir(&t.snippets, base_dir)

// 	if _, err := os.Stat(user_dir); os.IsNotExist(err) {
// 		return
// 	}
// 	snippetsFromDir(&t.user_snippets, user_dir)
}

func(s *snippets) parse() interface{} {
	return &s.snippets
}


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
