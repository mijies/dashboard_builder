package dashboard

import (
	"errors"
	"fmt"
	"log"
	"github.com/360EntSecGroup-Skylar/excelize"

	uac "github.com/mijies/dashboard_generator/internal/account"
	"github.com/mijies/dashboard_generator/internal/config"
	"github.com/mijies/dashboard_generator/pkg/utils"
)

func BuildDashboard(book_path string, account * uac.UserAccount) {
	d := dashboard{
		base_path:	book_path,
		account:	account,
		cfg:		config.NewConfig(),
	}
	d.initComponents()
	d.loadComponents()
	d.buildDashboard()
}

// commands and ttl_codes
type dashboard_component interface {
	loadData()
	getComponentLabel(cfg config.Config) string
	// insertRows()
}

type dashboard struct {
	base_path	string
	cfg			config.Config
	book		*excelize.File
	account		*uac.UserAccount
	commands	commands
}

func(d *dashboard) initComponents() {
	d.commands = commands{
		base_path:  d.cfg.GetCommandsDir() + d.cfg.GetCommandsFile(),
		user_path:  d.cfg.GetCommandsDir() + d.account.Name,
	}
}

func(d *dashboard) loadComponents() {
	file, err := excelize.OpenFile(d.base_path)
    if err != nil {
        log.Fatal(err)
	}
	d.book = file

	d.commands.loadData()
}

func(d *dashboard) buildDashboard() {
	// create a new book
	time_format := d.cfg.GetTimeFormat()
	new_path := utils.AddTimestampToFilename(d.base_path, time_format, "xlsm")
	if err := d.book.SaveAs(new_path); err != nil {
        log.Fatal(err)
	}

	// reopen a book with the new one
	d.base_path = new_path
	file, err := excelize.OpenFile(d.base_path)
    if err != nil {
        log.Fatal(err)
	}
	d.book = file

	// delete exsisting the macro sheet
	sheet_name := d.cfg.GetMacroSheetName()
	d.book.DeleteSheet(sheet_name)

	// copy the template sheet
	d.book.NewSheet(sheet_name)
	tmp_index := d.book.GetSheetIndex(d.cfg.GetMacroTmpSheetName())
	index  := d.book.GetSheetIndex(sheet_name)
	if err := d.book.CopySheet(tmp_index, index); err != nil {
        log.Fatal(err)
	}

	d.renderSheet(sheet_name, &d.commands)
	// d.renderSheet(sheet_name, d.ttl_codes)
	if err := d.book.Save(); err != nil {
        log.Fatal(err)
	}
}

func(d *dashboard) renderSheet(sheet_name string, comp dashboard_component) {
	label := comp.getComponentLabel(d.cfg)
	rows := [2]int{1, 5}
	cols := [2]int{1, 5}
	axis, err := d.locateCell(sheet_name, label, rows, cols)
	if err != nil {
		log.Fatal(err)
	}		
	fmt.Printf("%s\n", axis)
	// comp.insertRows()
}

func(d *dashboard) locateCell(sheet_name string, value string, rows [2]int, cols [2]int) (string, error) {
	if cols[1] > 27 {
		log.Fatal("column only supported upto Z")
	}

	seed := int('A' - 1)
	for col := cols[0]; col < cols[1]; col++ {
		for row := rows[0]; row < rows[1]; row++ {
			axis := fmt.Sprintf("%s%d", string(seed + col), row)
			val, err := d.book.GetCellValue(sheet_name, axis)
			if err != nil {
				log.Fatal(err)
			}		
			if val == value {
				return axis, nil
			}
		}	
	} 
	return "", errors.New("no cell found with the value:" + value)
}
