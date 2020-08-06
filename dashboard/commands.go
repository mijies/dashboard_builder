package dashboard

import (
	"fmt"
	"log"
	"sort"
	"strconv"
	"strings"
	"github.com/360EntSecGroup-Skylar/excelize"
)

type commands struct {
	chains		[]command
}

type command struct {
	index	int
	name	string
	chain	[]string
	args	map[string]string
}

func(c *commands) len() int {
	return len(c.chains)
}

func(c *commands) into_iter() iterator {
	i := commandsIterable{
		iterable: iterable{
			index: 0,
			length: c.len(),
		},
		component: c,
	}
	return iterator(&i)
}

func(c *commands) load(book *excelize.File) {
	// find label row
	rowi, err := findRow(book, MACRO_SHEET_NAME, COMMANDS_LABEL, [2]int{1, 100}, [2]int{1, 5})
	if err != nil {
		log.Fatal(err)
	}

	// iterate unless No. column is empty
	for {
		rowi++
		axis := fmt.Sprintf("%s%d", "A", rowi)
		index, err := book.GetCellValue(MACRO_SHEET_NAME, axis)
		if err != nil {
			log.Fatal(err)
		}
		if index == "" {
			break
		}
		int_idx, err := strconv.Atoi(index)
		if err != nil {
			log.Fatal(err)
		}
		cmd := command{index: int_idx, args: make(map[string]string)}
		cmd._load(book, rowi)
		c.chains = append(c.chains, cmd)
	}
	fmt.Printf("%d\n", len(c.chains))
}

func(c *command) _load(book *excelize.File, rowi int) {
	name, err := book.GetCellValue(MACRO_SHEET_NAME, fmt.Sprintf("%s%d", "B", rowi))
	if err != nil {
		log.Fatal(err)
	}	
	if name == "" {
		return
	}
	c.name = name

	chain, err := book.GetCellValue(MACRO_SHEET_NAME, fmt.Sprintf("%s%d", "C", rowi))
	if err != nil {
		log.Fatal(err)
	}	
	if chain == "" {
		return
	}
	c.chain = strings.Split(chain, ",")

	col := int('D')
	for i := 0;; i++ {
		args, err := book.GetCellValue(MACRO_SHEET_NAME, fmt.Sprintf("%s%d", string(col + i), rowi))
		if err != nil {
			log.Fatal(err)
		}
		if args == "" {
			return
		}
		kv := strings.Split(args, ",")
		c.args[kv[0]] = kv[1]
	}
}

func(c *commands) parse() interface{} {
	sort.SliceStable(c.chains, func(a, b int) bool { return c.chains[a].index < c.chains[b].index })
	return &c.chains
}
