package nominetuk

import (
	"fmt"
	"github.com/nbio/xx"
	"golang-epp/frames"
)

// Encode domain info
func (domain *Domain) encodeDomainInfo(domainName string) error {
	err := writeToBuffer(&domain.conn.buf, frames.DomainInfo(), map[string]string{
		"domain": domainName,
	})
	return err
}

// Process domain info
func (domain *Domain) processDomainInfo(domainName string) (DomainInfoResponse, error) {
	var domainInfoResponse DomainInfoResponse

	err := domain.conn.flushDataUnit()
	if err != nil {
		return domainInfoResponse, err
	}
	var res response_
	err = domain.conn.readResponse(&res)
	res.DomainInfoResponse.Result = res.Result
	if err != nil {
		return res.DomainInfoResponse, err
	}

	return res.DomainInfoResponse, nil
}

type DomainInfoResponse struct {
	Result
	Name, Roid, Registrant, ClID, CrID, CrDate, UpID, UpDate, ExDate string
	Ns                                                               []string
	DomainInfoExtension
}

type DomainInfoExtension struct {
	RegStatus, IgnoredField string
}

func init() {
	// Normal response
	resData := "epp > response > resData > " + ObjDomain + " infData"
	scanResponse.MustHandleCharData(resData+" > name", func(c *xx.Context) error {
		info := &c.Value.(*response_).DomainInfoResponse
		info.Name = string(c.CharData)
		return nil
	})
	scanResponse.MustHandleCharData(resData+" > roid", func(c *xx.Context) error {
		info := &c.Value.(*response_).DomainInfoResponse
		info.Roid = string(c.CharData)
		return nil
	})
	scanResponse.MustHandleCharData(resData+" > registrant", func(c *xx.Context) error {
		info := &c.Value.(*response_).DomainInfoResponse
		info.Registrant = string(c.CharData)
		return nil
	})
	scanResponse.MustHandleCharData(resData+" > clID", func(c *xx.Context) error {
		info := &c.Value.(*response_).DomainInfoResponse
		info.ClID = string(c.CharData)
		return nil
	})
	scanResponse.MustHandleCharData(resData+" > crID", func(c *xx.Context) error {
		info := &c.Value.(*response_).DomainInfoResponse
		info.CrID = string(c.CharData)
		return nil
	})
	scanResponse.MustHandleCharData(resData+" > crDate", func(c *xx.Context) error {
		info := &c.Value.(*response_).DomainInfoResponse
		info.CrDate = string(c.CharData)
		return nil
	})
	scanResponse.MustHandleCharData(resData+" > upID", func(c *xx.Context) error {
		info := &c.Value.(*response_).DomainInfoResponse
		info.UpID = string(c.CharData)
		return nil
	})
	scanResponse.MustHandleCharData(resData+" > upDate", func(c *xx.Context) error {
		info := &c.Value.(*response_).DomainInfoResponse
		info.UpDate = string(c.CharData)
		return nil
	})
	scanResponse.MustHandleCharData(resData+" > exDate", func(c *xx.Context) error {
		info := &c.Value.(*response_).DomainInfoResponse
		info.ExDate = string(c.CharData)
		return nil
	})
	scanResponse.MustHandleCharData(resData+" > ns > hostObj", func(c *xx.Context) error {
		info := &c.Value.(*response_).DomainInfoResponse
		info.Ns = append(info.Ns, string(c.CharData))
		return nil
	})

	// Extensions response
	extData := "epp > response > extension"
	scanResponse.MustHandleCharData(extData+" > infData > reg-status", func(c *xx.Context) error {
		info := &c.Value.(*response_).DomainInfoResponse
		info.DomainInfoExtension.RegStatus = string(c.CharData)
		return nil
	})
	scanResponse.MustHandleCharData(extData+" > ignored-field", func(c *xx.Context) error {
		info := &c.Value.(*response_).DomainInfoResponse
		info.DomainInfoExtension.IgnoredField = trimString(string(c.CharData))
		return nil
	})
}
