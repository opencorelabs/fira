package api

import (
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	v1 "github.com/opencorelabs/fira/gen/protos/go/protos/fira/v1"
)

type FiraApiService struct {
	v1.UnimplementedFiraServiceServer
	acctSvc   *AccountService
	appSvc    *AppService
	finAggSvc *FinancialAggregatorService
}

func New(
	acctService *AccountService,
	appService *AppService,
	finAggService *FinancialAggregatorService,
) v1.FiraServiceServer {
	fs := &FiraApiService{
		acctSvc:   acctService,
		appSvc:    appService,
		finAggSvc: finAggService,
	}
	return fs
}

func (*FiraApiService) GetApiInfo(context.Context, *v1.GetApiInfoRequest) (*v1.GetApiInfoResponse, error) {
	return &v1.GetApiInfoResponse{
		Version: &v1.GetApiInfoResponse_Version{
			Major: 1,
			Minor: 0,
			Patch: 0,
		},
	}, nil
}

///
/// Accounts API
///

func (s *FiraApiService) CreateAccount(ctx context.Context, request *v1.CreateAccountRequest) (*v1.CreateAccountResponse, error) {
	return s.acctSvc.CreateAccount(ctx, request)
}

func (s *FiraApiService) VerifyAccount(ctx context.Context, request *v1.VerifyAccountRequest) (*v1.VerifyAccountResponse, error) {
	return s.acctSvc.VerifyAccount(ctx, request)
}

func (s *FiraApiService) LoginAccount(ctx context.Context, request *v1.LoginAccountRequest) (*v1.LoginAccountResponse, error) {
	return s.acctSvc.LoginAccount(ctx, request)
}

func (s *FiraApiService) BeginPasswordReset(ctx context.Context, request *v1.BeginPasswordResetRequest) (*v1.BeginPasswordResetResponse, error) {
	return s.acctSvc.BeginPasswordReset(ctx, request)
}

func (s *FiraApiService) CompletePasswordReset(ctx context.Context, request *v1.CompletePasswordResetRequest) (*v1.CompletePasswordResetResponse, error) {
	return s.acctSvc.CompletePasswordReset(ctx, request)
}

func (s *FiraApiService) GetAccount(ctx context.Context, request *v1.GetAccountRequest) (*v1.GetAccountResponse, error) {
	return s.acctSvc.GetAccount(ctx, request)
}

///
/// Apps API
///

func (s *FiraApiService) CreateApp(ctx context.Context, request *v1.CreateAppRequest) (*v1.CreateAppResponse, error) {
	return s.appSvc.CreateApp(ctx, request)
}

func (s *FiraApiService) ListApps(ctx context.Context, request *v1.ListAppsRequest) (*v1.ListAppsResponse, error) {
	return s.appSvc.ListApps(ctx, request)
}

func (s *FiraApiService) GetApp(ctx context.Context, request *v1.GetAppRequest) (*v1.GetAppResponse, error) {
	return s.appSvc.GetApp(ctx, request)
}

func (s *FiraApiService) RotateAppToken(ctx context.Context, request *v1.RotateAppTokenRequest) (*v1.RotateAppTokenResponse, error) {
	return s.appSvc.RotateAppToken(ctx, request)
}

func (s *FiraApiService) InvalidateAppToken(ctx context.Context, request *v1.InvalidateAppTokenRequest) (*v1.InvalidateAppTokenResponse, error) {
	return s.appSvc.InvalidateAppToken(ctx, request)
}

///
/// Link Sessions API
///

func (s *FiraApiService) CreateLinkSession(ctx context.Context, request *v1.CreateLinkSessionRequest) (*v1.CreateLinkSessionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "not implemented")
}

func (s *FiraApiService) GetLinkSession(ctx context.Context, request *v1.GetLinkSessionRequest) (*v1.GetLinkSessionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "not implemented")
}

///
/// Financial aggregator API
///

func (s *FiraApiService) GetInstitutions(ctx context.Context, request *v1.GetInstitutionsRequest) (*v1.GetInstitutionsResponse, error) {
	return s.finAggSvc.GetInstitutions(ctx, request)
}
