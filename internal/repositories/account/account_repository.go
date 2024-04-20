package account

import (
	"github.com/dhiemaz/AccountService/internal/model"
)

// Repository Account
type Repository interface {
	InsertAccount(data model.Account) (model.Account, error)
	UpdateAccount(data model.Account) (model.Account, error)
	FindOneAccountByUsername(username string) (model.Account, error)
	GetAllAccounts() ([]model.Account, error)
}
