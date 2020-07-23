package dashboard

import (
	// "errors"
	// "fmt"
	"log"
	"github.com/360EntSecGroup-Skylar/excelize"
	// "github.com/mijies/dashboard_builder/pkg/utils"
)

// commands and snippets
type component interface {
	iterable()	iterator
	len() int
	load(book *excelize.File)
	// finalize()
	// intoRows()
}

type book interface {
	load()
}

type dashboard struct {
	path	string
	book		*excelize.File
	commands	commands
	// snippets	snippets
}

func(d *dashboard) load() {
	file, err := excelize.OpenFile(d.path)
    if err != nil {
        log.Fatal(err)
	}
	d.book = file
	d.commands.load(d.book)
	// d.commands.finalize()
	// d.snippets.load()
	// d.snippets.finalize()
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
