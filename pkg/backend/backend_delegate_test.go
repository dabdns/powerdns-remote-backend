package backend

import (
	"testing"
)

type TestBackend struct{}

func (TestBackend) Initialize() bool {
	return true
}

func (TestBackend) Service(_ *Request, _ *Response) (err error) {
	return
}

func (TestBackend) Lookup(_ string, _ string, _ string) (lookupResultArray []LookupResult, err error) {
	lookupResultArray = []LookupResult{}
	return
}

func (TestBackend) GetAllDomains(_ bool) (domainInfoResultArray []DomainInfoResult, err error) {
	domainInfoResultArray = []DomainInfoResult{}
	return
}

func (TestBackend) GetAllDomainMetadata(_ string) (metadata map[string][]string, err error) {
	metadata = map[string][]string{}
	return
}

func (TestBackend) GetDomainMetadata(_ string) (metadata []string, err error) {
	metadata = []string{}
	return
}

func TestNewDelegate(t *testing.T) {
	backendDelegate := NewDelegate()
	if backendDelegate == nil {
		t.FailNow()
	}
}

func TestDelegateAddDelegate(t *testing.T) {
	testBackend := TestBackend{}
	backendDelegate := NewDelegate()
	backendDelegate.AddDelegate(testBackend)
	if len(backendDelegate.Backends) != 1 && backendDelegate.Backends[0] != testBackend {
		t.FailNow()
	}
}
