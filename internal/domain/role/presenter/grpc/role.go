package grpc

import (
	"github.com/dhiemaz/AccountService/internal/domain/role/model"
)

type RolePresenter struct{}

// GetAllData : role presenter function
func (p *RolePresenter) GetAllData(data []model.Role, err error) ([]model.Role, error) {
	return data, err
}
