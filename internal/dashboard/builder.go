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
	b.appendBooks()
	for _, book := range b.books {
		book.load()
	}
}

func(b *builder) appendBooks() {
	d := dashboard{
		path:	b.target_path,
	}
	b.books = append(b.books, book(&d))

	d = dashboard{
		path:	getMasterPath(),
	}
	b.books = append(b.books, book(&d))

	if _, err := os.Stat(getUserPath(b.acc.Name)); os.IsNotExist(err) {
		return
	}
	d = dashboard{
		path:	getUserPath(b.acc.Name),
	}
	b.books = append(b.books, book(&d))
}
