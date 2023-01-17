package backend

import "github.com/dabdns/powerdns-remote-backend/pkg/config"

const (
	ANY   string = "ANY"
	SOA   string = "SOA"
	A     string = "A"
	AAAA  string = "AAAA"
	NS    string = "NS"
	CNAME string = "CNAME"
	DNAME string = "DNAME"
	TXT   string = "TXT"

	METHOD_INITIALIZE           string = "initialize"
	METHOD_LOOKUP               string = "lookup"
	METHOD_GETALLDOMAINS        string = "getalldomains"
	METHOD_GETALLDOMAINMETADATA string = "getalldomainmetadata"

	PARAM_QTYPE string = "qtype"
	PARAM_QNAME string = "qname"
	//PARAM_REMOTE     string = "remote"
	//PARAM_LOCAL      string = "local"
	//PARAM_REALREMOTE string = "real-remote"
	//PARAM_ZONEID     string = "zone-id"
)

type DelegateBase struct {
	Conf      *config.DelegateConfig
	resolvers []Resolver
}

func NewDelagateBase(conf *config.DelegateConfig) (delegate *DelegateBase, err error) {
	resolvers := []Resolver{}
	if conf.Lookup.Resolvers != nil {
		for _, resolverUri := range *conf.Lookup.Resolvers {
			var resolver Resolver
			var settings map[string]interface{}
			settings, err = conf.AsMap()
			if err == nil {
				resolver, err = NewResolver(resolverUri, settings)
				if err == nil {
					resolvers = append(resolvers, resolver)
				} else {
					break
				}
			} else {
				break
			}
		}
	}
	if err == nil {
		delegate = &DelegateBase{
			Conf:      conf,
			resolvers: resolvers,
		}
	}
	return
}

func (delegateBase *DelegateBase) Service(req *Request, resp *Response) (err error) {
	switch req.Method {
	case METHOD_INITIALIZE:
		resp.Result = delegateBase.Initialize()
	case METHOD_LOOKUP:
		qtype, qtypeOK := req.Parameters[PARAM_QTYPE].(string)
		qname, qnameOK := req.Parameters[PARAM_QNAME].(string)
		//zoneid, zoneidOK := req.Parameters[PARAM_ZONEID].(float64)
		if qtypeOK && qnameOK {
			resp.Result, err = delegateBase.Lookup(qtype, qname, "")
		}
	case METHOD_GETALLDOMAINS:
		resp.Result, err = delegateBase.GetAllDomains(false)
	case METHOD_GETALLDOMAINMETADATA:
		qname, qnameOK := req.Parameters[PARAM_QNAME].(string)
		if qnameOK {
			resp.Result, err = delegateBase.GetAllDomainMetadata(qname)
		}
	default:
	}
	return
}

func (*DelegateBase) Initialize() bool {
	return true
}

func (delegateBase *DelegateBase) GetAllDomains(_ bool) (domainInfoResultArray []DomainInfoResult, err error) {
	domainInfoResultArray = []DomainInfoResult{}
	for _, getAllDomainsConfig := range delegateBase.Conf.GetAllDomains.Default {
		domainInfoResult := DomainInfoResult{
			ID:             *getAllDomainsConfig.Id,
			Zone:           *delegateBase.Conf.Domain,
			Masters:        *getAllDomainsConfig.Masters,
			NotifiedSerial: *getAllDomainsConfig.NotifiedSerial,
			Serial:         *getAllDomainsConfig.Serial,
			LastCheck:      *getAllDomainsConfig.LastCheck,
			Kind:           *getAllDomainsConfig.Kind,
		}
		domainInfoResultArray = append(domainInfoResultArray, domainInfoResult)
	}
	return
}

func (delegateBase *DelegateBase) GetAllDomainMetadata(qname string) (metadata map[string][]string, err error) {
	metadata = map[string][]string{}
	if qname == *delegateBase.Conf.Domain {
		metadata["PRESIGNED"] = []string{"0"}
	}
	return
}
