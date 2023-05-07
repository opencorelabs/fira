package api

import (
	"context"

	v1 "github.com/opencorelabs/fira/gen/protos/go/protos/fira/v1"
)

type FiraApiService struct {
	v1.UnimplementedFiraServiceServer
}

func (FiraApiService) GetApiInfo(context.Context, *v1.GetApiInfoRequest) (*v1.GetApiInfoResponse, error) {
	return &v1.GetApiInfoResponse{
		Version: &v1.GetApiInfoResponse_Version{
			Major: 1,
			Minor: 0,
			Patch: 0,
		},
	}, nil
}
