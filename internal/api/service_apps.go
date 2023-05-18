package api

import (
	"context"
	"fmt"
	v1 "github.com/opencorelabs/fira/gen/protos/go/protos/fira/v1"
	"github.com/opencorelabs/fira/internal/auth"
	"github.com/opencorelabs/fira/internal/developer"
	"github.com/opencorelabs/fira/internal/logging"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
	"time"
)

type AppService struct {
	logger           *zap.SugaredLogger
	appJWTManager    auth.JWTManager
	appStoreProvider developer.AppStoreProvider
}

func NewAppService(
	log logging.Provider,
	appJWTManager auth.JWTManager,
	appStoreProvider developer.AppStoreProvider,
) *AppService {
	return &AppService{
		logger:           log.Logger().Named("api-service").Sugar(),
		appJWTManager:    appJWTManager,
		appStoreProvider: appStoreProvider,
	}
}

func (s *AppService) CreateApp(ctx context.Context, request *v1.CreateAppRequest) (*v1.CreateAppResponse, error) {
	acct, ok := auth.AccountFromContext(ctx)
	if !ok {
		return nil, status.Errorf(codes.Unauthenticated, "unauthenticated")
	}
	app := &developer.App{
		Name:      request.Name,
		AccountID: acct.ID,
		Tokens:    make(map[developer.Environment][]developer.Token),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	// create initial tokens
	for env := range developer.TokenExpiryMap {
		if err := app.Rotate(env); err != nil {
			s.logger.Errorw("failed to rotate app token", "error", err)
			return nil, status.Errorf(codes.Internal, "internal error")
		}
	}
	err := s.appStoreProvider.AppStore().CreateApp(ctx, app)
	if err != nil {
		s.logger.Errorw("failed to create app", "error", err)
		return nil, status.Errorf(codes.Internal, "internal error")
	}

	apiApp, convertErr := s.appToApi(ctx, app, acct)
	if convertErr != nil {
		s.logger.Errorw("failed to convert app to api", "error", convertErr)
		return nil, status.Errorf(codes.Internal, "internal error")
	}

	resp := &v1.CreateAppResponse{
		App: apiApp,
	}

	return resp, nil
}

func (s *AppService) ListApps(ctx context.Context, request *v1.ListAppsRequest) (*v1.ListAppsResponse, error) {
	acct, ok := auth.AccountFromContext(ctx)
	if !ok {
		return nil, status.Errorf(codes.Unauthenticated, "unauthenticated")
	}

	resp := &v1.ListAppsResponse{}

	apps, getErr := s.appStoreProvider.AppStore().GetAppsByAccountID(ctx, acct.ID)
	if getErr != nil {
		s.logger.Errorw("failed to get apps by account id", "error", getErr)
		return nil, status.Errorf(codes.Internal, "internal error")
	}

	for _, app := range apps {
		apiApp, convertErr := s.appToApi(ctx, app, acct)
		if convertErr != nil {
			s.logger.Errorw("failed to convert app to api", "error", convertErr)
			return nil, status.Errorf(codes.Internal, "internal error")
		}
		resp.Apps = append(resp.Apps, apiApp)
	}

	return resp, nil
}

func (s *AppService) GetApp(ctx context.Context, request *v1.GetAppRequest) (*v1.GetAppResponse, error) {
	acct, ok := auth.AccountFromContext(ctx)
	if !ok {
		return nil, status.Errorf(codes.Unauthenticated, "unauthenticated")
	}

	app, getErr := s.appStoreProvider.AppStore().GetAppByID(ctx, request.AppId)
	if getErr != nil {
		s.logger.Errorw("failed to get app by id", "error", getErr)
		return nil, status.Errorf(codes.Internal, "internal error")
	}

	if app == nil || app.AccountID != acct.ID {
		return nil, status.Errorf(codes.NotFound, "not found")
	}

	apiApp, convertErr := s.appToApi(ctx, app, acct)
	if convertErr != nil {
		s.logger.Errorw("failed to convert app to api", "error", convertErr)
		return nil, status.Errorf(codes.Internal, "internal error")
	}

	resp := &v1.GetAppResponse{
		App: apiApp,
	}

	return resp, nil
}

func (s *AppService) RotateAppToken(ctx context.Context, request *v1.RotateAppTokenRequest) (*v1.RotateAppTokenResponse, error) {
	acct, ok := auth.AccountFromContext(ctx)
	if !ok {
		return nil, status.Errorf(codes.Unauthenticated, "unauthenticated")
	}

	app, getErr := s.appStoreProvider.AppStore().GetAppByID(ctx, request.AppId)
	if getErr != nil {
		s.logger.Errorw("failed to get app by id", "error", getErr)
		return nil, status.Errorf(codes.Internal, "internal error")
	}

	if app == nil || app.AccountID != acct.ID {
		return nil, status.Errorf(codes.NotFound, "not found")
	}

	var devEnv developer.Environment
	switch request.Environment {
	case v1.Environment_ENVIRONMENT_SANDBOX:
		devEnv = developer.EnvironmentSandbox
	case v1.Environment_ENVIRONMENT_PRODUCTION:
		devEnv = developer.EnvironmentProduction
	case v1.Environment_ENVIRONMENT_DEVELOPER:
		devEnv = developer.EnvironmentDevelopment
	default:
		return nil, status.Errorf(codes.InvalidArgument, "invalid environment")
	}

	if rotErr := app.Rotate(devEnv); rotErr != nil {
		s.logger.Errorw("failed to rotate app token", "error", rotErr)
		return nil, status.Errorf(codes.Internal, "internal error")
	}

	if updErr := s.appStoreProvider.AppStore().UpdateApp(ctx, app); updErr != nil {
		s.logger.Errorw("failed to update app", "error", updErr)
		return nil, status.Errorf(codes.Internal, "internal error")
	}

	apiApp, convertErr := s.appToApi(ctx, app, acct)
	if convertErr != nil {
		s.logger.Errorw("failed to convert app to api", "error", convertErr)
		return nil, status.Errorf(codes.Internal, "internal error")
	}

	return &v1.RotateAppTokenResponse{
		App: apiApp,
	}, nil
}

func (s *AppService) InvalidateAppToken(ctx context.Context, request *v1.InvalidateAppTokenRequest) (*v1.InvalidateAppTokenResponse, error) {
	acct, ok := auth.AccountFromContext(ctx)
	if !ok {
		return nil, status.Errorf(codes.Unauthenticated, "unauthenticated")
	}

	app, getErr := s.appStoreProvider.AppStore().GetAppByID(ctx, request.AppId)
	if getErr != nil {
		s.logger.Errorw("failed to get app by id", "error", getErr)
		return nil, status.Errorf(codes.Internal, "internal error")
	}

	if app == nil || app.AccountID != acct.ID {
		return nil, status.Errorf(codes.NotFound, "not found")
	}

	if invErr := app.Invalidate(request.Jwt); invErr != nil {
		s.logger.Errorw("failed to invalidate app token", "error", invErr)
		return nil, status.Errorf(codes.Internal, "internal error")
	}

	return &v1.InvalidateAppTokenResponse{}, nil
}

func (s *AppService) appToApi(ctx context.Context, app *developer.App, account *auth.Account) (*v1.App, error) {
	var tokens []*v1.AppToken
	for env := range developer.TokenExpiryMap {
		token := app.Tokens[env][len(app.Tokens[env])-1]
		envCtx := developer.WithEnvironment(ctx, env)
		jwt, jwtErr := s.appJWTManager.Generate(envCtx, app)
		if jwtErr != nil {
			return nil, jwtErr
		}
		var apiEnv v1.Environment
		switch env {
		case developer.EnvironmentDevelopment:
			apiEnv = v1.Environment_ENVIRONMENT_DEVELOPER
		case developer.EnvironmentSandbox:
			apiEnv = v1.Environment_ENVIRONMENT_SANDBOX
		case developer.EnvironmentProduction:
			apiEnv = v1.Environment_ENVIRONMENT_PRODUCTION
		default:
			return nil, fmt.Errorf("unmapped environment: %s", env)
		}
		tokens = append(tokens, &v1.AppToken{
			Environment: apiEnv,
			Jwt:         jwt,
			ValidUtil:   timestamppb.New(token.ValidUntil),
		})
	}

	return &v1.App{
		AppId:     app.ID,
		Name:      app.Name,
		Owner:     accountToApi(account),
		Tokens:    tokens,
		CreatedAt: timestamppb.New(app.CreatedAt),
		UpdatedAt: timestamppb.New(app.UpdatedAt),
	}, nil
}
