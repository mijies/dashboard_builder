package dashboard

import (
	// "fmt"
	// "log"
	"github.com/360EntSecGroup-Skylar/excelize"

	uac "github.com/mijies/dashboard_generator/internal/account"
	inv "github.com/mijies/dashboard_generator/internal/inventory"
)

func BuildDashboard(book_path string, account * uac.UserAccount) {
	d := dashboard{
		account:	account,
		inventory:	inv.NewInventory(),
	}
	d.loadData(book_path)
}

type dashboard_component interface {
	loadData(base_path string)
	build()
	render(sheet_name string)
}

type dashboard struct {
	book		*excelize.File
	account		*uac.UserAccount
	inventory	inv.Inventory
	commands	commands
}

type commands struct {
	chains		[]command
	user_chains	[]command
}

type command struct {
	Chain	[]string			`json:"chain"`
	Args	map[string]string	`json:"args"`
}

func(d *dashboard) loadData(path string) {
// 	d.commands.loadData(BASE_PATH + COMMANDS_DIR, account)

// 	file, err := excelize.OpenFile(path)
//     if err != nil {
//         log.Fatal(err)
// 	}
// 	d.book = file
}

// func(c *commands) loadData(path string, account *user_account) {
// 	loadJSON(path + COMMANDS_FILE, 	&c.chains)
// 	loadJSON(path + account.name + ".json", &c.user_chains)

// 	fmt.Printf(c.chains[0].Chain[0])
// 	fmt.Printf(c.chains[0].Args["hostname"])
// 	fmt.Printf(c.user_chains[0].Chain[0])
// 	fmt.Printf(c.user_chains[0].Args["hostname"])
// }

// func(d *dashboard) build() {
// 	// d.commands.build()

// 	timestamp := getTimeStr(TIME_FORMAT)
// 	d.book.NewSheet(MACRO_SHEET_NAME_PREFIX + timestamp)
// 	d.render(timestamp)
// 	// if err := d.book.Save(); err != nil {
//     //     log.Fatal(err)
//     // }
// }

// func(d *dashboard) render(sheet_name string) {

// 	timestamp := getTimeStr(TIME_FORMAT)
// 	d.book.f.SetCellValue(sheet_name, "A2", "Hello world.")
// 	d.render(timestamp)
// 	// if err := d.book.Save(); err != nil {
//     //     log.Fatal(err)
//     // }
// }
