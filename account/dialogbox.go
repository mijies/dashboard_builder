package account

import (
	// "fmt"
	"log"
	"github.com/lxn/walk"
	. "github.com/lxn/walk/declarative"
)

type mainWindow struct {
	*walk.MainWindow
	// acc   *UserAccount
}

type dialogWindow struct{
	dlg      	*walk.Dialog
	user_name   *walk.LineEdit
	login_name	*walk.LineEdit
	acceptPB 	*walk.PushButton
	cancelPB 	*walk.PushButton
}

func(w *dialogWindow) acceptCliked(acc *UserAccount) { 
	acc.Name = w.user_name.Text()
	acc.Login_name = w.login_name.Text()
	w.dlg.Accept()
}

func get_account_from_dialogbox() *UserAccount {

	acc := &UserAccount{}
	mainW := &mainWindow{}

	mw := MainWindow{
		AssignTo: &mainW.MainWindow,
		Visible: true,
		Title: "DialogBoxテスト",
		MinSize: Size {150, 200},
		Size   : Size {300, 300},
		Layout: VBox {},
	}
	// func(MainWindow) {
	// 	return
	// }(mw)
	if _, err := mw.Run(); err != nil {
		log.Fatal(err)
		// os.Exit(1)
	}

	dlgW := new(dialogWindow)
	dlg := Dialog{
		AssignTo:      &dlgW.dlg,
		Title:         "User Information",
		DefaultButton: &dlgW.acceptPB,
		CancelButton:  &dlgW.cancelPB,
		MinSize: Size{300, 100},
		Layout:  VBox{},
		Children: []Widget{
			Composite{
				Layout: HBox{},
				Children: []Widget{
					Label{
						Text: "User Name:",
					},
					LineEdit{
						AssignTo: &dlgW.user_name,
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
						AssignTo: &dlgW.login_name,
					},
				},
			},
			Composite{
				Layout: HBox{},
				Children: []Widget{
					PushButton{
						AssignTo: &dlgW.acceptPB,
						Text:     "OK",
						OnClicked: func(){
							dlgW.acceptCliked(acc)
						},
					},
					PushButton{
						AssignTo:  &dlgW.cancelPB,
						Text:      "Cancel",
						OnClicked: func() { dlgW.dlg.Cancel() },
					},
				},
			},
		},
	}

	cmd, err := dlg.Run(walk.Form(mainW))
	if err != nil {
		log.Fatal(err)
	} else if cmd == walk.DlgCmdCancel { 
		return acc
	} else if cmd == walk.DlgCmdNone {
		return acc
	}
	return acc
}