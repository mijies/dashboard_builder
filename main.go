package main

import (
	"flag"
	"github.com/mijies/dashboard_builder/account"
	"github.com/mijies/dashboard_builder/dashboard"
)

func main() {

	// Parse comand-line
	book_path := *flag.String("f", "samples/dashboard.xlsm", "Path to the target bashboard.xlsm")
	user := *flag.String("u", "foo", "Your name")
	login_user := *flag.String("l", "foo15", "Your login name")
	flag.Parse()

	account := account.NewUserAccount(user, login_user)
	dashboard.Build(book_path, account)
}
