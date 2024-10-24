package xml

import (
	"bytes"
	"encoding/xml"
)

var (
	Unmarshal  = xml.Unmarshal
	Marshal    = xml.Marshal
	NewDecoder = xml.NewDecoder
	NewEncoder = xml.NewEncoder
)

func MarshalToString(v any) string {
	str, err := marshal(v)
	if err != nil {
		return ""
	}
	return string(str)
}

func MustToString(v any) string {
	str, err := marshal(v)
	if err != nil {
		panic(err)
	}
	return string(str)
}

func marshal(v any) ([]byte, error) {
	buf := new(bytes.Buffer)
	err := xml.NewEncoder(buf).Encode(v)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}
