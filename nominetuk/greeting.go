package nominetuk

import (
	"encoding/xml"
	"github.com/nbio/xx"
)

// Hello sends a <hello> command to request a <greeting> from the EPP server.
func (c *Conn) Hello() error {
	err := c.writeDataUnit(xmlHello)
	if err != nil {
		return err
	}
	_, err = c.readGreeting()
	return err
}

var xmlHello = []byte(xml.Header + startEPP + `<hello/>` + endEPP)

// Greeting is an EPP response that represents server status and capabilities.
// https://tools.ietf.org/html/rfc5730#section-2.4
type Greeting struct {
	ServerName string   `xml:"svID"`
	Versions   []string `xml:"svcMenu>version"`
	Languages  []string `xml:"svcMenu>lang"`
	Objects    []string `xml:"svcMenu>objURI"`
	Extensions []string `xml:"svcMenu>svcExtension>extURI,omitempty"`
}

// SupportsExtension returns true if the EPP server supports
// the object specified by uri.
func (g *Greeting) SupportsObject(uri string) bool {
	for _, v := range g.Objects {
		if v == uri {
			return true
		}
	}
	return false
}

// SupportsExtension returns true if the EPP server supports
// the extension specified by uri.
func (g *Greeting) SupportsExtension(uri string) bool {
	for _, v := range g.Extensions {
		if v == uri {
			return true
		}
	}
	return false
}

// ExtURNNames maps short extension names to their full URN.
var ExtURNNames = map[string]string{
	"secDNS-1.1":       ExtSecDNS,
	"rgp-1.0":          ExtRGP,
	"launch-1.0":       ExtLaunch,
	"idn-1.0":          ExtIDN,
	"charge-1.0":       ExtCharge,
	"fee-0.5":          ExtFee05,
	"fee-0.6":          ExtFee06,
	"fee-0.7":          ExtFee07,
	"fee-0.8":          ExtFee08,
	"fee-0.9":          ExtFee09,
	"fee-0.11":         ExtFee11,
	"price-1.1":        ExtPrice,
	"namestoreExt-1.1": ExtNamestore,
	"neulevel":         ExtNeulevel,
	"neulevel-1.0":     ExtNeulevel10,
}

func (c *Conn) readGreeting() (Greeting, error) {
	err := c.readDataUnit()
	if err != nil {
		return Greeting{}, err
	}
	deleteBufferRange(&c.buf, []byte(`<dcp>`), []byte(`</dcp>`))
	var res response_
	err = IgnoreEOF(scanResponse.Scan(c.decoder, &res))
	if err != nil {
		return Greeting{}, err
	}
	return res.Greeting, nil
}

func init() {
	path := "epp>greeting"
	scanResponse.MustHandleCharData(path+">svID", func(c *xx.Context) error {
		res := c.Value.(*response_)
		res.Greeting.ServerName = string(c.CharData)
		return nil
	})
	scanResponse.MustHandleCharData(path+">svcMenu>version", func(c *xx.Context) error {
		res := c.Value.(*response_)
		res.Greeting.Versions = append(res.Greeting.Versions, string(c.CharData))
		return nil
	})
	scanResponse.MustHandleCharData(path+">svcMenu>lang", func(c *xx.Context) error {
		res := c.Value.(*response_)
		res.Greeting.Languages = append(res.Greeting.Languages, string(c.CharData))
		return nil
	})
	scanResponse.MustHandleCharData(path+">svcMenu>objURI", func(c *xx.Context) error {
		res := c.Value.(*response_)
		res.Greeting.Objects = append(res.Greeting.Objects, string(c.CharData))
		return nil
	})
	scanResponse.MustHandleCharData(path+">svcMenu>svcExtension>extURI", func(c *xx.Context) error {
		res := c.Value.(*response_)
		res.Greeting.Extensions = append(res.Greeting.Extensions, string(c.CharData))
		return nil
	})
}
