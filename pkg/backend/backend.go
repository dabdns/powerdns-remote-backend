package backend

type LookupResult struct {
	QType   string `json:"qType"`
	QName   string `json:"qname"`
	Content string `json:"content"`
	TTL     uint32 `json:"ttl"`
}

type DomainInfoResult struct {
	ID             uint32   `json:"id"`
	Zone           string   `json:"zone"`
	Masters        []string `json:"masters"`
	NotifiedSerial uint32   `json:"notified_serial"`
	Serial         uint32   `json:"serial"`
	LastCheck      uint32   `json:"last_check"`
	Kind           string   `json:"kind"`
}

// Backend @see https://doc.powerdns.com/authoritative/backends/remote.html
type Backend interface {
	// Always required:
	Initialize() bool
	Lookup(qtype string, qname string, zoneId string) (lookupResultArray []LookupResult, err error)

	// Filling the Zone Cache:
	GetAllDomains(includeDisabled bool) (domainInfoResultArray []DomainInfoResult, err error)
}
