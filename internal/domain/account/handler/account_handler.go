package handler

import (
	"context"
	"github.com/dhiemaz/AccountService/infrastructures/server/grpc/proto/account_svc"
	accessModel "github.com/dhiemaz/AccountService/internal/domain/access/model"
	"github.com/dhiemaz/AccountService/internal/domain/account/model"
	"github.com/dhiemaz/AccountService/internal/domain/account/usecase"
	officeModel "github.com/dhiemaz/AccountService/internal/domain/office/model"
	roleModel "github.com/dhiemaz/AccountService/internal/domain/role/model"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"strings"
)

type AccountHandler struct {
	accountUseCase usecase.Input
}

func NewAccountHandler(accountUseCase usecase.Input) AccountHandler {
	return AccountHandler{
		accountUseCase: accountUseCase,
	}
}

func (h AccountHandler) AccountRegistration(ctx context.Context, in *account_svc.AccountRegistrationRequest) (*account_svc.AccountRegistrationResponse, error) {
	// map offices from grpc model to db model
	var offices []officeModel.Office
	for _, val := range in.Office {
		offices = append(offices, officeModel.Office{
			BranchId:   int(val.Id),
			BranchName: val.Name,
		})
	}

	// map access from grpc model to db model
	var access []accessModel.Access
	for _, val := range in.Access {
		access = append(access, accessModel.Access{
			AccessId:   int(val.Id),
			AccessName: val.Name,
		})
	}

	// map roles from grpc model to db model
	var roles []roleModel.Role
	for _, val := range in.Role {
		roles = append(roles, roleModel.Role{
			RoleId:   int(val.Id),
			RoleName: val.Name,
		})
	}

	registrationData := model.Account{
		Username: in.Username,
		Password: in.Password,
		Fullname: in.UserData.Fullname,
		Address:  in.UserData.Address,
		Zipcode:  in.UserData.Zipcode,
		Province: in.UserData.Province,
		Office:   offices,
		Role:     roles,
		Access:   access,
	}

	_, err := h.accountUseCase.CreateNewAccount(registrationData)
	if err != nil {
		if strings.Contains(err.Error(), "username is already taken, please choose another one") {
			response := account_svc.AccountRegistrationResponse{
				Status: &account_svc.Response{
					Code:    "01",
					Message: err.Error(),
				},
				Data: nil,
			}
			return &response, nil
		}

		if strings.Contains(err.Error(), "invalid branch_id or branch_name") {
			response := account_svc.AccountRegistrationResponse{
				Status: &account_svc.Response{
					Code:    "02",
					Message: err.Error(),
				},
				Data: nil,
			}
			return &response, nil
		}

		if strings.Contains(err.Error(), "invalid role_id or role_name") {
			response := account_svc.AccountRegistrationResponse{
				Status: &account_svc.Response{
					Code:    "02",
					Message: err.Error(),
				},
				Data: nil,
			}
			return &response, nil
		}

		if strings.Contains(err.Error(), "invalid access_id or access_name") {
			response := account_svc.AccountRegistrationResponse{
				Status: &account_svc.Response{
					Code:    "02",
					Message: err.Error(),
				},
				Data: nil,
			}
			return &response, nil
		}

		return nil, status.Errorf(codes.Unknown, err.Error())
	}

	response := account_svc.AccountRegistrationResponse{
		Status: &account_svc.Response{
			Code:    "00",
			Message: "success",
		},
		Data: &account_svc.RegistrationData{
			Username: in.Username,
			UserData: &account_svc.UserData{
				Fullname: in.UserData.Fullname,
				Address:  in.UserData.Address,
				Zipcode:  in.UserData.Zipcode,
				Province: in.UserData.Province,
			},
			Office: in.Office,
			Role:   in.Role,
			Access: in.Access,
		},
	}

	return &response, nil
}

