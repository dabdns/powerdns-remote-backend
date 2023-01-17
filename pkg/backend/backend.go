package backend

type Request struct {
	Method     string                 `json:"method"`
	Parameters map[string]interface{} `json:"parameters"`
}

type Response struct {
	Result interface{} `json:"result"`
	Log    []string    `json:"log"` // logged in PowerDNS at loglevel info (6).
}

type LookupResult struct {
	QType   string `json:"qtype"`
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
	Service(req *Request, resp *Response) (err error)
	// Always required:
	Initialize() bool
	Lookup(qtype string, qname string, zoneId string) (lookupResultArray []LookupResult, err error)

	// Filling the Zone Cache:
	GetAllDomains(includeDisabled bool) (domainInfoResultArray []DomainInfoResult, err error)
	GetAllDomainMetadata(qname string) (metadata map[string][]string, err error)
}
