package dashboard

import (
	"fmt"
	// "log"
	"os"
	"github.com/mijies/dashboard_generator/internal/account"
	"github.com/mijies/dashboard_generator/internal/config"
	"github.com/mijies/dashboard_generator/pkg/utils"
)

type commands struct {
	chains		[]command
	user_chains	[]command
}

type command struct {
	Index	int					`json:"index"`
	Name	string				`json:"name"`
	Chain	[]string			`json:"chain"`
	Args	map[string]string	`json:"args"`
}

func(c *commands) getLength() int {
	return len((*c).chains) + len((*c).user_chains)
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

func(c *commands) ComponentIterator() []string {
	return []string{""}
}