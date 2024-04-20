package account

import (
	"github.com/dhiemaz/AccountService/internal/model"
)

// Input : account input function
type Input interface {
	CreateNewAccount(data model.Account) (model.Account, error)
	UpdateAccount(data model.Account) (model.Account, error)
	GetAllData() ([]model.Account, error)
	FindAccountByUsername(username string) (model.Account, error)
}

// Output : account output function
type Output interface {
	CreateNewAccountOutput(data model.Account, err error) (model.Account, error)
	UpdateAccountOutput(data model.Account, err error) (model.Account, error)
	GetAllData(data []model.Account, err error) ([]model.Account, error)
	FindAccountByUsername(data model.Account, err error) (model.Account, error)
}
