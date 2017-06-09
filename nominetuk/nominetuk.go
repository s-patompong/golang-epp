package nominetuk

type NominetUk struct {
	Conn *Conn
}

func (nominet *NominetUk) NewDomain() Domain {
	var domain Domain
	domain.conn = nominet.Conn
	return domain
}
