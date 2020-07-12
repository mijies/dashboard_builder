package main

import (
	// "flag"
	// "fmt"
	"github.com/mijies/dashboard_generator/internal/account"
	"github.com/mijies/dashboard_generator/internal/dashboard"
)

func main() {

	// Parse comand-line flags
	// book_path := flag.String("s", "", "Path to the source bashboard excel file")
	// user := flag.String("u", "foo", "Your name")
	// login_user := flag.String("l", "foo15", "Your login name")
	// flag.Parse()

	user := "foo"
	login_user := "foo15"
	book_path := "samples/dashboard.xlsm"
	account := account.NewUserAccount(user, login_user)
	dashboard.Build(book_path, account)
}
