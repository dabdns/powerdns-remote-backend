package ip

import (
	"github.com/dabdns/powerdns-remote-backend/pkg/backend"
)

const (
	defaultAContent    string = "127.0.0.1"
	defaultAAAAContent string = "::1"
	defaultTXTContent  string = "foobar"
)

func (delegateIP *DelegateIP) Lookup(qtype string, qname string, _ string) (lookupResultArray []backend.LookupResult, err error) {
	switch qtype {
	case A:
		lookupResultArray, err = delegateIP.lookupA(qname)
	case AAAA:
		lookupResultArray, err = delegateIP.lookupAAAA(qname)
	case ANY:
		lookupResultArray, err = delegateIP.lookupANY(qname)
	case SOA:
		lookupResultArray, err = delegateIP.lookupSOA(qname)
	case TXT:
		lookupResultArray, err = delegateIP.lookupTXT(qname)
	default:
		lookupResultArray = []backend.LookupResult{}
	}
	return
}

func (delegateIP *DelegateIP) lookupANY(qname string) (lookupResultArray []backend.LookupResult, err error) {
	lookupResultArray = []backend.LookupResult{}
	aLookupResultArray, aErr := delegateIP.lookupA(qname)
	if aErr != nil {
		err = aErr
	} else {
		lookupResultArray = append(lookupResultArray, aLookupResultArray...)
		aaaaLookupResultArray, aaaaErr := delegateIP.lookupAAAA(qname)
		if aaaaErr != nil {
			err = aaaaErr
		} else {
			lookupResultArray = append(lookupResultArray, aaaaLookupResultArray...)
			txtLookupResultArray, txtErr := delegateIP.lookupTXT(qname)
			if txtErr != nil {
				err = txtErr
			} else {
				lookupResultArray = append(lookupResultArray, txtLookupResultArray...)
			}
		}
	}
	return
}

func (*DelegateIP) lookupA(qname string) (lookupResultArray []backend.LookupResult, err error) {
	lookupResultArray = []backend.LookupResult{}
	lookupResult := backend.LookupResult{
		QType:   A,
		QName:   qname,
		Content: defaultAContent,
		TTL:     defaultTTL,
	}
	lookupResultArray = append(lookupResultArray, lookupResult)
	return
}

func (*DelegateIP) lookupAAAA(qname string) (lookupResultArray []backend.LookupResult, err error) {
	lookupResultArray = []backend.LookupResult{}
	lookupResult := backend.LookupResult{
		QType:   AAAA,
		QName:   qname,
		Content: defaultAAAAContent,
		TTL:     defaultTTL,
	}
	lookupResultArray = append(lookupResultArray, lookupResult)
	return
}

func (delegateIP *DelegateIP) lookupSOA(_ string) (lookupResultArray []backend.LookupResult, err error) {
	lookupResultArray = []backend.LookupResult{}
	lookupResult := backend.LookupResult{
		QType:   SOA,
		QName:   delegateIP.Config.Domain,
		Content: delegateIP.Config.SOAConfig.Content(),
		TTL:     delegateIP.Config.SOAConfig.TTL,
	}
	lookupResultArray = append(lookupResultArray, lookupResult)
	return
}

func (*DelegateIP) lookupTXT(qname string) (lookupResultArray []backend.LookupResult, err error) {
	lookupResultArray = []backend.LookupResult{}
	lookupResult := backend.LookupResult{
		QType:   TXT,
		QName:   qname,
		Content: defaultTXTContent,
		TTL:     defaultTTL,
	}
	lookupResultArray = append(lookupResultArray, lookupResult)
	return
}
