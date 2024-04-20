package role

import (
	"github.com/dhiemaz/AccountService/internal/model"
)

// Repository Role
type Repository interface {
	GetAllRoleData() ([]model.Role, error)
}
