package grpc

import (
	"github.com/dhiemaz/AccountService/internal/domain/office/model"
)

type OfficePresenter struct{}

// GetAllData : office presenter function
func (p *OfficePresenter) GetAllData(data []model.Office, err error) ([]model.Office, error) {
	return data, err
}
