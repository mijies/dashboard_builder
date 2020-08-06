package dashboard

import (
	"fmt"
	"log"
	"github.com/360EntSecGroup-Skylar/excelize"
	"github.com/mijies/dashboard_builder/account"
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

func(s *snippets) into_iter(acc *account.UserAccount) iterator {
	i := snippetsIterable{
		iterable: iterable{
			acc:	acc,
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
}

func(s *snippets) parse() interface{} {
	return &s.snippets
}
