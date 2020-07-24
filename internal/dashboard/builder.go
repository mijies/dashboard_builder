package dashboard

import (
	"os"
	"github.com/mijies/dashboard_builder/internal/account"
)

type Builder interface {
	load()
	appendBooks()
}

type builder struct {
	target_path	string
	acc			*account.UserAccount
	books		[]book
}

func Build(target_path string, acc *account.UserAccount) {
	b := builder{
		target_path:	target_path,
		acc:			acc,
	}
	b.load()
}

func(b *builder) load() {

	paths := []string{
		b.target_path,
		getMasterPath(),
	}
	if _, err := os.Stat(getUserPath(b.acc.Name)); !os.IsNotExist(err) {
		paths = append(paths, getUserPath(b.acc.Name))
	}

	done := make(chan bool)
	for _, path := range paths {
		go func(path string) {
			bk := b.newBook(path)
			bk.load()
			b.books = append(b.books, bk)
			done <- true
		}(path)
	}

	for _ = range paths {
		<- done
	}
}

func(b *builder) newBook(path string) book {
	var d book
	switch(path) {
		case b.target_path:
			d = book(&targetBook{dashboard{path: path}})
		case getMasterPath():
			d = book(&masterBook{dashboard{path: path}})
		case getUserPath(b.acc.Name):
			d = book(&userBook{dashboard{path: path}})
	}
	return d
}
