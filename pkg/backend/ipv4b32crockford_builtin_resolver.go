package backend

import (
	"encoding/base32"
	"fmt"
	"net"
	"regexp"
	"strings"
)

// Base32 alphabets
const (
	IPV4B32CROCKFORD_BUILTIN_RESOLVER_NAME string = "builtin:ipv4b32crockford"
	lowercaseAlphabet                      string = "0123456789abcdefghjkmnpqrstvwxyz"
)

// Base32 encodings
var (
	lowerEncoding = base32.NewEncoding(lowercaseAlphabet).WithPadding(base32.NoPadding)
)

type IPv4B32CrockfordResolver struct {
	settings map[string]interface{}
	regexp   *regexp.Regexp
	ttl      uint32
}

func NewIPv4B32CrockfordResolver(settings map[string]interface{}) *IPv4B32CrockfordResolver {
	var domain, pattern string
	var ttl uint32
	domain, _ = settings["domain"].(string)
	ttl, _ = settings["ttl"].(uint32)
	pattern = fmt.Sprintf("([0123456789abcdefghjkmnpqrstvwxyzABCDEFGHJKMNPQRSTVWXYZ]{5,7})\\.%s$", strings.ReplaceAll(domain, ".", "\\."))
	re, _ := regexp.Compile(pattern)
	return &IPv4B32CrockfordResolver{
		settings: settings,
		regexp:   re,
		ttl:      ttl,
	}
}

func (resolver *IPv4B32CrockfordResolver) Lookup(qtype string, qname string, _ string) (lookupResultArray []LookupResult, err error) {
	lookupResultArray = []LookupResult{}
	if qtype == A || qtype == AAAA || qtype == ANY {
		a := resolver.regexp.FindStringSubmatch(qname)
		if len(a) > 1 {
			lower := strings.ToLower(a[1])
			if len(lower) == 6 {
				// add padding
				lower = fmt.Sprintf("0%s", lower)
			}
			ipValue, convErr := lowerEncoding.DecodeString(lower)
			if convErr == nil && len(ipValue) == 4 {
				var a, b, c, d byte
				a = ipValue[0]
				b = ipValue[1]
				c = ipValue[2]
				d = ipValue[3]
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
