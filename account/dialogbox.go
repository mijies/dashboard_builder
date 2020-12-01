package account

import (
	// "fmt"
	"log"
	"github.com/lxn/walk"
	. "github.com/lxn/walk/declarative"
)

type mainWindow struct {
	*walk.MainWindow
	user_name   *walk.LineEdit
	login_name	*walk.LineEdit
	acceptPB 	*walk.PushButton
}

func(w *mainWindow) acceptClick(acc *UserAccount) { 
	acc.Name = w.user_name.Text()
	acc.Login_name = w.login_name.Text()
	w.Close()
}

func get_account_from_dialogbox() *UserAccount {
	acc := &UserAccount{}
	mainW := &mainWindow{}

	if _, err := (MainWindow{
		AssignTo: &mainW.MainWindow,
		Visible: true,
		Title: "Account Info",
		MinSize: Size {300, 100},
		// Size   : Size {300, 300},
		Layout: VBox {},
		Children: []Widget{
			Composite{
				Layout: HBox{},
				Children: []Widget{
					Label{
						Text: "User Name :",
					},
					LineEdit{
						AssignTo: &mainW.user_name,
					},
				},
			},
			Composite{
				Layout: HBox{},
				Children: []Widget{
					Label{
						Text: "Login Name:",
					},
					LineEdit{
						AssignTo: &mainW.login_name,
					},
				},
			},
			Composite{
				Layout: HBox{},
				Children: []Widget{
					PushButton{
						AssignTo: &mainW.acceptPB,
						Text:     "OK",
						OnClicked: func(){
							mainW.acceptClick(acc)
						},
					},
				},
			},
		},
	}).Run(); err != nil {
		log.Fatal(err)
	}
	return acc
}