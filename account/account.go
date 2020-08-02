package account

type UserAccount struct {
	Name		string
	Login_name	string
}

func NewUserAccount(name string, login_name string) *UserAccount {
	return &UserAccount{
		Name: 		name,
		Login_name: login_name,
	}
}