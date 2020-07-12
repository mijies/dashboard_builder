package dashboard

import (
	"fmt"
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
	finalized		[]snippet
	rows			[][]string // made by intoRows()
}

type snippet struct {
	snipMap		map[string][]byte
}

func(t *ttl_codes) iterable() iterator {
	i := iter{
		index:	0,
		length:	t.getLength(),
		items:	&t.rows,
	}
	return iterator(&i)
}

func(t *ttl_codes) getLength() int {
	if len(t.rows) != 0 {
		return len(t.rows)
	}
	if len(t.finalized) != 0 {
		return len(t.finalized)
	}
	return -1 // length is unknown until finalized
}

func(t *ttl_codes) getComponentLabel(cfg config.Config) string {
	return cfg.GetTTLCodesLabel()
}

func(t *ttl_codes) loadData(cfg config.Config, acc *account.UserAccount) {
	base_dir := cfg.GetTTLCodesDir()
	user_dir := filepath.FromSlash(cfg.GetTTLCodesDir() + acc.Name)

	snippetsFromDir(&t.snippets, base_dir)

	if _, err := os.Stat(user_dir); os.IsNotExist(err) {
		return
	}
	snippetsFromDir(&t.user_snippets, user_dir)
}

func(t *ttl_codes) finalize() {
	// title(file name) duplication not allowed
	for _, s := range t.snippets {
		for _, u := range t.user_snippets {
			for sk, _ := range s.snipMap {
				for uk, _ := range u.snipMap {
					if sk == uk {
						log.Fatal("code name duplication with your custom code: " + sk)
					}
				}
			}
		}
	}

	t.finalized = append(t.snippets, t.user_snippets...)
	t.snippets 		= nil
	t.user_snippets = nil
	t.intoRows()
}

func(t *ttl_codes) intoRows() {
	for _, s := range t.finalized {
		for k, v := range s.snipMap {
			t.rows = append(t.rows, []string{"", "", "[" + k + "]"}) // 1st and 2nd columns are empty
			t.rows = append(t.rows, []string{"", "", string(v)})
			t.rows = append(t.rows, []string{"", "", ""})
		}
	}
	t.finalized	= nil
}

func snippetsFromDir(snips *[]snippet, dir string) {
	file_names := utils.DirWalk(dir, onlyTextFile)
	for _, name := range file_names {
		bs, err := ioutil.ReadFile(filepath.Join(dir, name))
		if err != nil {
			log.Fatal(err)
		}
		snip := snippet{
			snipMap: map[string][]byte{name[:len(name)-4]: bs}, // remove .txt from name
		}
		*snips = append(*snips, snip)
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
