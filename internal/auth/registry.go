package auth

import "sync"

type DefaultRegistry struct {
	mtx      *sync.Mutex
	backends map[CredentialsType]Backend
}

func NewDefaultRegistry() Registry {
	return &DefaultRegistry{
		backends: make(map[CredentialsType]Backend),
		mtx:      &sync.Mutex{},
	}
}

func (d *DefaultRegistry) RegisterBackend(ct CredentialsType, backend Backend) {
	d.mtx.Lock()
	defer d.mtx.Unlock()
	d.backends[ct] = backend
}

func (d *DefaultRegistry) GetBackend(ct CredentialsType) (Backend, error) {
	d.mtx.Lock()
	defer d.mtx.Unlock()
	backend, ok := d.backends[ct]
	if !ok {
		return nil, ErrUnknownAuthorizer
	}
	return backend, nil
}
