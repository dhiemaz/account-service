package usecase

import (
	"github.com/dhiemaz/AccountService/internal/domain/account/model"
	paginate "github.com/gobeam/mongo-go-pagination"
)

// Input : account input function
type Input interface {
	CreateNewAccount(data model.Account) (model.Account, error)
	UpdateAccount(data model.Account) (model.Account, error)
	GetAllData(searchFilter string, page, limit int) ([]model.Account, paginate.PaginationData, error)
	FindAccountByUsername(username string) (model.Account, error)
}

// Output : account output function
type Output interface {
	CreateNewAccountOutput(data model.Account, err error) (model.Account, error)
	UpdateAccountOutput(data model.Account, err error) (model.Account, error)
	GetAllData(data []model.Account, paging paginate.PaginationData, err error) ([]model.Account, paginate.PaginationData, error)
	FindAccountByUsername(data model.Account, err error) (model.Account, error)
}
