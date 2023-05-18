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

type BigInterface interface {
	DoA() string
	DoB() int
	DoC() bool
}

type ImplA struct{}

func (ia *ImplA) DoA() string {
	return "Doing A"
}

type ImplB struct{}

func (ib *ImplB) DoB() int {
	return 123
}

type ImplC struct{}

func (ic *ImplC) DoC() bool {
	return true
}

type CombinedImpl struct {
	ImplA
	ImplB
	ImplC
}

func BS() {
	var bi BigInterface
	bi = &CombinedImpl{}
	bi.DoA()
	bi.DoB()
	bi.DoC()
}
