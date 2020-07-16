package main

import (
	"flag"
	"github.com/mijies/dashboard_builder/internal/account"
	"github.com/mijies/dashboard_builder/internal/dashboard"
)

func main() {

	// Parse comand-line flags
	book_path := flag.String("f", "", "Path to the source bashboard excel file")
	user := flag.String("u", "foo", "Your name")
	login_user := flag.String("l", "foo15", "Your login name")
	flag.Parse()

	// user := "foo"
	// login_user := "foo15"
	// book_path := "samples/dashboard.xlsm"
	account := account.NewUserAccount(*user, *login_user)
	dashboard.Build(*book_path, account)
}
