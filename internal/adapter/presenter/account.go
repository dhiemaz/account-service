package presenter

import "github.com/dhiemaz/AccountService/internal/model"

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
func (p *AccountPresenter) GetAllData(data []model.Account, err error) ([]model.Account, error) {
	return data, err
}

// FindAccountByUsername : account presenter function
func (p *AccountPresenter) FindAccountByUsername(data model.Account, err error) (model.Account, error) {
	return data, err
}
