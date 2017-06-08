package nominetuk

import "github.com/nbio/xx"

type response_ struct {
	Result
	Greeting
	DomainInfoResponse
	DomainCheckResponse
}

var scanResponse = xx.NewScanner()

func init() {
	scanResponse.MustHandleStartElement("epp", func(c *xx.Context) error {
		*c.Value.(*response_) = response_{}
		return nil
	})
}
