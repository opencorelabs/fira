package backend

import "context"

type InstitutionSearch struct {
	CountryCodes []string
	Search       string
}

type Interface interface {
	ID() string
	Name() string
	URL() string
	LogoURL() string
	GetInstitutions(ctx context.Context, search InstitutionSearch) ([]Institution, error)
}

type Provider interface {
	Backend() Interface
}
