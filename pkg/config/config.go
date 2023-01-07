package config

import (
	"embed"
	"encoding/json"
	"fmt"
	"github.com/dabdns/powerdns-remote-backend/pkg/util"
	"strings"
)

var (
	//go:embed config.json
	embedFS embed.FS
)

type Config struct {
	Connector *ConnectorConfig  `json:"connector"`
	Delegates []*DelegateConfig `json:"delegates"`
}

func (c *Config) Merge(o *Config) {
	if o != nil {
		if o.Connector != nil {
			if c.Connector != nil {
				c.Connector.Merge(o.Connector)
			} else {
				c.Connector = o.Connector
			}
		}
		if o.Delegates != nil {
			delegates := []*DelegateConfig{}
			i := 0
			for ; i < len(o.Delegates); i = i + 1 {
				if i < len(c.Delegates) {
					c.Delegates[i].Merge(o.Delegates[i])
					delegates = append(delegates, c.Delegates[i])
				} else {
					delegates = append(delegates, o.Delegates[i])
				}
			}
			for ; i < len(c.Delegates); i = i + 1 {
				delegates = append(delegates, c.Delegates[i])
			}
		}
	}
}

type ConnectorConfig struct {
	Type *string `json:"type"`
	Host *string `json:"host"`
	Port *uint16 `json:"port"`
}

func (c *ConnectorConfig) Merge(o *ConnectorConfig) {
	if o != nil {
		if o.Type != nil {
			c.Type = o.Type
		}
		if o.Host != nil {
			c.Host = o.Host
		}
		if o.Port != nil {
			c.Port = o.Port
		}
	}
}

type DelegateConfig struct {
	Domain               *string                             `json:"domain"`
	TTL                  *uint32                             `json:"ttl"`
	Initialize           *DelegateInitializeConfig           `json:"initialize"`
	Lookup               *DelegateLookupConfig               `json:"lookup"`
	GetAllDomains        *DelegateGetAllDomainsConfig        `json:"getAllDomains"`
	GetAllDomainMetadata *DelegateGetAllDomainMetadataConfig `json:"getAllDomainMetadata"`
}

func (d *DelegateConfig) Merge(o *DelegateConfig) {
	if o != nil {
		if o.Domain != nil {
			d.Domain = o.Domain
		}
		if o.Initialize != nil {
			d.Initialize.Merge(o.Initialize)
		}
		if o.Lookup != nil {
			d.Lookup.Merge(o.Lookup)
		}
		if o.GetAllDomains != nil {
			d.GetAllDomains.Merge(o.GetAllDomains)
		}
		if o.GetAllDomainMetadata != nil {
			d.GetAllDomainMetadata.Merge(o.GetAllDomainMetadata)
		}
	}
}

type DelegateInitializeConfig struct{}

func (*DelegateInitializeConfig) Merge(_ *DelegateInitializeConfig) {
	// No attributes
}

type DelegateLookupConfig struct {
	SOA       *DelegateLookupSOAConfig  `json:"soa"`
	A         *DelegateLookupAConfig    `json:"a"`
	AAAA      *DelegateLookupAAAAConfig `json:"aaaa"`
	NS        *DelegateLookupNSConfig   `json:"ns"`
	CNAME     *DelegateLookupNSConfig   `json:"cname"`
	DNAME     *DelegateLookupNSConfig   `json:"dname"`
	TXT       *DelegateLookupTXTConfig  `json:"txt"`
	Resolvers *[]string                 `json:"resolvers"`
}

func (d *DelegateLookupConfig) Merge(o *DelegateLookupConfig) {
	if o != nil {
		d.SOA.Merge(o.SOA)
		d.A.Merge(o.A)
		d.AAAA.Merge(o.AAAA)
		d.NS.Merge(o.NS)
		d.CNAME.Merge(o.CNAME)
		d.DNAME.Merge(o.DNAME)
		d.TXT.Merge(o.TXT)
	}
	d.MergeResolvers(o)
}

func (d *DelegateLookupConfig) MergeResolvers(o *DelegateLookupConfig) {
	if o != nil {
		if o.Resolvers != nil {
			if d.Resolvers != nil {
				resolvers := []string{}
				i := 0
				for ; i < len(*o.Resolvers); i = i + 1 {
					resolvers = append(resolvers, (*o.Resolvers)[i])
				}
				for ; i < len(*d.Resolvers); i = i + 1 {
					resolvers = append(resolvers, (*d.Resolvers)[i])
				}
				d.Resolvers = &resolvers
			} else {
				d.Resolvers = o.Resolvers
			}
		}
	}
}

