package models

// Account of the user
type Account struct {
	Status    AccountStatus
	AccountId string
	Username  string
	Password  string
	Email     string
}

func (a *Account) resetPassword() bool {
	if a.Password == "abc" {
		return true
	}
	return false
}

type AccountStatus string
