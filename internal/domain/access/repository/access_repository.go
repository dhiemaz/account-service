package repository

import (
	"github.com/dhiemaz/AccountService/internal/domain/access/model"
)

// Repository Access
type Repository interface {
	GetAllAccessData() ([]model.Access, error)
}
