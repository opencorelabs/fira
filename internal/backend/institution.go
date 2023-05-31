package backend

type Institution interface {
	ID() string
	Name() string
	URL() string
	LogoURL() string
	CountryCodes() []string
	UsesOAuth() bool
}
