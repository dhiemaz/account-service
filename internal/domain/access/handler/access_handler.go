package handler

import (
	"context"
	"github.com/dhiemaz/AccountService/infrastructures/server/grpc/proto/access_svc"
	"github.com/dhiemaz/AccountService/internal/domain/access/usecase"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type AccessHandler struct {
	accessInput usecase.Input
}

func NewAccessHandler(accessInput usecase.Input) AccessHandler {
	return AccessHandler{
		accessInput: accessInput,
	}
}

func (h AccessHandler) SearchAccessData(ctx context.Context, request *access_svc.SearchAccessRequest) (*access_svc.SearchAccessResponse, error) {
	result, err := h.accessInput.GetAllData()
	if err != nil {
		if e, ok := status.FromError(err); ok {
			return nil, status.Errorf(e.Code(), e.Message())
		}

		return nil, status.Errorf(codes.Unknown, err.Error())
	}

	var datas []*access_svc.Data
	for _, val := range result {
		datas = append(datas, &access_svc.Data{
			Id:   int32(val.AccessId),
			Name: val.AccessName,
		})
	}

	return &access_svc.SearchAccessResponse{
		Meta: nil,
		Data: datas,
	}, nil
}