type DelegateLookupSOAConfig struct {
	Default *DelegateLookupSOAObjectConfig            `json:"default"`
	Entries map[string]*DelegateLookupSOAObjectConfig `json:"entries"`
}

func (d *DelegateLookupSOAConfig) Merge(o *DelegateLookupSOAConfig) {
	if o != nil {
		if o.Default != nil {
			if d.Default != nil {
				d.Default.Merge(o.Default)
			} else {
				d.Default = o.Default
			}
		}
		if o.Entries != nil {
			if o.Entries != nil {
				for k, v := range o.Entries {
					d.Entries[k] = v
				}
			} else {
				d.Entries = o.Entries
			}
		}
	}
}

type DelegateLookupSOAObjectConfig struct {
	MNAME   *string `json:"mname"`
	RNAME   *string `json:"rname"`
	SERIAL  *uint32 `json:"serial"`
	REFRESH *uint32 `json:"refresh"`
	RETRY   *uint32 `json:"retry"`
	EXPIRE  *uint32 `json:"expire"`
	TTL     *uint32 `json:"ttl"`
}

func (d *DelegateLookupSOAObjectConfig) Merge(o *DelegateLookupSOAObjectConfig) {
	if o != nil {
		if o.MNAME != nil {
			d.MNAME = o.MNAME
		}
		if o.RNAME != nil {
			d.RNAME = o.RNAME
		}
		if o.SERIAL != nil {
			d.SERIAL = o.SERIAL
		}
		if o.REFRESH != nil {
			d.REFRESH = o.REFRESH
		}
		if o.RETRY != nil {
			d.RETRY = o.RETRY
		}
		if o.EXPIRE != nil {
			d.EXPIRE = o.EXPIRE
		}
		if o.TTL != nil {
			d.TTL = o.TTL
		}
	}
}

func (soaObjectConfig *DelegateLookupSOAObjectConfig) Serial() (serial string) {
	if soaObjectConfig.SERIAL != nil && *soaObjectConfig.SERIAL > 0 {
		serial = fmt.Sprintf("%02d", *soaObjectConfig.SERIAL)
	} else {
		now := util.GetDefaultTimeProvider().Now()
		serial = fmt.Sprintf("%s00",
			now.Format("20060102"),
		)
	}
	return
}

func (soaObjectConfig *DelegateLookupSOAObjectConfig) Content(ttl uint32) string {
	return fmt.Sprintf("%s %s %s %d %d %d %d",
		*soaObjectConfig.MNAME,
		*soaObjectConfig.RNAME,
		soaObjectConfig.Serial(),
		*soaObjectConfig.REFRESH,
		*soaObjectConfig.RETRY,
		*soaObjectConfig.EXPIRE,
		ttl,
	)
}

type DelegateLookupAConfig struct {
	Default *string            `json:"default"`
	Entries map[string]*string `json:"entries"`
}

func (d *DelegateLookupAConfig) Merge(o *DelegateLookupAConfig) {
	if o != nil {
		if o.Default != nil {
			d.Default = o.Default
		}
		if o.Entries != nil {
			if o.Entries != nil {
				for k, v := range o.Entries {
					d.Entries[k] = v
				}
			} else {
				d.Entries = o.Entries
			}
		}
	}
}

type DelegateLookupAAAAConfig struct {
	Default *string            `json:"default"`
	Entries map[string]*string `json:"entries"`
}

func (d *DelegateLookupAAAAConfig) Merge(o *DelegateLookupAAAAConfig) {
	if o != nil {
		if o.Default != nil {
			d.Default = o.Default
		}
		if o.Entries != nil {
			if o.Entries != nil {
				for k, v := range o.Entries {
					d.Entries[k] = v
				}
			} else {
				d.Entries = o.Entries
			}
		}
	}
}

type DelegateLookupTXTConfig struct {
	Default *string            `json:"default"`
	Entries map[string]*string `json:"entries"`
}

func (d *DelegateLookupTXTConfig) Merge(o *DelegateLookupTXTConfig) {
	if o != nil {
		if o.Default != nil {
			d.Default = o.Default
		}
		if o.Entries != nil {
			if o.Entries != nil {
				for k, v := range o.Entries {
					d.Entries[k] = v
				}
			} else {
				d.Entries = o.Entries
			}
		}
	}
}

