package dashboard

import (
	"fmt"
	// "log"
	"os"
	"sort"
	"strings"
	"github.com/mijies/dashboard_generator/internal/account"
	"github.com/mijies/dashboard_generator/internal/config"
	"github.com/mijies/dashboard_generator/pkg/utils"
)

type commands struct {
	chains		[]command
	user_chains	[]command
	finalized	[]command
	rows		[][]string // made by intoRows()
}

type command struct {
	Index	int					`json:"index"`
	Name	string				`json:"name"`
	Chain	[]string			`json:"chain"`
	Args	map[string]string	`json:"args"`
}

func(c *commands) iterable() iterator {
	i := iter{
		index:	0,
		length:	c.getLength(),
		items:	&c.rows,
	}
	return iterator(&i)
}

func(c *commands) getLength() int {
	if len(c.rows) != 0 {
		return len(c.rows)
	}
	if len(c.finalized) != 0 {
		return len(c.finalized)
	}
	return -1 // length is unknown until finalized
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
}

func(c *commands) finalize() {
	c.finalized   = append(c.user_chains, c.chains...)
	c.chains 	  = nil
	c.user_chains = nil
	sort.SliceStable(c.finalized, func(a, b int) bool { return c.finalized[a].Index < c.finalized[b].Index })
	c.intoRows()
}

func(c *commands) intoRows() {
	for _, cmd := range c.finalized {
		row := []string{
			fmt.Sprintf("%d", cmd.Index),
			cmd.Name,
			strings.Join(cmd.Chain, ","),
		}
		for k, v := range cmd.Args {
			row = append(row, k + "," + v)
		}
		c.rows = append(c.rows, row)
	}
	c.finalized	= nil
}