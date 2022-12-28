package ip

import (
	"github.com/dabdns/powerdns-remote-backend/pkg/backend"
)

const (
	defaultAContent    string = "127.0.0.1"
	defaultAAAAContent string = "::1"
	defaultSOAContent  string = "ns1.kune.one. hostmaster.kune.one. 2012081600 7200 3600 1209600 3600"
	defaultTXTContent  string = "foobar"
	defaultQName       string = "nip.kune.one."
)

func (ipDelegate *DelegateIP) Lookup(qtype string, qname string, _ string) (lookupResultArray []backend.LookupResult, err error) {
	switch qtype {
	case A:
		lookupResultArray, err = ipDelegate.lookupA(qname)
	case AAAA:
		lookupResultArray, err = ipDelegate.lookupAAAA(qname)
	case ANY:
		lookupResultArray, err = ipDelegate.lookupANY(qname)
	case SOA:
		lookupResultArray, err = ipDelegate.lookupSOA(qname)
	case TXT:
		lookupResultArray, err = ipDelegate.lookupTXT(qname)
	default:
		lookupResultArray = []backend.LookupResult{}
	}
	return
}

func (ipDelegate *DelegateIP) lookupANY(qname string) (lookupResultArray []backend.LookupResult, err error) {
	lookupResultArray = []backend.LookupResult{}
	aLookupResultArray, aErr := ipDelegate.lookupA(qname)
	if aErr != nil {
		err = aErr
	} else {
		lookupResultArray = append(lookupResultArray, aLookupResultArray...)
		aaaaLookupResultArray, aaaaErr := ipDelegate.lookupAAAA(qname)
		if aaaaErr != nil {
			err = aaaaErr
		} else {
			lookupResultArray = append(lookupResultArray, aaaaLookupResultArray...)
			txtLookupResultArray, txtErr := ipDelegate.lookupTXT(qname)
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

func (*DelegateIP) lookupSOA(_ string) (lookupResultArray []backend.LookupResult, err error) {
	lookupResultArray = []backend.LookupResult{}
	lookupResult := backend.LookupResult{
		QType:   SOA,
		QName:   defaultQName,
		Content: defaultSOAContent,
		TTL:     defaultTTL,
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
