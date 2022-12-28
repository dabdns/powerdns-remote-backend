package backend

import (
	"testing"
)

type TestBackend struct{}

func (TestBackend) Initialize() bool {
	return true
}

func (TestBackend) Lookup(_ string, _ string, _ string) (lookupResultArray []LookupResult, err error) {
	lookupResultArray = []LookupResult{}
	return
}

func (TestBackend) GetAllDomains(_ bool) (domainInfoResultArray []DomainInfoResult, err error) {
	domainInfoResultArray = []DomainInfoResult{}
	return
}

func TestNewBackendDelegate(t *testing.T) {
	backendDelegate := NewBackendDelegate()
	if backendDelegate == nil {
		t.FailNow()
	}
}

func TestDelegateAddDelegate(t *testing.T) {
	testBackend := TestBackend{}
	backendDelegate := NewBackendDelegate()
	backendDelegate.AddDelegate(testBackend)
	if len(backendDelegate.Backends) != 1 && backendDelegate.Backends[0] != testBackend {
		t.FailNow()
	}
}
