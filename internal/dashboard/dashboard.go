package dashboard

import (
	// "fmt"
	"log"
	"github.com/360EntSecGroup-Skylar/excelize"

	uac "github.com/mijies/dashboard_generator/internal/account"
	"github.com/mijies/dashboard_generator/internal/config"
)

func BuildDashboard(book_path string, account * uac.UserAccount) {
	d := dashboard{
		base_path:	book_path,
		account:	account,
		cfg:		config.NewConfig(),
	}
	d.initComponents()
	d.loadData()
	d.generate()
}

type dashboard_component interface {
	loadData()
	render(book *excelize.File, sheet_name string)
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

func(d *dashboard) loadData() {
	file, err := excelize.OpenFile(d.base_path)
    if err != nil {
        log.Fatal(err)
	}
	d.book = file

	d.commands.loadData()
}

func(d *dashboard) generate() {
	sheet_name := d.cfg.GetNewSheetName()
	d.book.NewSheet(sheet_name)
	d.render(d.book, sheet_name)
	// if err := d.book.Save(); err != nil {
    //     log.Fatal(err)
    // }
}

func(d *dashboard) render(book *excelize.File, sheet_name string) {
	d.commands.render(book, sheet_name)
}
