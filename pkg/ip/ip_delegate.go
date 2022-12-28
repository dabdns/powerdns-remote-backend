package ip

import "github.com/dabdns/powerdns-remote-backend/pkg/backend"

const (
	A          string = "A"
	AAAA       string = "AAAA"
	ANY        string = "ANY"
	SOA        string = "SOA"
	TXT        string = "TXT"
	defaultTTL int16  = 60
)

type DelegateIP struct {
	Domain string
}

func NewDelegateIP(domain string) *DelegateIP {
	return &DelegateIP{
		Domain: domain,
	}
}

func (*DelegateIP) Initialize() bool {
	return true
}

func (*DelegateIP) GetAllDomains(_ bool) (domainInfoResultArray []backend.DomainInfoResult, err error) {
	domainInfoResultArray = []backend.DomainInfoResult{}
	domainInfoResult := backend.DomainInfoResult{
		ID:             1,
		Zone:           "nip.kune.one.",
		Masters:        []string{"0.0.0.0"},
		NotifiedSerial: 1,
		Serial:         1,
		LastCheck:      0,
		Kind:           "native",
	}
	domainInfoResultArray = append(domainInfoResultArray, domainInfoResult)
	return
}
