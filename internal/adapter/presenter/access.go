package presenter

import "github.com/dhiemaz/AccountService/internal/model"

type AccessPresenter struct{}

// GetAllData : access presenter function
func (p *AccessPresenter) GetAllData(data []model.Access, err error) ([]model.Access, error) {
	return data, err
}
