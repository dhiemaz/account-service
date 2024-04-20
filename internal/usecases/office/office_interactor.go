package office

import (
	"github.com/dhiemaz/AccountService/infrastructures/logger"
	"github.com/dhiemaz/AccountService/internal/model"
	"github.com/dhiemaz/AccountService/internal/repositories/office"
)

type OfficeInteractor struct {
	officeRepository office.Repository
	out              Output
}

// NewOfficeInteractor object creation
func NewOfficeInteractor(officeRepository office.Repository, out Output) *OfficeInteractor {
	return &OfficeInteractor{
		officeRepository: officeRepository,
		out:              out,
	}
}

// GetAllData access use case
func (i *OfficeInteractor) GetAllData() ([]model.Office, error) {
	result, err := i.officeRepository.GetAllOfficeData()
	if err != nil {
		logger.WithFields(logger.Fields{"component": "usecase", "action": "GetAllData"}).
			Errorf("failed get all office data, error : %v", err)
	}

	return i.out.GetAllData(result, err)
}
