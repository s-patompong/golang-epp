package nominetuk

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"io"
	"os"
)

func logXML(pfx string, p []byte) {
	if !LogXml {
		return
	}

	var b bytes.Buffer
	enc := xml.NewEncoder(&b)
	enc.Indent("", "\t")

	dec := xml.NewDecoder(bytes.NewReader(p))
	var t xml.Token
	var err error
	for {
		t, err = dec.RawToken()
		if err == io.EOF {
			err = enc.Flush()
			break
		}
		if err != nil {
			break
		}
		err = enc.EncodeToken(t)
		if err != nil {
			break
		}
	}
	if err != nil {
		fmt.Fprintf(os.Stdout, "Indentation error. Raw XML: %s\n%s\n\n", pfx, string(p))
		return
	}

	fmt.Fprintf(os.Stdout, "%s (pretty-printed)\n", pfx)
	io.Copy(os.Stdout, &b)
	fmt.Fprint(os.Stdout, "\n\n")
}
