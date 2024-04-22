package usecase

import (
	"github.com/dhiemaz/AccountService/infrastructures/logger"
	"github.com/dhiemaz/AccountService/internal/domain/access/model"
	"github.com/dhiemaz/AccountService/internal/domain/access/repository"
)

type AccessUsecase struct {
	accessRepository repository.Repository
	out              Output
}

// NewAccessUsecase object creation
func NewAccessUsecase(accessRepository repository.Repository, out Output) AccessUsecase {
	return AccessUsecase{
		accessRepository: accessRepository,
		out:              out,
	}
}

// GetAllData access use case
func (u AccessUsecase) GetAllData() ([]model.Access, error) {
	result, err := u.accessRepository.GetAllAccessData()
	if err != nil {
		logger.WithFields(logger.Fields{"component": "usecase", "action": "GetAllData"}).
			Errorf("failed get all access data, error : %v", err)
	}

	return u.out.GetAllData(result, err)
}
