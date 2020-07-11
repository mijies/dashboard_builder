package dashboard

import (
	// "fmt"
	"io/ioutil"
	"log"
	"os"
	"regexp"
	"path/filepath"
	"github.com/mijies/dashboard_generator/internal/account"
	"github.com/mijies/dashboard_generator/internal/config"
	"github.com/mijies/dashboard_generator/pkg/utils"
)

type ttl_codes struct {
	snippets		[]snippet
	user_snippets	[]snippet
	parsed_snippets []snippet
}

type snippet struct {
	snipMap		map[string][]byte
}

func(t *ttl_codes) iterable() iterator {
	i := iter{
		index:	0,
		length:	t.getLength(),
		comp:	dashboard_component(t),
	}
	return iterator(&i)
}

func(t *ttl_codes) getLength() int {
	return len((*t).parsed_snippets)
}

func(t *ttl_codes) getComponentLabel(cfg config.Config) string {
	return cfg.GetTTLCodesLabel()
}

func(t *ttl_codes) loadData(cfg config.Config, acc *account.UserAccount) {
	base_dir := cfg.GetTTLCodesDir()
	user_dir := filepath.FromSlash(cfg.GetTTLCodesDir() + acc.Name)

	snippetsFromDir(t.snippets, base_dir)

	if _, err := os.Stat(user_dir); os.IsNotExist(err) {
		return
	}
	snippetsFromDir(t.user_snippets, user_dir)
}

func(t *ttl_codes) parseData() {
	// no same title(file name) snippet allowed

	t.parsed_snippets = append(t.snippets, t.user_snippets...)
}

func(t *ttl_codes) intoRow(index int) [][]string {
	var rows [][]string
	for k, v := range t.parsed_snippets[index].snipMap {
		rows[0] = []string{"", "", k}		// 1st and 2nd columns are empty
		rows[1] = []string{"", "", string(v)}
		rows[2] = []string{"", "", ""}		// empty row between snippets
	} 
	return rows
}

func snippetsFromDir(snips []snippet, dir string) {
	file_names := utils.DirWalk(dir, onlyTextFile)
	for _, name := range file_names {
		bs, err := ioutil.ReadFile(filepath.Join(dir, name))
		if err != nil {
			log.Fatal(err)
		}
		snip := snippet{
			snipMap: map[string][]byte{name: bs},
		}
		snips = append(snips, snip)
	}
}

// used for utils.DirWalk
func onlyTextFile(dir string, file os.FileInfo) string {
	r := regexp.MustCompile(`txt$`)
	if file.IsDir() || !r.MatchString(file.Name()) {
		return ""
	}
	return file.Name()
}
