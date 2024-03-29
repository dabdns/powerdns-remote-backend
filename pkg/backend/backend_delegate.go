package backend

type Delegate struct {
	Backends []Backend
}

func NewDelegate() *Delegate {
	return &Delegate{
		Backends: []Backend{},
	}
}

func (delegate *Delegate) AddDelegate(backend Backend) {
	delegate.Backends = append(delegate.Backends, backend)
}

func (delegate *Delegate) Service(req *Request, resp *Response) (err error) {
	for _, backend := range delegate.Backends {
		err = backend.Service(req, resp)
	}
	return
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

func (delegate *Delegate) List(qname string, domainId string, zoneId string) (listResultArray []ListResult, err error) {
	listResultArray = []ListResult{}
	for _, backend := range delegate.Backends {
		backendListResultArray, backendErr := backend.List(qname, domainId, zoneId)
		if backendErr != nil {
			err = backendErr
		} else {
			listResultArray = append(listResultArray, backendListResultArray...)
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

func (delegate *Delegate) GetAllDomainMetadata(qname string) (metadata map[string][]string, err error) {
	metadata = map[string][]string{}
	for _, backend := range delegate.Backends {
		backendMetadata, backendErr := backend.GetAllDomainMetadata(qname)
		if backendErr != nil {
			err = backendErr
		} else {
			metadata = backendMetadata
		}
	}
	return
}

func (delegate *Delegate) GetDomainMetadata(qname string, qtype string) (metadata []string, err error) {
	metadata = []string{}
	for _, backend := range delegate.Backends {
		backendMetadata, backendErr := backend.GetDomainMetadata(qname, qtype)
		if backendErr != nil {
			err = backendErr
		} else {
			metadata = backendMetadata
		}
	}
	return
}

func (delegate *Delegate) GetDomainInfo(qname string) (domainInfoResult DomainInfoResult, err error) {
	for _, backend := range delegate.Backends {
		backendDomainInfoResult, backendErr := backend.GetDomainInfo(qname)
		if backendErr != nil {
			err = backendErr
		} else {
			domainInfoResult = backendDomainInfoResult
		}
	}
	return
}
