package usecase

import (
	"github.com/dhiemaz/AccountService/internal/domain/access/model"
)

// Input : access input function
type Input interface {
	GetAllData() ([]model.Access, error)
}

// Output : access output function
type Output interface {
	GetAllData(data []model.Access, err error) ([]model.Access, error)
}
