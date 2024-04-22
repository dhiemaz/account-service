package usecase

import (
	"github.com/dhiemaz/AccountService/infrastructures/logger"
	"github.com/dhiemaz/AccountService/internal/domain/role/model"
	"github.com/dhiemaz/AccountService/internal/domain/role/repository"
)

type RoleUsecase struct {
	roleRepository repository.Repository
	out            Output
}

// NewRoleUsecase object creation
func NewRoleUsecase(roleRepository repository.Repository, out Output) RoleUsecase {
	return RoleUsecase{
		roleRepository: roleRepository,
		out:            out,
	}
}

// GetAllData role use case
func (i *RoleUsecase) GetAllData() ([]model.Role, error) {
	result, err := i.roleRepository.GetAllRoleData()
	if err != nil {
		logger.WithFields(logger.Fields{"component": "usecase", "action": "GetAllData"}).
			Errorf("failed get all access data, error : %v", err)
	}

	return i.out.GetAllData(result, err)
}
