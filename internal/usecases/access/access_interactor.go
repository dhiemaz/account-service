package access

import (
	"github.com/dhiemaz/AccountService/infrastructures/logger"
	"github.com/dhiemaz/AccountService/internal/model"
	"github.com/dhiemaz/AccountService/internal/repositories/access"
)

type AccessInteractor struct {
	accessRepository access.Repository
	out              Output
}

// NewAccessInteractor object creation
func NewAccessInteractor(accessRepository access.Repository, out Output) *AccessInteractor {
	return &AccessInteractor{
		accessRepository: accessRepository,
		out:              out,
	}
}

// GetAllData access use case
func (i *AccessInteractor) GetAllData() ([]model.Access, error) {
	result, err := i.accessRepository.GetAllAccessData()
	if err != nil {
		logger.WithFields(logger.Fields{"component": "usecase", "action": "GetAllData"}).
			Errorf("failed get all access data, error : %v", err)
	}

	return i.out.GetAllData(result, err)
}