func (h AccountHandler) AccountUpdate(ctx context.Context, in *account_svc.AccountUpdateRequest) (*account_svc.AccountUpdateResponse, error) {
	// map offices from grpc model to db model
	var offices []officeModel.Office
	for _, val := range in.Office {
		offices = append(offices, officeModel.Office{
			BranchId:   int(val.Id),
			BranchName: val.Name,
		})
	}

	// map access from grpc model to db model
	var access []accessModel.Access
	for _, val := range in.Access {
		access = append(access, accessModel.Access{
			AccessId:   int(val.Id),
			AccessName: val.Name,
		})
	}

	// map roles from grpc model to db model
	var roles []roleModel.Role
	for _, val := range in.Role {
		roles = append(roles, roleModel.Role{
			RoleId:   int(val.Id),
			RoleName: val.Name,
		})
	}

	updateData := model.Account{
		Username: in.Username,
		Fullname: in.Data.Fullname,
		Address:  in.Data.Address,
		Zipcode:  in.Data.Zipcode,
		Province: in.Data.Province,
		Office:   offices,
		Role:     roles,
		Access:   access,
	}

	_, err := h.accountUseCase.UpdateAccount(updateData)
	if err != nil {
		if strings.Contains(err.Error(), "username is not found") {
			response := account_svc.AccountUpdateResponse{
				Status: &account_svc.Response{
					Code:    "01",
					Message: err.Error(),
				},
			}
			return &response, nil
		}

		if strings.Contains(err.Error(), "invalid branch_id or branch_name") {
			response := account_svc.AccountUpdateResponse{
				Status: &account_svc.Response{
					Code:    "02",
					Message: err.Error(),
				},
			}
			return &response, nil
		}

		if strings.Contains(err.Error(), "invalid role_id or role_name") {
			response := account_svc.AccountUpdateResponse{
				Status: &account_svc.Response{
					Code:    "02",
					Message: err.Error(),
				},
			}
			return &response, nil
		}

		if strings.Contains(err.Error(), "invalid access_id or access_name") {
			response := account_svc.AccountUpdateResponse{
				Status: &account_svc.Response{
					Code:    "02",
					Message: err.Error(),
				},
			}
			return &response, nil
		}

		return nil, status.Errorf(codes.Unknown, err.Error())
	}

	response := account_svc.AccountUpdateResponse{
		Status: &account_svc.Response{
			Code:    "00",
			Message: "success",
		},
		Username: in.Username,
		UserData: in.Data,
		Role:     in.Role,
		Access:   in.Access,
		Office:   in.Office,
	}

	return &response, nil
}

func (h AccountHandler) AccountSearch(ctx context.Context, in *account_svc.AccountSearchRequest) (*account_svc.AccountSearchResponse, error) {
	results, pagination, err := h.accountUseCase.GetAllData(in.SearchFilter.Value, int(in.Pagination.Page), int(in.Pagination.Limit))
	if err != nil {
		if e, ok := status.FromError(err); ok {
			return nil, status.Errorf(e.Code(), e.Message())
		}

		return nil, status.Errorf(codes.Unknown, err.Error())
	}

	var accounts []*account_svc.RegistrationData
	for _, result := range results {
		account := account_svc.RegistrationData{
			Username: result.Username,
			UserData: &account_svc.UserData{
				Fullname: result.Fullname,
				Address:  result.Address,
				Zipcode:  result.Zipcode,
				Province: result.Province,
			},
		}

		var offices []*account_svc.Data
		// map offices data
		for _, val := range result.Office {
			offices = append(offices, &account_svc.Data{
				Id:   int32(val.BranchId),
				Name: val.BranchName,
			})
		}

		var roles []*account_svc.Data
		// map roles data
		for _, val := range result.Role {
			roles = append(roles, &account_svc.Data{
				Id:   int32(val.RoleId),
				Name: val.RoleName,
			})
		}

		var access []*account_svc.Data
		// map access data
		for _, val := range result.Access {
			access = append(access, &account_svc.Data{
				Id:   int32(val.AccessId),
				Name: val.AccessName,
			})
		}

		account.Office = offices
		account.Role = roles
		account.Access = access

		accounts = append(accounts, &account)
	}

	paginationData := account_svc.Pagination{
		Total:     int32(pagination.Total),
		Page:      int32(pagination.Page),
		PerPage:   int32(pagination.Total),
		Prev:      int32(pagination.Prev),
		Next:      int32(pagination.Next),
		TotalPage: int32(pagination.TotalPage),
	}

	status := account_svc.Response{
		Code:    "00",
		Message: "success",
	}

	var response account_svc.AccountSearchResponse
	response.Status = &status
	response.Pagination = &paginationData
	response.Data = accounts

	return &response, nil
}

func (h AccountHandler) AccountSearchByUsername(ctx context.Context, in *account_svc.AccountSearchByUsernameRequest) (*account_svc.AccountSearchByUsernameResponse, error) {
	result, err := h.accountUseCase.FindAccountByUsername(in.Username)
	if err != nil {
		if e, ok := status.FromError(err); ok {
			return nil, status.Errorf(e.Code(), e.Message())
		}

		return nil, status.Errorf(codes.Unknown, err.Error())
	}

	var offices []*account_svc.Data
	var roles []*account_svc.Data
	var access []*account_svc.Data

	// map offices data
	for _, val := range result.Office {
		offices = append(offices, &account_svc.Data{
			Id:   int32(val.BranchId),
			Name: val.BranchName,
		})
	}

	// map roles data
	for _, val := range result.Role {
		roles = append(roles, &account_svc.Data{
			Id:   int32(val.RoleId),
			Name: val.RoleName,
		})
	}

	// map roles data
	for _, val := range result.Access {
		access = append(access, &account_svc.Data{
			Id:   int32(val.AccessId),
			Name: val.AccessName,
		})
	}

	response := account_svc.AccountSearchByUsernameResponse{
		Status: &account_svc.Response{
			Code:    "00",
			Message: "success",
		},
		Data: &account_svc.RegistrationData{
			Username: result.Username,
			UserData: &account_svc.UserData{
				Fullname: result.Fullname,
				Address:  result.Address,
				Zipcode:  result.Zipcode,
				Province: result.Province,
			},
			Office: offices,
			Role:   roles,
			Access: access,
		},
	}

	return &response, nil
}
