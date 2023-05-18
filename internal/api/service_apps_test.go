package api

import (
	"fmt"
	v1 "github.com/opencorelabs/fira/gen/protos/go/protos/fira/v1"
	"time"
)

func (s *FiraApiSuite) Test_CreateApp() {
	acct, ctx := s.acctCtx()
	createTime := time.Now().Add(-time.Second)
	req := &v1.CreateAppRequest{
		Name: "test-app",
	}
	resp, respErr := s.api.CreateApp(ctx, req)
	s.Require().NoError(respErr, "failed to create app")
	s.Require().NotNil(resp, "response is nil")
	s.Require().NotNil(resp.App, "app is nil")
	s.Require().NotEmpty(resp.App.AppId, "app id is empty")
	s.Equal(req.Name, resp.App.Name, "app name is not equal")
	s.Len(resp.App.Tokens, 3, "expect 3 tokens")
	for _, tok := range resp.App.Tokens {
		s.NotEqual(v1.Environment_ENVIRONMENT_UNSPECIFIED, tok.Environment, "env is unspecified")
		s.NotEmpty(tok.Jwt, "jwt is empty")
		s.True(createTime.Before(tok.ValidUtil.AsTime()), "valid until should be in the future")
	}
	s.Equal(acct.ID, resp.App.Owner.Id, "account id is not equal")
	s.True(createTime.Before(resp.App.CreatedAt.AsTime()), "created at should be in the future")
	s.True(createTime.Before(resp.App.UpdatedAt.AsTime()), "updated at should be in the future")
}

func (s *FiraApiSuite) Test_ListApps_Empty() {
	_, ctx := s.acctCtx()
	req := &v1.ListAppsRequest{}
	resp, respErr := s.api.ListApps(ctx, req)
	s.Require().NoError(respErr, "failed to list apps")
	s.Require().NotNil(resp, "response is nil")
	s.Len(resp.Apps, 0, "expect 0 apps")
}

func (s *FiraApiSuite) Test_ListApps() {
	_, ctx := s.acctCtx()
	expectedIDs := make(map[string]bool)
	for i := 0; i < 10; i++ {
		req := &v1.CreateAppRequest{
			Name: fmt.Sprintf("test-app-%d", i),
		}
		resp, respErr := s.api.CreateApp(ctx, req)
		s.Require().NoError(respErr, "failed to create app")
		expectedIDs[resp.App.AppId] = true
	}

	req := &v1.ListAppsRequest{}
	resp, respErr := s.api.ListApps(ctx, req)
	s.Require().NoError(respErr, "failed to list apps")
	s.Require().NotNil(resp, "response is nil")

	s.Len(resp.Apps, 10, "expect 10 apps")

	for _, app := range resp.Apps {
		s.True(expectedIDs[app.AppId], "unexpected app id")
	}
}

func (s *FiraApiSuite) Test_GetApp() {
	acct, ctx := s.acctCtx()
	req := &v1.CreateAppRequest{
		Name: "test-app",
	}
	createResp, createRespErr := s.api.CreateApp(ctx, req)
	s.Require().NoError(createRespErr, "failed to create app")
	s.Require().NotNil(createResp, "response is nil")
	s.Require().NotNil(createResp.App, "app is nil")

	getReq := &v1.GetAppRequest{
		AppId: createResp.App.AppId,
	}
	getResp, getRespErr := s.api.GetApp(ctx, getReq)
	s.Require().NoError(getRespErr, "failed to get app")
	s.Require().NotNil(getResp, "response is nil")
	s.Require().NotNil(getResp.App, "app is nil")
	s.Equal(createResp.App.AppId, getResp.App.AppId, "app id is not equal")
	s.Equal(createResp.App.Name, getResp.App.Name, "app name is not equal")
	s.Equal(acct.ID, getResp.App.Owner.Id, "account id is not equal")
	s.Equal(createResp.App.CreatedAt, getResp.App.CreatedAt, "created at is not equal")
	s.Equal(createResp.App.UpdatedAt, getResp.App.UpdatedAt, "updated at is not equal")
}

func (s *FiraApiSuite) Test_RotateAppToken() {
	_, ctx := s.acctCtx()
	req := &v1.CreateAppRequest{
		Name: "test-app",
	}
	createResp, createRespErr := s.api.CreateApp(ctx, req)
	s.Require().NoError(createRespErr, "failed to create app")
	s.Require().NotNil(createResp, "response is nil")
	s.Require().NotNil(createResp.App, "app is nil")
	s.Require().Len(createResp.App.Tokens, 3, "expect 3 tokens")

	var productionToken *v1.AppToken
	for _, tok := range createResp.App.Tokens {
		if tok.Environment == v1.Environment_ENVIRONMENT_PRODUCTION {
			productionToken = tok
			break
		}
	}

	if productionToken == nil {
		s.Fail("production token not found")
		return
	}

	rotateReq := &v1.RotateAppTokenRequest{
		AppId:       createResp.App.AppId,
		Environment: v1.Environment_ENVIRONMENT_PRODUCTION,
	}
	rotateResp, rotateRespErr := s.api.RotateAppToken(ctx, rotateReq)
	s.Require().NoError(rotateRespErr, "failed to rotate app token")
	s.Require().NotNil(rotateResp, "response is nil")
	s.Require().NotNil(rotateResp.App, "app is nil")
	s.Require().Len(rotateResp.App.Tokens, 3, "expect 3 tokens")
	for _, tok := range rotateResp.App.Tokens {
		s.NotEmpty(tok.Jwt, "jwt is empty")
		if tok.Environment == v1.Environment_ENVIRONMENT_PRODUCTION {
			s.NotEqual(productionToken.Jwt, tok.Jwt, "jwt is not rotated")
			return
		}
	}
}

func (s *FiraApiSuite) Test_InvalidateAppToken() {
	s.T().Skip("not implemented")
}
