package nominetuk

import (
	"bytes"
	"golang-epp/frames"
)

// Login initializes an authenticated EPP session.
// https://tools.ietf.org/html/rfc5730#section-2.9.1.1
func (c *Conn) Login(user, password string) error {
	err := c.writeLogin(user, password)
	if err != nil {
		return err
	}
	var res response_
	return c.readResponse(&res)
}

func (c *Conn) writeLogin(user, password string) error {
	ver, lang := "1.0", "en"
	if len(c.Greeting.Versions) > 0 {
		ver = c.Greeting.Versions[0]
	}
	if len(c.Greeting.Languages) > 0 {
		lang = c.Greeting.Languages[0]
	}
	err := encodeLogin(&c.buf, user, password, ver, lang)
	if err != nil {
		return err
	}
	return c.flushDataUnit()
}

func encodeLogin(buf *bytes.Buffer, user, password, version, language string) error {
	err := writeToBuffer(buf, frames.Login(), map[string]string{
		"username":   user,
		"password":   password,
		"version":    version,
		"lang":       language,
		"objects":    getObjectsStringForFrame(NominetUkObjects),
		"extensions": getExtensionStringForFrame(NominetUkExtensions),
	})

	return err
}

// Logout sends a <logout> command to terminate an EPP session.
// https://tools.ietf.org/html/rfc5730#section-2.9.1.2
func (c *Conn) Logout() error {
	err := c.writeDataUnit(xmlLogout)
	if err != nil {
		return err
	}
	var res response_
	return c.readResponse(&res)
}

var xmlLogout = []byte(xmlCommandPrefix + `<logout/>` + xmlCommandSuffix)
