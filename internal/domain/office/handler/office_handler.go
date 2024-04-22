package handler

import (
	"context"
	"github.com/dhiemaz/AccountService/infrastructures/server/grpc/proto/office_svc"
	"github.com/dhiemaz/AccountService/internal/domain/office/usecase"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type OfficeHandler struct {
	officeUseCase usecase.OfficeUsecase
}

func NewOfficeHandler(officeUseCase usecase.OfficeUsecase) OfficeHandler {
	return OfficeHandler{
		officeUseCase: officeUseCase,
	}
}

func (h OfficeHandler) SearchOffice(ctx context.Context, input *office_svc.SearchOfficeRequest) (*office_svc.SearchOfficeResponse, error) {
	result, err := h.officeUseCase.GetAllData()
	if err != nil {
		if e, ok := status.FromError(err); ok {
			return nil, status.Errorf(e.Code(), e.Message())
		}

		return nil, status.Errorf(codes.Unknown, err.Error())
	}

	var datas []*office_svc.Data
	for _, val := range result {
		datas = append(datas, &office_svc.Data{
			Id:   int32(val.BranchId),
			Name: val.BranchName,
		})
	}

	return &office_svc.SearchOfficeResponse{
		Meta: nil,
		Data: datas,
	}, nil
}
