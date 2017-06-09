package nominetuk

// System setting
const (
	EppUrl = "epp.nominet.org.uk:700"
	LogXml = false
)

// Frame config
const (
	// Frame file name and folder config
	FrameDir        = "frames"
	FrameLogin      = "login.xml"
	FrameDomainInfo = "domain-info.xml"

	// Frame variable config
	FrameVariableOpen  = "{"
	FrameVariableClose = "}"
)

// Object config
const (
	ObjDomain     = "urn:ietf:params:xml:ns:domain-1.0"
	ObjHost       = "urn:ietf:params:xml:ns:host-1.0"
	ObjContact    = "urn:ietf:params:xml:ns:contact-1.0"
	ObjFinance    = "http://www.unitedtld.com/epp/finance-1.0"
	ExtSecDNS     = "urn:ietf:params:xml:ns:secDNS-1.1"
	ExtRGP        = "urn:ietf:params:xml:ns:rgp-1.0"
	ExtLaunch     = "urn:ietf:params:xml:ns:launch-1.0"
	ExtIDN        = "urn:ietf:params:xml:ns:idn-1.0"
	ExtCharge     = "http://www.unitedtld.com/epp/charge-1.0"
	ExtFee05      = "urn:ietf:params:xml:ns:fee-0.5"
	ExtFee06      = "urn:ietf:params:xml:ns:fee-0.6"
	ExtFee07      = "urn:ietf:params:xml:ns:fee-0.7"
	ExtFee08      = "urn:ietf:params:xml:ns:fee-0.8"
	ExtFee09      = "urn:ietf:params:xml:ns:fee-0.9"
	ExtFee11      = "urn:ietf:params:xml:ns:fee-0.11"
	ExtPrice      = "urn:ar:params:xml:ns:price-1.1"
	ExtNamestore  = "http://www.verisign-grs.com/epp/namestoreExt-1.1"
	ExtNeulevel   = "urn:ietf:params:xml:ns:neulevel"
	ExtNeulevel10 = "urn:ietf:params:xml:ns:neulevel-1.0"
)
