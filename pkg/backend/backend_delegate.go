package backend

type Delegate struct {
	Backends []Backend
}

func NewBackendDelegate() *Delegate {
	return &Delegate{
		Backends: []Backend{},
	}
}

func (delegate *Delegate) AddDelegate(backend Backend) {
	delegate.Backends = append(delegate.Backends, backend)
}

func (delegate *Delegate) Initialize() bool {
	var initialized = true
	for _, backend := range delegate.Backends {
		initialized = initialized && backend.Initialize()
	}
	return initialized
}

func (delegate *Delegate) Lookup(qtype string, qname string, zoneId string) (lookupResultArray []LookupResult, err error) {
	lookupResultArray = []LookupResult{}
	for _, backend := range delegate.Backends {
		backendLookupResultArray, backendErr := backend.Lookup(qtype, qname, zoneId)
		if backendErr != nil {
			err = backendErr
		} else {
			lookupResultArray = append(lookupResultArray, backendLookupResultArray...)
		}
	}
	return
}

func (delegate *Delegate) GetAllDomains(includeDisabled bool) (domainInfoResultArray []DomainInfoResult, err error) {
	domainInfoResultArray = []DomainInfoResult{}
	for _, backend := range delegate.Backends {
		backendDomainInfoResultArray, backendErr := backend.GetAllDomains(includeDisabled)
		if backendErr != nil {
			err = backendErr
		} else {
			domainInfoResultArray = append(domainInfoResultArray, backendDomainInfoResultArray...)
		}
	}
	return
}
