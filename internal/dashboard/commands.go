package dashboard

import (
	"fmt"
	// "log"
	"github.com/360EntSecGroup-Skylar/excelize"

	"github.com/mijies/dashboard_generator/pkg/utils"
)

type commands struct {
	base_path	string
	user_path	string
	chains		[]command
	user_chains	[]command
}

type command struct {
	Chain	[]string			`json:"chain"`
	Args	map[string]string	`json:"args"`
}

func(c *commands) loadData() {
	utils.LoadJSON(c.base_path, &c.chains)
	utils.LoadJSON(c.user_path + ".json", &c.user_chains)

	fmt.Printf(c.chains[0].Chain[0])
	fmt.Printf(c.chains[0].Args["hostname"])
	fmt.Printf(c.user_chains[0].Chain[0])
	fmt.Printf(c.user_chains[0].Args["hostname"])
}

func(c *commands) render(book *excelize.File, sheet_name string) {
	// book.SetCellValue(sheet_name, d.inv., "Hello world.")
}