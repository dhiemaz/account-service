package role

import (
	"github.com/dhiemaz/AccountService/infrastructures/logger"
	"github.com/dhiemaz/AccountService/internal/model"
	"github.com/dhiemaz/AccountService/internal/repositories/role"
)

type RoleInteractor struct {
	roleRepository role.Repository
	out            Output
}

// NewRoleInteractor object creation
func NewRoleInteractor(roleRepository role.Repository, out Output) *RoleInteractor {
	return &RoleInteractor{
		roleRepository: roleRepository,
		out:            out,
	}
}

// GetAllData role use case
func (i *RoleInteractor) GetAllData() ([]model.Role, error) {
	result, err := i.roleRepository.GetAllRoleData()
	if err != nil {
		logger.WithFields(logger.Fields{"component": "usecase", "action": "GetAllData"}).
			Errorf("failed get all access data, error : %v", err)
	}

	return i.out.GetAllData(result, err)
}
