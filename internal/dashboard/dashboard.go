package dashboard

import (
	// "errors"
	"fmt"
	"log"
	"github.com/360EntSecGroup-Skylar/excelize"
	"github.com/mijies/dashboard_builder/pkg/utils"
)

type dbook interface {
	load()
	parse(mtx chan bool, sch chan *[]snippet, cch chan *[]command)
	build()
}

// commands and snippets
type component interface {
	len() int
	load(book *excelize.File)
	parse() interface{}
	into_iter()
}

type dashboard struct {
	path		string
	book		*excelize.File
	commands	commands
	snippets	snippets
}
type targetBook struct {
	dashboard
}
type masterBook struct {
	dashboard
}
type userBook struct {
	dashboard
}

func(d *dashboard) load() {
	file, err := excelize.OpenFile(d.path)
    if err != nil {
        log.Fatal(err)
	}
	d.book = file
}
func(d *targetBook) load() {
	d.dashboard.load()
}
func(d *masterBook) load() {
	d.dashboard.load()
	d.commands.load(d.book)
	d.snippets.load(d.book)
	d.book = nil
}
func(d *userBook) load() {
	d.dashboard.load()
	d.commands.load(d.book)
	d.snippets.load(d.book)
	d.book = nil
}

func(d *dashboard) parse(mtx chan bool, sch chan *[]snippet, cch chan *[]command) {
}
func(d *targetBook) parse(mtx chan bool, sch chan *[]snippet, cch chan *[]command) {
	d._parse_snippets(*<- sch, *<- sch)
	d._parse_commands(*<- cch, *<- cch)
}
func(d *masterBook) parse(mtx chan bool, sch chan *[]snippet, cch chan *[]command) {
	sch <- d.snippets.parse().(*[]snippet)
	mtx <- true
	cch <- d.commands.parse().(*[]command)
	mtx <- true
}
func(d *userBook) parse(mtx chan bool, sch chan *[]snippet, cch chan *[]command) {
	<- mtx // wait for masterBook snippets
	sch <- d.snippets.parse().(*[]snippet)
	<- mtx
	cch <- d.commands.parse().(*[]command)
}

func(d *targetBook) _parse_snippets(ms []snippet, us []snippet) {
	for _, master := range ms {
		for _, user := range us {
			for k, _ := range master.snipMap {
				if _, dup := user.snipMap[k]; dup {
					log.Fatal("code name duplication with your custom code: " + k)
				}
			}
		}
	}
	ms = append(ms, us...)
	d.snippets.snippets = append(d.snippets.snippets, ms...)
}

func(d *targetBook) _parse_commands(mc []command, uc []command) {
	var c command
	m, u, offset := 0, 0, 0
	// prioritize user chains indces
	for i := 1; i <= len(mc) + len(uc); i++ {
		if m < len(mc) {
			c = mc[m]
		}
		if u < len(uc) && uc[u].index <= c.index + offset {
			c = uc[u]
			offset++
			u++
		} else {
			m++
		}
		c.index = i
		d.commands.chains = append(d.commands.chains, c)
	}
	for _, cmd := range d.commands.chains {
		fmt.Printf("%d %s\n", cmd.index, cmd.name)
	}
}

func(d *dashboard) build() {
}
func(d *targetBook) build() {
	// create a new book
	new_path := utils.AddTimestampToFilename(d.path, getTimeFormat(), "xlsm")
	if err := d.book.SaveAs(new_path); err != nil {
        log.Fatal(err)
	}

// 	// swap with the new one
// 	d.base_path = new_path
// 	file, err := excelize.OpenFile(d.base_path)
//     if err != nil {
//         log.Fatal(err)
// 	}
// 	d.book = file

// 	// delete the exsisting macro sheet
// 	sheet_name := d.cfg.GetMacroSheetName()
// 	d.book.DeleteSheet(sheet_name)

// 	// copy the template sheet
// 	d.book.NewSheet(sheet_name)
// 	tmp_index := d.book.GetSheetIndex(d.cfg.GetMacroTmpSheetName())
// 	index  := d.book.GetSheetIndex(sheet_name)
// 	if err := d.book.CopySheet(tmp_index, index); err != nil {
//         log.Fatal(err)
}

// 	d.renderSheet(sheet_name, &d.commands)
// 	d.renderSheet(sheet_name, &d.snippets)
// 	if err := d.book.Save(); err != nil {
//         log.Fatal(err)
// 	}
// }

// func(d *dashboard) renderSheet(sheet_name string, comp component) {
// 	label := comp.getComponentLabel(d.cfg)
// 	rowc  := d.commands.getLength() + 10 // row count to cover
// 	rows  := [2]int{1, rowc}
// 	cols  := [2]int{1, 5}
// 	rowi, err := d.findRow(sheet_name, label, rows, cols)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	seed := int('A')
// 	itr  := comp.iterable()
// 	for itr.hasNext() {
// 		rowi++
// 		d.book.DuplicateRow(sheet_name, rowi)
// 		cols, stys := itr.next()
// 		for i, v := range cols {
// 			axis := fmt.Sprintf("%s%d", string(seed + i), rowi)
// 			d.book.SetCellValue(sheet_name, axis, v)
// 			if len(stys[i]) != 0 {
// 				d.setRowStyle(sheet_name, axis, stys[i])
// 			}
// 		}
// 	}
// }

// }

// func(d *dashboard) setRowStyle(sheet_name string, axis string, style_str string) {
// 	style, err := d.book.NewStyle(style_str)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	err = d.book.SetCellStyle(sheet_name, axis, axis, style)
// }
