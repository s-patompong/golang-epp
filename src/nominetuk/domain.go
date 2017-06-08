package nominetuk

type Domain struct {
	conn *Conn
}

// Get domain info
func (domain *Domain) Info(domainName string) (DomainInfoResponse, error) {
	var domainInfoResponse DomainInfoResponse
	err := domain.encodeDomainInfo(domainName)
	if err != nil {
		return domainInfoResponse, err
	}
	return domain.processDomainInfo(domainName)
}
