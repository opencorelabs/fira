package backend

import (
	"context"
	"fmt"
	"github.com/opencorelabs/fira/internal/logging"
	"go.uber.org/zap"
)

// Meta is a provider that aggregates other interfaces.
type Meta struct {
	Interfaces []Interface
	logger     *zap.Logger
}

// assert Meta implements Provider
var _ Interface = (*Meta)(nil)

func NewMetaInterface(lp logging.Provider, backends []Interface) *Meta {
	return &Meta{
		Interfaces: backends,
		logger:     lp.Logger().Named("meta-backend"),
	}
}

func (m *Meta) ID() string {
	return "meta"
}

func (m *Meta) Name() string {
	return "Fira"
}

func (m *Meta) URL() string {
	return "https://fira.opencorelabs.org"
}

func (m *Meta) LogoURL() string {
	return "https://fira.opencorelabs.org/logo.png"
}

func (m *Meta) GetInstitutions(ctx context.Context, search InstitutionSearch) ([]Institution, error) {
	var institutions []Institution
	for _, backend := range m.Interfaces {
		providerInstitutions, err := backend.GetInstitutions(ctx, search)
		if err != nil {
			return nil, err
		}
		metaInstitutions := make([]Institution, len(providerInstitutions))
		for i, institution := range providerInstitutions {
			metaInstitutions[i] = &MetaInstitution{
				backend:     backend,
				institution: institution,
			}
		}
		institutions = append(institutions, metaInstitutions...)
	}
	return institutions, nil
}

type MetaInstitution struct {
	backend     Interface
	institution Institution
}

func (m *MetaInstitution) ID() string {
	return fmt.Sprintf("%s:%s", m.backend.ID(), m.institution.ID())
}

func (m *MetaInstitution) Name() string {
	return m.institution.Name()
}

func (m *MetaInstitution) ProviderID() string {
	return m.backend.ID()
}

func (m *MetaInstitution) URL() string {
	return m.institution.URL()
}

func (m *MetaInstitution) LogoURL() string {
	return m.institution.LogoURL()
}

func (m *MetaInstitution) CountryCodes() []string {
	return m.institution.CountryCodes()
}

func (m *MetaInstitution) UsesOAuth() bool {
	return m.institution.UsesOAuth()
}

// assert MetaInstitution implements Institution
var _ Institution = (*MetaInstitution)(nil)
