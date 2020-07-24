package dashboard

import (
	"fmt"
	"log"
	"sort"
	"strconv"
	"strings"
	"github.com/360EntSecGroup-Skylar/excelize"
	// "github.com/mijies/dashboard_builder/pkg/utils"
)

type commands struct {
	chains		[]command
	// styles		[][]string // cell styles
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
		items: c,
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
		cmd.load(book, rowi)
		c.chains = append(c.chains, cmd)
	}
	fmt.Printf("%d\n", len(c.chains))
}

func(c *commands) parse() interface{} {
	sort.SliceStable(c.chains, func(a, b int) bool { return c.chains[a].index < c.chains[b].index })
	return &c.chains
}

func(c *command) load(book *excelize.File, rowi int) {
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

// func(c *commands) finalize() {
// 	c.finalized   = append(c.user_chains, c.chains...)
// 	c.chains 	  = nil
// 	c.user_chains = nil
// 	sort.SliceStable(c.finalized, func(a, b int) bool { return c.finalized[a].Index < c.finalized[b].Index })
// 	c.intoRows()
// }

// func(c *commands) intoRows() {
// 	for _, cmd := range c.finalized {
// 		row := []string{
// 			fmt.Sprintf("%d", cmd.Index),
// 			cmd.Name,
// 			strings.Join(cmd.Chain, ","),
// 		}
// 		style := []string{
// 			STYLE_NO, STYLE_NAME, STYLE_CHAIN,
// 		}

// 		for k, v := range cmd.Args {
// 			row   = append(row, k + "," + v)
// 			style = append(style, STYLE_ARGS)
// 		}
// 		c.rows   = append(c.rows,   row)
// 		c.styles = append(c.styles, style)
// 	}
// 	c.finalized	= nil
// }

// const (
// 	STYLE_NO    = ``
// 	STYLE_NAME  = `{"font":{"bold":true}}`
// 	STYLE_CHAIN = `{"border":[{"type":"left","style":1},{"type":"right","style":1},{"type":"top","style":1},{"type":"bottom","style":1}],
// 					"fill":{"type":"gradient","color":["#FFFFFF","#E0EBF5"],"shading":5}}`
// 	STYLE_ARGS  = ``
// )