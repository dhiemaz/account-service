package usecase

import (
	"github.com/dhiemaz/AccountService/internal/domain/office/model"
)

// Input : office input function
type Input interface {
	GetAllData() ([]model.Office, error)
}

// Output : office output function
type Output interface {
	GetAllData(data []model.Office, err error) ([]model.Office, error)
}
