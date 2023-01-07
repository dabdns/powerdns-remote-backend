package backend

import (
	"github.com/dabdns/powerdns-remote-backend/pkg/config"
	"strings"
)

func (delegateBase *DelegateBase) Lookup(qtype string, qname string, _ string) (lookupResultArray []LookupResult, err error) {
	if strings.HasSuffix(qname, *delegateBase.Conf.Domain) {
		switch qtype {
		case ANY:
			lookupResultArray, err = delegateBase.lookupANY(qname)
		case SOA:
			lookupResultArray, err = delegateBase.lookupSOA(qname)
		case A:
			lookupResultArray, err = delegateBase.lookupA(qname)
		case AAAA:
			lookupResultArray, err = delegateBase.lookupAAAA(qname)
		case NS:
			lookupResultArray, err = delegateBase.lookupNS(qname)
		case CNAME:
			lookupResultArray, err = delegateBase.lookupCNAME(qname)
		case DNAME:
			lookupResultArray, err = delegateBase.lookupDNAME(qname)
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
	type lookupFuncType func(string) (lookupResultArray []LookupResult, err error)
	lookupFuncs := []lookupFuncType{
		delegateBase.lookupA,
		delegateBase.lookupAAAA,
		delegateBase.lookupNS,
		delegateBase.lookupCNAME,
		delegateBase.lookupDNAME,
		delegateBase.lookupTXT,
	}
	for _, lookupFunc := range lookupFuncs {
		lookupFuncResultArray, lookupFuncErr := lookupFunc(qname)
		if lookupFuncErr != nil {
			err = lookupFuncErr
			break
		} else {
			lookupResultArray = append(lookupResultArray, lookupFuncResultArray...)
		}
	}
	return
}

func (delegateBase *DelegateBase) lookupA(qname string) (lookupResultArray []LookupResult, err error) {
	lookupResultArray = []LookupResult{}
	a := delegateBase.Conf.Lookup.A.Entries[qname]
	if a == nil && strings.HasSuffix(qname, *delegateBase.Conf.Domain) {
		a = delegateBase.Conf.Lookup.A.Default
	}
	if a != nil {
		lookupResult := LookupResult{
			QType:   A,
			QName:   qname,
			Content: *a,
			TTL:     *delegateBase.Conf.TTL,
		}
		lookupResultArray = append(lookupResultArray, lookupResult)
	}
	return
}

func (delegateBase *DelegateBase) lookupAAAA(qname string) (lookupResultArray []LookupResult, err error) {
	lookupResultArray = []LookupResult{}
	lookupResult := LookupResult{
		QType:   AAAA,
		QName:   qname,
		Content: *delegateBase.Conf.Lookup.AAAA.Default,
		TTL:     *delegateBase.Conf.TTL,
	}
	lookupResultArray = append(lookupResultArray, lookupResult)
	return
}

func (delegateBase *DelegateBase) lookupSOA(qname string) (lookupResultArray []LookupResult, err error) {
	lookupResultArray = []LookupResult{}
	var soaObjectConfig *config.DelegateLookupSOAObjectConfig
	var found bool
	soaObjectConfig, found = delegateBase.Conf.Lookup.SOA.Entries[qname]
	if soaObjectConfig == nil {
		soaObjectConfig = delegateBase.Conf.Lookup.SOA.Default
	}
	if soaObjectConfig != nil && (found || strings.EqualFold(qname, *delegateBase.Conf.Domain)) {
		ttl := *delegateBase.Conf.TTL
		if soaObjectConfig.TTL != nil {
			ttl = *soaObjectConfig.TTL
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
		TTL:     *delegateBase.Conf.TTL,
	}
	lookupResultArray = append(lookupResultArray, lookupResult)
	return
}

func (delegateBase *DelegateBase) lookupNS(qname string) (lookupResultArray []LookupResult, err error) {
	lookupResultArray = []LookupResult{}
	ns := delegateBase.Conf.Lookup.NS.Entries[qname]
	if ns == nil && strings.HasSuffix(qname, *delegateBase.Conf.Domain) {
		ns = delegateBase.Conf.Lookup.NS.Default
	}
	if ns != nil {
		lookupResult := LookupResult{
			QType:   NS,
			QName:   qname,
			Content: *ns,
			TTL:     *delegateBase.Conf.TTL,
		}
		lookupResultArray = append(lookupResultArray, lookupResult)
	}
	return
}

func (delegateBase *DelegateBase) lookupCNAME(qname string) (lookupResultArray []LookupResult, err error) {
	lookupResultArray = []LookupResult{}
	cname := delegateBase.Conf.Lookup.CNAME.Entries[qname]
	if cname == nil && strings.HasSuffix(qname, *delegateBase.Conf.Domain) {
		cname = delegateBase.Conf.Lookup.CNAME.Default
	}
	if cname != nil {
		lookupResult := LookupResult{
			QType:   CNAME,
			QName:   qname,
			Content: *cname,
			TTL:     *delegateBase.Conf.TTL,
		}
		lookupResultArray = append(lookupResultArray, lookupResult)
	}
	return
}

func (delegateBase *DelegateBase) lookupDNAME(qname string) (lookupResultArray []LookupResult, err error) {
	lookupResultArray = []LookupResult{}
	dname := delegateBase.Conf.Lookup.DNAME.Entries[qname]
	if dname == nil && strings.HasSuffix(qname, *delegateBase.Conf.Domain) {
		dname = delegateBase.Conf.Lookup.DNAME.Default
	}
	if dname != nil {
		lookupResult := LookupResult{
			QType:   DNAME,
			QName:   qname,
			Content: *dname,
			TTL:     *delegateBase.Conf.TTL,
		}
		lookupResultArray = append(lookupResultArray, lookupResult)
	}
	return
}
