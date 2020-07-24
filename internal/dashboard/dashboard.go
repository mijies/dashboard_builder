package dashboard

import (
	// "errors"
	// "fmt"
	"log"
	"github.com/360EntSecGroup-Skylar/excelize"
	// "github.com/mijies/dashboard_builder/pkg/utils"
)

type dbook interface {
	get_path() string
	ref_commands()	*commands
	ref_snippets()	*snippets
	load()
	parse(cmds *commands, snip *snippets, sch chan bool)
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

func(d *dashboard) get_path() string {
	return d.path
}

func(d *dashboard) ref_commands() *commands {
	return &d.commands
}
func(d *dashboard) ref_snippets() *snippets {
	return &d.snippets
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
	d.book = nil	// relase as book can be surely opened
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

func(d *dashboard) parse(cmds *commands, snip *snippets, sch chan bool) {
	// cs := d.commands.parse().(*commands)
}
func(d *masterBook) parse(cmds *commands, snip *snippets, sch chan bool) {
	ss := d.snippets.parse().(*[]snippet)
	snip.snippets = append(snip.snippets, *ss...)
	sch <- true
}
func(d *userBook) parse(cmds *commands, snip *snippets, sch chan bool) {
	ss := d.snippets.parse().(*[]snippet)
	<- sch // wait for masterBook
	snip.snippets = append(snip.snippets, *ss...)
}

// func(d *dashboard) build() {
// 	// create a new book
// 	time_format := d.cfg.GetTimeFormat()
// 	new_path := utils.AddTimestampToFilename(d.base_path, time_format, "xlsm")
// 	if err := d.book.SaveAs(new_path); err != nil {
//         log.Fatal(err)
// 	}

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
// 	}

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
