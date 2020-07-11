package dashboard

import (
	"fmt"
	// "log"
	"os"
	"strings"
	"github.com/mijies/dashboard_generator/internal/account"
	"github.com/mijies/dashboard_generator/internal/config"
	"github.com/mijies/dashboard_generator/pkg/utils"
)

type commands struct {
	chains			[]command
	user_chains		[]command
	parsed_chains	[]command // made by parseData()
}

type command struct {
	Index	int					`json:"index"`
	Name	string				`json:"name"`
	Chain	[]string			`json:"chain"`
	Args	map[string]string	`json:"args"`
}

type commandIterator struct {
	index	int
	length	int
	items	[][]string
}

func(c *commands) iterable() iterator {
	i := iter{
		index:	0,
		length:	c.getLength(),
		comp:	dashboard_component(c),
	}
	return iterator(&i)
}

func(c *commands) getLength() int {
	return len((*c).parsed_chains)
}

func(c *commands) getComponentLabel(cfg config.Config) string {
	return cfg.GetCommandsLabel()
}

func(c *commands) loadData(cfg config.Config, acc *account.UserAccount) {
	base_path := cfg.GetCommandsDir() + cfg.GetCommandsFile()
	user_path := cfg.GetCommandsDir() + acc.Name + ".json"

	utils.LoadJSON(base_path, &c.chains)

	if _, err := os.Stat(user_path); os.IsNotExist(err) {
		return
	}
	utils.LoadJSON(user_path, &c.user_chains)

	fmt.Printf("%s\n", c.user_chains[0].Chain[0])
	fmt.Printf("%s\n", c.user_chains[0].Args["hostname"])
}

func(c *commands) parseData() {
	c.parsed_chains = c.chains 
}

func(c *commands) intoRow(index int) [][]string {
	var rows [][]string
	cmd := c.parsed_chains[index]
	rows[0] = []string{
		string(cmd.Index),
		cmd.Name,
		strings.Join(cmd.Chain, ","),
	}
	for k, v := range cmd.Args {
		rows[0] = append(rows[0], k + "," + v)
	}
	return rows
}
