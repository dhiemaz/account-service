package grpc

import (
	"github.com/dhiemaz/AccountService/internal/domain/access/model"
)

type AccessPresenter struct{}

// GetAllData : access presenter function
func (p *AccessPresenter) GetAllData(data []model.Access, err error) ([]model.Access, error) {
	return data, err
}
