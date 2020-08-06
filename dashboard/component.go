package dashboard

import (
	"fmt"
	"strings"
	"github.com/360EntSecGroup-Skylar/excelize"
)

type iterator interface {
    hasNext()	bool
    next()		([]string, []string)
}

// commands and snippets
type component interface {
	iterator
	len() int
	load(book *excelize.File)
	parse() interface{}
	into_iter()
}

type iterable struct {
	index	int
	length	int
}
type commandsIterable struct {
	iterable
	component	*commands
}
type snippetsIterable struct {
	iterable
	component	*snippets
}

func (i *iterable) hasNext() bool {
    if i.index < i.length {
        return true
    }
    return false
}

func(i *commandsIterable) next() ([]string, []string) {
	styles := COMMANDS_STYLE[:]
	item   := &(*i.component).chains[i.index]
	cols   := []string{
		fmt.Sprintf("%d", item.index),
		item.name,
		strings.Join(item.chain, ","),
	}
	for k, v := range item.args {
		cols   = append(cols, k + "," + v)
		styles = append(styles, COMMANDS_STYLE_ARGS)
	}
	i.index++
	return cols, styles
}

func(i *snippetsIterable) next() ([]string, []string) {
	styles := SNIPPETS_STYLE_BLANKS[:]
	cols   := []string{"","",""}
	// a snippet takes 3 rows
	switch(i.index % 3) {
	case 0: // title
		item := &(*i.component).snippets[i.index / 3]
		for k, _ := range item.snipMap {
			cols[2] = k
		}
		styles = SNIPPETS_STYLE_NAMES[:]
	case 1: // code
		item := &(*i.component).snippets[i.index / 3]
		for _, v := range item.snipMap {
			cols[2] = string(v)
		}
	case 2: // blank
	}
	i.index++
	return cols, styles
}
