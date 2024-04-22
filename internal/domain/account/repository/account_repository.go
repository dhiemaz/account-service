package repository

import (
	"github.com/dhiemaz/AccountService/internal/domain/account/model"
	paginate "github.com/gobeam/mongo-go-pagination"
)

// Repository Account
type Repository interface {
	InsertAccount(data model.Account) (model.Account, error)
	UpdateAccount(data model.Account) (model.Account, error)
	FindOneAccountByUsername(username string) (model.Account, error)
	GetAllAccounts(searchFilter string, page, limit int) ([]model.Account, paginate.PaginationData, error)
}
