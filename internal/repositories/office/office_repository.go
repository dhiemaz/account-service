package office

import "github.com/dhiemaz/AccountService/internal/model"

// Repository Office
type Repository interface {
	GetAllOfficeData() ([]model.Office, error)
}
