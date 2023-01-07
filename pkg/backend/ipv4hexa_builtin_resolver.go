package backend

import (
	"fmt"
	"net"
	"regexp"
	"strconv"
	"strings"
)

const IPV4HEXA_BUILTIN_RESOLVER_NAME = "builtin:ipv4hexa"

type IPv4HexaResolver struct {
	settings map[string]interface{}
	regexp   *regexp.Regexp
	ttl      uint32
}

func NewIPv4HexaResolver(settings map[string]interface{}) *IPv4HexaResolver {
	var domain, pattern string
	var ttl uint32
	domain, _ = settings["domain"].(string)
	ttl, _ = settings["ttl"].(uint32)
	pattern = fmt.Sprintf("([0-9a-fA-F]{8})\\.%s$", strings.ReplaceAll(domain, ".", "\\."))
	re, _ := regexp.Compile(pattern)
	return &IPv4HexaResolver{
		settings: settings,
		regexp:   re,
		ttl:      ttl,
	}
}

func (resolver *IPv4HexaResolver) Lookup(qtype string, qname string, _ string) (lookupResultArray []LookupResult, err error) {
	lookupResultArray = []LookupResult{}
	if qtype == A || qtype == AAAA || qtype == ANY {
		a := resolver.regexp.FindStringSubmatch(qname)
		if len(a) > 1 {
			ipValue, convErr := strconv.ParseInt(a[1], 16, 33)
			if convErr == nil {
				var a, b, c, d byte
				a = byte(ipValue & 0xFF000000 >> 24)
				b = byte(ipValue & 0x00FF0000 >> 16)
				c = byte(ipValue & 0x0000FF00 >> 8)
				d = byte(ipValue & 0x000000FF)
				ip := net.IPv4(a, b, c, d)
				if ip != nil {
					var lookupResult LookupResult
					if qtype == A || qtype == ANY {
						lookupResult = LookupResult{
							QType:   A,
							QName:   qname,
							Content: ip.String(),
							TTL:     uint32(60),
						}
					} else {
						lookupResult = LookupResult{
							QType:   AAAA,
							QName:   qname,
							Content: fmt.Sprintf("::ffff:%s", ip.String()),
							TTL:     uint32(60),
						}
					}
					lookupResultArray = append(lookupResultArray, lookupResult)
				}
			}
		}
	}
	return
}
