package grpc

import (
	"github.com/dhiemaz/AccountService/internal/domain/account/model"
	paginate "github.com/gobeam/mongo-go-pagination"
)

type AccountPresenter struct{}

// CreateNewAccountOutput : account presenter function
func (p *AccountPresenter) CreateNewAccountOutput(data model.Account, err error) (model.Account, error) {
	return data, err
}

// UpdateAccountOutput : account presenter function
func (p *AccountPresenter) UpdateAccountOutput(data model.Account, err error) (model.Account, error) {
	return data, err
}

// GetAllData : account presenter function
func (p *AccountPresenter) GetAllData(data []*model.Account, paging paginate.PaginationData, err error) ([]*model.Account, paginate.PaginationData, error) {
	return data, paging, err
}

// FindAccountByUsername : account presenter function
func (p *AccountPresenter) FindAccountByUsername(data model.Account, err error) (model.Account, error) {
	return data, err
}
