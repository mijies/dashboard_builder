package dashboard

import (
	"fmt"
	"os"
	"sort"
	"strings"
	"github.com/mijies/dashboard_builder/internal/account"
	"github.com/mijies/dashboard_builder/internal/config"
	"github.com/mijies/dashboard_builder/pkg/utils"
)

type commands struct {
	chains		[]command
	user_chains	[]command
	finalized	[]command
	rows		[][]string // made by intoRows()
	styles		[][]string // cell styles
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
		values:	&c.rows,
		styles:	&c.styles,
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
		style := []string{
			STYLE_NO, STYLE_NAME, STYLE_CHAIN,
		}

		for k, v := range cmd.Args {
			row   = append(row, k + "," + v)
			style = append(style, STYLE_ARGS)
		}
		c.rows   = append(c.rows,   row)
		c.styles = append(c.styles, style)
	}
	c.finalized	= nil
}

const (
	STYLE_NO    = ``
	STYLE_NAME  = `{"font":{"bold":true}}`
	STYLE_CHAIN = `{"border":[{"type":"left","style":1},{"type":"right","style":1},{"type":"top","style":1},{"type":"bottom","style":1}],
					"fill":{"type":"gradient","color":["#FFFFFF","#E0EBF5"],"shading":5}}`
	STYLE_ARGS  = ``
)