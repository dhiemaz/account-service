package access

import (
	"github.com/dhiemaz/AccountService/internal/model"
)

// Repository Access
type Repository interface {
	GetAllAccessData() ([]model.Access, error)
}
