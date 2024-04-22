package repository

import (
	"github.com/dhiemaz/AccountService/internal/domain/office/model"
)

// Repository Office
type Repository interface {
	GetAllOfficeData() ([]model.Office, error)
}
