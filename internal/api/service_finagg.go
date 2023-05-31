package api

import (
	"context"
	v1 "github.com/opencorelabs/fira/gen/protos/go/protos/fira/v1"
	"github.com/opencorelabs/fira/internal/backend"
)

type FinancialAggregatorService struct {
	backendProvider backend.Provider
}

func NewFinancialAggregatorService(backendProvider backend.Provider) *FinancialAggregatorService {
	return &FinancialAggregatorService{
		backendProvider: backendProvider,
	}
}

func (s *FinancialAggregatorService) GetInstitutions(ctx context.Context, request *v1.GetInstitutionsRequest) (*v1.GetInstitutionsResponse, error) {
	be := s.backendProvider.Backend()

	institutions, err := be.GetInstitutions(ctx, backend.InstitutionSearch{
		CountryCodes: request.CountryCodes,
		Search:       request.SearchString,
	})
	if err != nil {
		return nil, err
	}

	resp := &v1.GetInstitutionsResponse{}

	for _, institution := range institutions {
		resp.Institutions = append(resp.Institutions, &v1.Institution{
			Id:      institution.ID(),
			Name:    institution.Name(),
			Website: institution.URL(),
			Logo:    institution.LogoURL(),
			Oauth:   institution.UsesOAuth(),
		})
	}

	return resp, nil
}
