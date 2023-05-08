package backend

func (delegateBase *DelegateBase) List(qname string, domainId string, zoneid string) (listResultArray []ListResult, err error) {
	listResultArray = []ListResult{}
	var lookupArray []LookupResult
	lookupArray, err = delegateBase.lookupSOA(qname)
	if err == nil {
		for _, lookupSOA := range lookupArray {
			listResult := ListResult{
				QName:     lookupSOA.QName,
				QType:     lookupSOA.QType,
				Content:   lookupSOA.Content,
				TTL:       lookupSOA.TTL,
				DomainId:  &domainId,
				ScopeMask: nil,
				Auth:      nil,
			}
			listResultArray = append(listResultArray, listResult)
		}
	}
	lookupArray, err = delegateBase.lookupANY(qname)
	if err == nil {
		for _, lookupSOA := range lookupArray {
			listResult := ListResult{
				QName:     lookupSOA.QName,
				QType:     lookupSOA.QType,
				Content:   lookupSOA.Content,
				TTL:       lookupSOA.TTL,
				DomainId:  &domainId,
				ScopeMask: nil,
				Auth:      nil,
			}
			listResultArray = append(listResultArray, listResult)
		}
	}
	return
}
