package dashboard

import (
	// "fmt"
	"os"
	"github.com/mijies/dashboard_builder/internal/account"
)

type Builder interface {
	loadBooks()
	parseBooks()
}

type builder struct {
	target_path	string
	acc			*account.UserAccount
	books		[]dbook
}

func Build(target_path string, acc *account.UserAccount) {
	b := builder{
		target_path:	target_path,
		acc:			acc,
	}
	b.loadBooks()
	b.parseBooks()
}

func(b *builder) newBook(path string) dbook {
	var d dbook
	switch(path) {
		case b.target_path:
			d = dbook(&targetBook{dashboard{path: path}})
		case getMasterPath():
			d = dbook(&masterBook{dashboard{path: path}})
		case getUserPath(b.acc.Name):
			d = dbook(&userBook{dashboard{path: path}})
	}
	return d
}

func(b *builder) loadBooks() {
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
			book := b.newBook(path)
			book.load()
			b.books = append(b.books, book)
			done <- true
		}(path)
	}

	for _ = range paths {
		<- done
	}
}

func(b *builder) parseBooks() {
	sch  := make(chan bool)
	done := make(chan bool)
	var cmds	*commands
	var snip	*snippets

	for _, book := range b.books {
		if book.get_path() == b.target_path {
			cmds = book.ref_commands()
			snip = book.ref_snippets()
			continue
		}
		go func(book dbook) {
			book.parse(cmds, snip, sch)
			done <- true
		}(book)
	}

	for _ = range b.books[1:] { // -1 for targetBook
		<- done
	}
}

