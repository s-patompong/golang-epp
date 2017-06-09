package nominetuk

import (
	"bytes"
	"strings"
)

// Write the map to buffer
func writeToBuffer(buf *bytes.Buffer, xml string, m map[string]string) error {
	buf.Reset()

	m["clTRID"] = generateclTRID()

	xml = replaceStringInFrame(xml, m)

	buf.WriteString(xml)

	return nil
}

func getObjectsStringForFrame(objects []string) string {
	str := ""
	for _, object := range objects {
		str += "<objURI>" + object + "</objURI>"
	}
	return str
}

func getExtensionStringForFrame(extensions []string) string {
	str := ""

	if len(extensions) > 0 {
		str += "<svcExtension>"
	}

	for _, extension := range extensions {
		str += "<extURI>" + extension + "</extURI>"
	}

	if len(extensions) > 0 {
		str += "</svcExtension>"
	}

	return str
}

func replaceStringInFrame(stringXml string, m map[string]string) string {
	for key, value := range m {
		stringXml = strings.Replace(stringXml, FrameVariableOpen+key+FrameVariableClose, value, -1)
	}
	return stringXml
}
