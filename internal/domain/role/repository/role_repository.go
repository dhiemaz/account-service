package repository

import (
	"github.com/dhiemaz/AccountService/internal/domain/role/model"
)

// Repository Role
type Repository interface {
	GetAllRoleData() ([]model.Role, error)
}
