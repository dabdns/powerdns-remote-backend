package ip

import "github.com/dabdns/powerdns-remote-backend/pkg/backend"

const (
	A          string = "A"
	AAAA       string = "AAAA"
	ANY        string = "ANY"
	SOA        string = "SOA"
	TXT        string = "TXT"
	defaultTTL uint32 = 60
)

type DelegateIP struct {
	Config *DelegateIPConfig
}

type DelegateIPConfig struct {
	Domain    string
	SOAConfig SOAConfig
}

func NewDelegateIP(domain string, soaConfig SOAConfig) *DelegateIP {
	return &DelegateIP{
		Config: &DelegateIPConfig{
			Domain:    domain,
			SOAConfig: soaConfig,
		},
	}
}

func (*DelegateIP) Initialize() bool {
	return true
}

func (delegateIP *DelegateIP) GetAllDomains(_ bool) (domainInfoResultArray []backend.DomainInfoResult, err error) {
	domainInfoResultArray = []backend.DomainInfoResult{}
	domainInfoResult := backend.DomainInfoResult{
		ID:             1,
		Zone:           delegateIP.Config.Domain,
		Masters:        []string{"0.0.0.0"},
		NotifiedSerial: 1,
		Serial:         1,
		LastCheck:      0,
		Kind:           "native",
	}
	domainInfoResultArray = append(domainInfoResultArray, domainInfoResult)
	return
}
