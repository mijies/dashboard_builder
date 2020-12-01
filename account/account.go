package account

import (
	// "log"
	"os"
	"github.com/mijies/dashboard_builder/utils"
)

const DEFAULT_DASHBOARD_PATH = "samples/dashboard.xlsm"
const DEFAULT_ACCOUNT_TXT    = "samples/dashboard.txt"

type UserAccount struct {
	Name		string
	Login_name	string
}

func NewUserAccount(name string, login_name string) *UserAccount {
	if name != "" && login_name != "" {
		return &UserAccount{
			Name: 		name,
			Login_name: login_name,
		}
	}
	return get_account_info()
}

func get_account_info() *UserAccount {
	// first try a file
	if _, err := os.Stat(DEFAULT_ACCOUNT_TXT); !os.IsNotExist(err) {
		lines := utils.ReadLineSlice(DEFAULT_ACCOUNT_TXT)
		if len(lines) == 2 {
			return &UserAccount{
				Name: 		lines[0],
				Login_name: lines[1],
			}
		}
	}
	// try gui input
	return get_account_from_dialogbox()
}
