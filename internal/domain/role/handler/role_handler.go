package handler

import (
	"context"
	"github.com/dhiemaz/AccountService/infrastructures/server/grpc/proto/role_svc"
	"github.com/dhiemaz/AccountService/internal/domain/role/usecase"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type RoleHandler struct {
	roleUseCase usecase.RoleUsecase
}

func NewRoleHandler(roleUseCase usecase.RoleUsecase) RoleHandler {
	return RoleHandler{
		roleUseCase: roleUseCase,
	}
}

func (h RoleHandler) SearchRole(ctx context.Context, request *role_svc.SearchRoleRequest) (*role_svc.SearchRoleResponse, error) {
	result, err := h.roleUseCase.GetAllData()
	if err != nil {
		if e, ok := status.FromError(err); ok {
			return nil, status.Errorf(e.Code(), e.Message())
		}

		return nil, status.Errorf(codes.Unknown, err.Error())
	}

	var datas []*role_svc.Data
	for _, val := range result {
		datas = append(datas, &role_svc.Data{
			Id:   int32(val.RoleId),
			Name: val.RoleName,
		})
	}

	return &role_svc.SearchRoleResponse{
		Meta: nil,
		Data: datas,
	}, nil
}
