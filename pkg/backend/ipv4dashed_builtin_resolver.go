package backend

import (
	"fmt"
	"net"
	"regexp"
	"strings"
)

const IPV4DASHED_BUILTIN_RESOLVER_NAME = "builtin:ipv4dashed"

type IPv4DashedResolver struct {
	settings map[string]interface{}
	regexp   *regexp.Regexp
	ttl      uint32
}

func NewIPv4DashedResolver(settings map[string]interface{}) *IPv4DashedResolver {
	var domain, pattern string
	var ttl uint32
	domain, _ = settings["domain"].(string)
	ttl, _ = settings["ttl"].(uint32)
	pattern = fmt.Sprintf("(([0-9]{1,3}-){3}([0-9]{1,3}))\\.%s$", strings.ReplaceAll(domain, ".", "\\."))
	re, _ := regexp.Compile(pattern)
	return &IPv4DashedResolver{
		settings: settings,
		regexp:   re,
		ttl:      ttl,
	}
}

func (resolver *IPv4DashedResolver) Lookup(qtype string, qname string, _ string) (lookupResultArray []LookupResult, err error) {
	lookupResultArray = []LookupResult{}
	if qtype == A || qtype == AAAA || qtype == ANY {
		a := resolver.regexp.FindStringSubmatch(qname)
		if len(a) > 1 {
			ip := net.ParseIP(strings.ReplaceAll(a[1], "-", "."))
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
	return
}
