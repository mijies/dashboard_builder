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
}

type snippet struct {
	snipMap		map[string][]byte
}

func(t *ttl_codes) getLength() int {
	return len((*t).snippets) + len((*t).user_snippets)
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

func(c *ttl_codes) ComponentIterator() []string {
	// book.SetCellValue(sheet_name, address, label)
	return []string{""}
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