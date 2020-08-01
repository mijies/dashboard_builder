package dashboard

import (
	// "fmt"
	"os"
	"github.com/mijies/dashboard_builder/internal/account"
)

type Builder interface {
	loadBooks()
	parseBooks()
	buildBook()
}

type builder struct {
	// acc			*account.UserAccount
	books		[]dbook
}

func Build(target_path string, acc *account.UserAccount) {
	paths := []string{
		target_path,
		getMasterPath(),
	}
	if _, err := os.Stat(getUserPath(acc.Name)); !os.IsNotExist(err) {
		paths = append(paths, getUserPath(acc.Name))
	}
	b := builder{}
	b.loadBooks(paths)
	b.parseBooks()
	b.buildBook()
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

func(b *builder) loadBooks(paths []string) {
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
	mtx  := make(chan bool)
	sch  := make(chan *[]snippet)
	cch  := make(chan *[]command)
	done := make(chan bool)

	for _, book := range b.books {
		go func(book dbook) {
			book.parse(mtx, sch, cch)
			done <- true
		}(book)
	}

	for _ = range b.books {
		<- done
	}
}

func(b *builder) buildBook() {
	done := make(chan bool)
	for _, book := range b.books {
		go func() {
			book.build()
			done <- true
		}()
	}

	for _ = range b.books {
		<- done
	}
}

