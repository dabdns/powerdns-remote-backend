package backend

import (
	"fmt"
	"plugin"
	"strings"
)

const (
	builtinPrefix string = "builtin:"
)

type Resolver interface {
	Lookup(qtype string, qname string, zoneId string) (lookupResultArray []LookupResult, err error)
}

type NewResolverType func(map[string]interface{}) Resolver

type BuiltinResolverError struct {
	msg string
}

func (err BuiltinResolverError) Error() string {
	return err.msg
}

func BuiltinResolver(name string, settings map[string]interface{}) (resolver Resolver, err error) {
	switch name {
	case IPV4DOTTED_BUILTIN_RESOLVER_NAME:
		resolver = NewIPv4DottedResolver(settings)
	case IPV4DASHED_BUILTIN_RESOLVER_NAME:
		resolver = NewIPv4DashedResolver(settings)
	case IPV4HEXA_BUILTIN_RESOLVER_NAME:
		resolver = NewIPv4HexaResolver(settings)
	case IPV4B32_BUILTIN_RESOLVER_NAME:
		resolver = NewIPv4B32Resolver(settings)
	case IPV4B32CROCKFORD_BUILTIN_RESOLVER_NAME:
		resolver = NewIPv4B32CrockfordResolver(settings)
	default:
		err = BuiltinResolverError{
			msg: fmt.Sprintf("No resolver with name \"%s\"", name),
		}
	}
	return
}

func LoadResolver(uri string, settings map[string]interface{}) (resolver Resolver, err error) {
	var p *plugin.Plugin
	var s plugin.Symbol
	p, err = plugin.Open(uri)
	if err == nil {
		s, err = p.Lookup("NewResolver")
		if err == nil {
			newResolver, isNewResolverType := s.(NewResolverType)
			if isNewResolverType {
				resolver = newResolver(settings)
			}
		}
	}
	return
}

func NewResolver(uri string, settings map[string]interface{}) (resolver Resolver, err error) {
	if strings.HasPrefix(uri, builtinPrefix) {
		resolver, err = BuiltinResolver(uri, settings)
	} else {
		resolver, err = LoadResolver(uri, settings)
	}
	return
}
