package usecase

import (
	"github.com/dhiemaz/AccountService/internal/domain/role/model"
)

// Input : role input function
type Input interface {
	GetAllData() ([]model.Role, error)
}

// Output : role output function
type Output interface {
	GetAllData(data []model.Role, err error) ([]model.Role, error)
}
