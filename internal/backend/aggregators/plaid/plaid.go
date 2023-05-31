package plaid

import (
	"context"
	"encoding/base64"
	"fmt"
	"github.com/opencorelabs/fira/internal/backend"
	"github.com/opencorelabs/fira/internal/logging"
	"github.com/plaid/plaid-go/v12/plaid"
	"go.uber.org/zap"
	"time"
)

const (
	getInstitutionsTimeout   = 30 * time.Second
	getInstitutionFetchLimit = 500
)

var logoSVG = `<svg height="48" viewBox="0 0 126 48" xmlns="http://www.w3.org/2000/svg"><g fill="#111" fill-rule="evenodd"><path d="M66.248 16.268c-1.057-.889-2.861-1.333-5.413-1.333h-5.756v17.788h4.304v-5.575h1.928c2.34 0 4.056-.515 5.148-1.546 1.23-1.155 1.849-2.693 1.849-4.613 0-1.991-.687-3.565-2.06-4.721m-5.044 6.855h-1.821V18.96h1.636c1.99 0 2.985.698 2.985 2.094 0 1.378-.934 2.068-2.8 2.068M75.673 14.934h-4.488v17.788h9.69v-4.026h-5.202zM89.668 14.934l-7.05 17.788h4.832l.924-2.586H94.5l.845 2.586h4.886l-7-17.788zm-.053 11.601l1.849-6.08 1.82 6.08z"></path><path d="M102.473 32.722h4.489V14.934h-4.489zM124.39 18.268a7.376 7.376 0 00-2.14-2.053c-1.355-.854-3.204-1.28-5.545-1.28h-5.914v17.787h6.918c2.5 0 4.506-.817 6.02-2.453 1.514-1.635 2.27-3.805 2.27-6.508 0-2.15-.537-3.981-1.61-5.493m-7.182 10.427h-1.927v-9.734h1.954c1.373 0 2.428.43 3.168 1.287.74.857 1.11 2.073 1.11 3.647 0 3.2-1.435 4.8-4.305 4.8M18.637 0L4.09 3.81.081 18.439l5.014 5.148L0 28.65l3.773 14.693 14.484 4.047 5.096-5.064 5.014 5.147 14.547-3.81 4.008-14.63-5.013-5.146 5.095-5.063L43.231 4.13 28.745.083l-5.094 5.063zM9.71 6.624l7.663-2.008 3.351 3.44-4.887 4.856zm16.822 1.478l3.405-3.383 7.63 2.132-6.227 6.187zM4.672 17.238l2.111-7.705 6.125 6.288-4.886 4.856zm29.547-1.243l6.227-6.189 1.986 7.74-3.404 3.384zm-15.502-.127l4.887-4.856 4.807 4.936-4.886 4.856zm-7.814 7.765l4.886-4.856 4.81 4.936-4.888 4.856zm15.503.127l4.886-4.856L36.1 23.84l-4.887 4.856zM4.57 29.927l3.406-3.385 4.807 4.937-6.225 6.186zm14.021 1.598l4.887-4.856 4.808 4.936-4.886 4.856zm15.502.128l4.887-4.856 3.351 3.439-2.11 7.705zm-24.656 8.97l6.226-6.189 4.81 4.936-3.406 3.385zm16.843-1.206l4.886-4.856 6.126 6.289-7.662 2.007z"></path></g></svg>`

type Backend struct {
	logoDataURL string
	cli         *plaid.APIClient
	logger      *zap.Logger
}

type Env plaid.Environment

func New(lp logging.Provider, clientID, secret, env string) *Backend {
	be := &Backend{
		logger: lp.Logger().Named("plaid-backend"),
	}

	be.logoDataURL = "data:image/svg+xml;base64," + base64.StdEncoding.EncodeToString([]byte(logoSVG))

	cfg := plaid.NewConfiguration()
	cfg.AddDefaultHeader("PLAID-CLIENT-ID", clientID)
	cfg.AddDefaultHeader("PLAID-SECRET", secret)
	cfg.UseEnvironment(plaid.Environment(env))

	be.cli = plaid.NewAPIClient(cfg)

	return be
}

func (b *Backend) ID() string {
	return "plaid"
}

func (b *Backend) Name() string {
	return "Plaid"
}

func (b *Backend) URL() string {
	return "https://plaid.com/"
}

func (b *Backend) LogoURL() string {
	return b.logoDataURL
}

func (b *Backend) GetInstitutions(ctx context.Context, search backend.InstitutionSearch) ([]backend.Institution, error) {
	timeoutCtx, cancel := context.WithTimeout(ctx, getInstitutionsTimeout)
	defer cancel()

	var institutions []backend.Institution

	now := time.Now()

	req := plaid.NewInstitutionsSearchRequest(search.Search, []plaid.Products{plaid.PRODUCTS_BALANCE}, b.plaidCountryCodes(search.CountryCodes))
	req.SetOptions(plaid.InstitutionsSearchRequestOptions{
		IncludeOptionalMetadata: plaid.PtrBool(true),
	})

	b.logger.Debug("fetching institutions", zap.Any("req", req))

	resp, _, err := b.cli.PlaidApi.InstitutionsSearch(timeoutCtx).InstitutionsSearchRequest(*req).Execute()

	b.logger.Debug("fetched institutions", zap.Duration("duration", time.Since(now)), zap.Int("returned", len(resp.Institutions)))

	if err != nil {
		return nil, fmt.Errorf("failed to get institutions: %w", err)
	}

	for _, inst := range resp.Institutions {
		institutions = append(institutions, &Institution{inst: inst})
	}

	return institutions, nil
}

// assert Backend implements backend.Interface
var _ backend.Interface = (*Backend)(nil)

func (b *Backend) plaidCountryCodes(wanted []string) []plaid.CountryCode {
	knownCodes := []string{
		"US",
		"GB",
		"ES",
		"NL",
		"FR",
		"IE",
		"CA",
		"DE",
		"IT",
		"PL",
		"DK",
		"NO",
		"SE",
		"EE",
		"LT",
		"LV",
		"PT",
	}
	if len(wanted) == 0 {
		wanted = knownCodes
	}
	var result []plaid.CountryCode
	for _, code := range wanted {
		plaidCountryCode, err := plaid.NewCountryCodeFromValue(code)
		if err != nil {
			b.logger.Error("failed to convert country code", zap.String("code", code), zap.Error(err))
			continue
		}
		result = append(result, *plaidCountryCode)
	}
	return result
}

type Institution struct {
	inst plaid.Institution
}

func (i *Institution) ID() string {
	return i.inst.GetInstitutionId()
}

func (i *Institution) Name() string {
	return i.inst.GetName()
}

func (i *Institution) URL() string {
	return i.inst.GetUrl()
}

func (i *Institution) LogoURL() string {
	return i.inst.GetLogo()
}

func (i *Institution) CountryCodes() []string {
	ret := make([]string, len(i.inst.GetCountryCodes()))
	codes := i.inst.GetCountryCodes()
	for ii, code := range codes {
		ret[ii] = string(code)
	}
	return ret
}

func (i *Institution) UsesOAuth() bool {
	return i.inst.GetOauth()
}

// assert Institution implements backend.Institution
var _ backend.Institution = (*Institution)(nil)
