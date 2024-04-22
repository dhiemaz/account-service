package usecase

import (
	"github.com/dhiemaz/AccountService/infrastructures/logger"
	"github.com/dhiemaz/AccountService/internal/domain/office/model"
	"github.com/dhiemaz/AccountService/internal/domain/office/repository"
)

type OfficeUsecase struct {
	officeRepository repository.Repository
	out              Output
}

// NewOfficeUsecase object creation
func NewOfficeUsecase(officeRepository repository.Repository, out Output) OfficeUsecase {
	return OfficeUsecase{
		officeRepository: officeRepository,
		out:              out,
	}
}

// GetAllData access use case
func (i OfficeUsecase) GetAllData() ([]model.Office, error) {
	result, err := i.officeRepository.GetAllOfficeData()
	if err != nil {
		logger.WithFields(logger.Fields{"component": "usecase", "action": "GetAllData"}).
			Errorf("failed get all office data, error : %v", err)
	}

	return i.out.GetAllData(result, err)
}
