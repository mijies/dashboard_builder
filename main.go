package main

import (
	"flag"
	"log"
	"os"
	"path/filepath"
	"github.com/mijies/dashboard_builder/account"
	"github.com/mijies/dashboard_builder/dashboard"
)

func main() {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
    if err != nil {
		log.Fatal(err)
	}
	os.Chdir(dir)

	book_path  := account.DEFAULT_DASHBOARD_PATH
	user       := flag.String("u", "", "Your name")
	login_user := flag.String("l", "", "Your login name")

	flag.Parse()
	args := flag.Args()
	if len(args) != 0 {
		book_path = args[0]
	}

	account := account.NewUserAccount(*user, *login_user)
	dashboard.Build(book_path, account)
}
