package config

import (
	_ "embed"
	"fmt"
	"github.com/dabdns/powerdns-remote-backend/pkg/util"
)

var (
	//go:embed config.json
	DefaultConfString string
)

type Config struct {
	Connector *ConnectorConfig `json:"connector"`
	Delegates []DelegateConfig `json:"delegates"`
}

type ConnectorConfig struct {
	Type string `json:"type"`
	Host string `json:"host"`
	Port uint16 `json:"port"`
}

type DelegateConfig struct {
	Domain               string                              `json:"domain"`
	TTL                  uint32                              `json:"ttl"`
	Initialize           *DelegateInitializeConfig           `json:"initialize"`
	Lookup               *DelegateLookupConfig               `json:"lookup"`
	GetAllDomains        *DelegateGetAllDomainsConfig        `json:"getAllDomains"`
	GetAllDomainMetadata *DelegateGetAllDomainMetadataConfig `json:"getAllDomainMetadata"`
}

type DelegateInitializeConfig struct{}

type DelegateLookupConfig struct {
	SOA  *DelegateLookupSOAConfig  `json:"soa"`
	A    *DelegateLookupAConfig    `json:"a"`
	AAAA *DelegateLookupAAAAConfig `json:"aaaa"`
	TXT  *DelegateLookupTXTConfig  `json:"txt"`
}

type DelegateLookupSOAConfig struct {
	Default *DelegateLookupSOAObjectConfig            `json:"default"`
	Entries map[string]*DelegateLookupSOAObjectConfig `json:"entries"`
}
type DelegateLookupSOAObjectConfig struct {
	MNAME   string `json:"mname"`
	RNAME   string `json:"rname"`
	SERIAL  uint32 `json:"serial"`
	REFRESH uint32 `json:"refresh"`
	RETRY   uint32 `json:"retry"`
	EXPIRE  uint32 `json:"expire"`
	TTL     uint32 `json:"ttl"`
}

func (soaObjectConfig *DelegateLookupSOAObjectConfig) Serial() (serial string) {
	if soaObjectConfig.SERIAL > 0 {
		serial = fmt.Sprintf("%02d", soaObjectConfig.SERIAL)
	} else {
		now := util.GetDefaultTimeProvider().Now()
		serial = fmt.Sprintf("%s%02d",
			now.Format("20060102"),
			soaObjectConfig.SERIAL,
		)
	}
	return
}

func (soaObjectConfig *DelegateLookupSOAObjectConfig) Content(ttl uint32) string {
	return fmt.Sprintf("%s %s %s %d %d %d %d",
		soaObjectConfig.MNAME,
		soaObjectConfig.RNAME,
		soaObjectConfig.Serial(),
		soaObjectConfig.REFRESH,
		soaObjectConfig.RETRY,
		soaObjectConfig.EXPIRE,
		ttl,
	)
}

type DelegateLookupAConfig struct {
	Default string            `json:"default"`
	Entries map[string]string `json:"entries"`
}

type DelegateLookupAAAAConfig struct {
	Default string            `json:"default"`
	Entries map[string]string `json:"entries"`
}

type DelegateLookupTXTConfig struct {
	Default string            `json:"default"`
	Entries map[string]string `json:"entries"`
}
type DelegateGetAllDomainsConfig struct {
	Default []*DelegateGetAllDomainsObjectConfig            `json:"default"`
	Entries map[string][]*DelegateGetAllDomainsObjectConfig `json:"entries"`
}

type DelegateGetAllDomainsObjectConfig struct {
	Id             uint32   `json:"id"`
	Masters        []string `json:"masters"`
	NotifiedSerial uint32   `json:"notifiedSerial"`
	Serial         uint32   `json:"serial"`
	LastCheck      uint32   `json:"lastCheck"`
	Kind           string   `json:"kind"`
}

type DelegateGetAllDomainMetadataConfig struct {
	Default []map[string][]string            `json:"default"`
	Entries map[string][]map[string][]string `json:"entries"`
}