type DelegateLookupNSConfig struct {
	Default *string            `json:"default"`
	Entries map[string]*string `json:"entries"`
}

func (d *DelegateLookupNSConfig) Merge(o *DelegateLookupNSConfig) {
	if o != nil {
		if o.Default != nil {
			d.Default = o.Default
		}
		if o.Entries != nil {
			if o.Entries != nil {
				for k, v := range o.Entries {
					d.Entries[k] = v
				}
			} else {
				d.Entries = o.Entries
			}
		}
	}
}

type DelegateLookupCNAMEConfig struct {
	Default *string            `json:"default"`
	Entries map[string]*string `json:"entries"`
}

func (d *DelegateLookupCNAMEConfig) Merge(o *DelegateLookupCNAMEConfig) {
	if o != nil {
		if o.Default != nil {
			d.Default = o.Default
		}
		if o.Entries != nil {
			if o.Entries != nil {
				for k, v := range o.Entries {
					d.Entries[k] = v
				}
			} else {
				d.Entries = o.Entries
			}
		}
	}
}

type DelegateLookupDNAMEConfig struct {
	Default *string            `json:"default"`
	Entries map[string]*string `json:"entries"`
}

func (d *DelegateLookupDNAMEConfig) Merge(o *DelegateLookupDNAMEConfig) {
	if o != nil {
		if o.Default != nil {
			d.Default = o.Default
		}
		if o.Entries != nil {
			if o.Entries != nil {
				for k, v := range o.Entries {
					d.Entries[k] = v
				}
			} else {
				d.Entries = o.Entries
			}
		}
	}
}

type DelegateGetAllDomainsConfig struct {
	Default []*DelegateGetAllDomainsObjectConfig            `json:"default"`
	Entries map[string][]*DelegateGetAllDomainsObjectConfig `json:"entries"`
}

func (d *DelegateGetAllDomainsConfig) Merge(o *DelegateGetAllDomainsConfig) {
	if o != nil {
		if o.Default != nil {
			d.Default = o.Default
		}
		if o.Entries != nil {
			if o.Entries != nil {
				for k, v := range o.Entries {
					d.Entries[k] = v
				}
			} else {
				d.Entries = o.Entries
			}
		}
	}
}

type DelegateGetAllDomainsObjectConfig struct {
	Id             *uint32   `json:"id"`
	Masters        *[]string `json:"masters"`
	NotifiedSerial *uint32   `json:"notifiedSerial"`
	Serial         *uint32   `json:"serial"`
	LastCheck      *uint32   `json:"lastCheck"`
	Kind           *string   `json:"kind"`
}

func (d *DelegateGetAllDomainsObjectConfig) Merge(o *DelegateGetAllDomainsObjectConfig) {
	if o != nil {
		if o.Id != nil {
			d.Id = o.Id
		}
		if o.Masters != nil {
			d.Masters = o.Masters
		}
		if o.NotifiedSerial != nil {
			d.NotifiedSerial = o.NotifiedSerial
		}
		if o.Serial != nil {
			d.Serial = o.Serial
		}
		if o.LastCheck != nil {
			d.LastCheck = o.LastCheck
		}
		if o.Kind != nil {
			d.Kind = o.Kind
		}
	}
}

type DelegateGetAllDomainMetadataConfig struct {
	Default []map[string][]*string            `json:"default"`
	Entries map[string][]map[string][]*string `json:"entries"`
}

func (d *DelegateGetAllDomainMetadataConfig) Merge(o *DelegateGetAllDomainMetadataConfig) {
	if o != nil {
		if o.Default != nil {
			d.Default = o.Default
		}
		if o.Entries != nil {
			if o.Entries != nil {
				for k, v := range o.Entries {
					d.Entries[strings.ToUpper(k)] = v
				}
			} else {
				d.Entries = o.Entries
			}
		}
	}
}

func GetDefaultConfig() (defaults Config) {
	// Unmarshalling default config json
	configData, err := embedFS.ReadFile("config.json")
	if err != nil {
		panic(fmt.Errorf("error reading embedded config.json, %s", err))
	}
	err = json.Unmarshal(configData, &defaults)
	if err != nil {
		panic(fmt.Errorf("error unmarshalling default config json, %s", err))
	}
	return
}

func (delegateConfig *DelegateConfig) AsMap() (settings map[string]interface{}, err error) {
	settings = make(map[string]interface{})
	var data []byte
	data, err = json.Marshal(delegateConfig)
	if err == nil {
		err = json.Unmarshal(data, &settings)
	}
	return
}
