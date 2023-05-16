package developer

import (
	"context"
	"fmt"
	"github.com/opencorelabs/fira/internal/auth"
)

type AppJWTManager struct {
}

func NewAppJWTManager() auth.JWTManager {
	return &AppJWTManager{}
}

func (a *AppJWTManager) Generate(ctx context.Context, principal interface{}) (string, error) {
	_, isApp := principal.(*App)
	if !isApp {
		return "", fmt.Errorf("principal is not an app")
	}
	panic("implement me")
}

func (a *AppJWTManager) Verify(ctx context.Context, tokenStr string) (context.Context, error) {
	//TODO implement me
	panic("implement me")
}
