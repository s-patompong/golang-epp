package frames

func Login() string {
	return `
        <?xml version="1.0" encoding="UTF-8"?>
          <epp xmlns="urn:ietf:params:xml:ns:epp-1.0"
               xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance"
               xsi:schemaLocation="urn:ietf:params:xml:ns:epp-1.0 epp-1.0.xsd">
            <command>
              <login>
                <clID>{username}</clID>
                <pw>{password}</pw>
                <options>
                  <version>{version}</version>
                  <lang>{lang}</lang>
                </options>
                <svcs>
                   {objects}
                   {extensions}
                </svcs>
              </login>
              <clTRID>EPP-{clTRID}</clTRID>
            </command>
          </epp>
  `
}
