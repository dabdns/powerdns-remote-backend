package ip

import "fmt"

const (
	defaultSERIAL  uint32 = 01
	defaultREFRESH uint32 = 86400
	defaultRETRY   uint32 = 7200
	defaultEXPIRE  uint32 = 86400
)

type SOAConfig struct {
	MNAME   string
	RNAME   string
	SERIAL  uint32
	REFRESH uint32
	RETRY   uint32
	EXPIRE  uint32
	TTL     uint32
}

func NewSOAConfig(mname string, rname string) SOAConfig {
	return SOAConfig{
		MNAME:   mname,
		RNAME:   rname,
		SERIAL:  defaultSERIAL,
		REFRESH: defaultREFRESH,
		RETRY:   defaultRETRY,
		EXPIRE:  defaultEXPIRE,
		TTL:     defaultTTL,
	}
}

func (soaConfig SOAConfig) Content() string {
	return fmt.Sprintf("%s %s %d %d %d %d %d",
		soaConfig.MNAME,
		soaConfig.RNAME,
		soaConfig.SERIAL,
		soaConfig.REFRESH,
		soaConfig.RETRY,
		soaConfig.EXPIRE,
		soaConfig.TTL,
	)
}
