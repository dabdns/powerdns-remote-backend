package backend

import (
	"github.com/dabdns/powerdns-remote-backend/pkg/config"
	"strings"
)

func (delegateBase *DelegateBase) Lookup(qtype string, qname string, _ string) (lookupResultArray []LookupResult, err error) {
	if strings.HasSuffix(qname, delegateBase.Conf.Domain) {
		switch qtype {
		case A:
			lookupResultArray, err = delegateBase.lookupA(qname)
		case AAAA:
			lookupResultArray, err = delegateBase.lookupAAAA(qname)
		case ANY:
			lookupResultArray, err = delegateBase.lookupANY(qname)
		case SOA:
			lookupResultArray, err = delegateBase.lookupSOA(qname)
		case TXT:
			lookupResultArray, err = delegateBase.lookupTXT(qname)
		default:
			lookupResultArray = []LookupResult{}
		}
	}
	return
}

func (delegateBase *DelegateBase) lookupANY(qname string) (lookupResultArray []LookupResult, err error) {
	lookupResultArray = []LookupResult{}
	aLookupResultArray, aErr := delegateBase.lookupA(qname)
	if aErr != nil {
		err = aErr
	} else {
		lookupResultArray = append(lookupResultArray, aLookupResultArray...)
		aaaaLookupResultArray, aaaaErr := delegateBase.lookupAAAA(qname)
		if aaaaErr != nil {
			err = aaaaErr
		} else {
			lookupResultArray = append(lookupResultArray, aaaaLookupResultArray...)
			txtLookupResultArray, txtErr := delegateBase.lookupTXT(qname)
			if txtErr != nil {
				err = txtErr
			} else {
				lookupResultArray = append(lookupResultArray, txtLookupResultArray...)
			}
		}
	}
	return
}

func (delegateBase *DelegateBase) lookupA(qname string) (lookupResultArray []LookupResult, err error) {
	lookupResultArray = []LookupResult{}
	lookupResult := LookupResult{
		QType:   A,
		QName:   qname,
		Content: delegateBase.Conf.Lookup.A.Default,
		TTL:     delegateBase.Conf.TTL,
	}
	lookupResultArray = append(lookupResultArray, lookupResult)
	return
}

func (delegateBase *DelegateBase) lookupAAAA(qname string) (lookupResultArray []LookupResult, err error) {
	lookupResultArray = []LookupResult{}
	lookupResult := LookupResult{
		QType:   AAAA,
		QName:   qname,
		Content: delegateBase.Conf.Lookup.AAAA.Default,
		TTL:     delegateBase.Conf.TTL,
	}
	lookupResultArray = append(lookupResultArray, lookupResult)
	return
}

func (delegateBase *DelegateBase) lookupSOA(qname string) (lookupResultArray []LookupResult, err error) {
	lookupResultArray = []LookupResult{}
	var soaObjectConfig *config.DelegateLookupSOAObjectConfig
	soaObjectConfig = delegateBase.Conf.Lookup.SOA.Entries[qname]
	if soaObjectConfig == nil {
		soaObjectConfig = delegateBase.Conf.Lookup.SOA.Default
		// FIXME: rewrite qname with default Domain name
		qname = delegateBase.Conf.Domain
	}
	if soaObjectConfig != nil {
		ttl := delegateBase.Conf.TTL
		if soaObjectConfig.TTL > 0 {
			ttl = soaObjectConfig.TTL
		}
		lookupResult := LookupResult{
			QType:   SOA,
			QName:   qname,
			Content: soaObjectConfig.Content(ttl),
			TTL:     ttl,
		}
		lookupResultArray = append(lookupResultArray, lookupResult)
	}
	return
}

func (delegateBase *DelegateBase) lookupTXT(qname string) (lookupResultArray []LookupResult, err error) {
	lookupResultArray = []LookupResult{}
	lookupResult := LookupResult{
		QType:   TXT,
		QName:   qname,
		Content: "~",
		TTL:     delegateBase.Conf.TTL,
	}
	lookupResultArray = append(lookupResultArray, lookupResult)
	return
}
