package extract

import "errors"

type extractorFromXml struct {
	xmlBytes []byte
}

func FromXml(xmlBytes []byte) *extractorFromXml {
	return &extractorFromXml{xmlBytes: xmlBytes}
}

func (e *extractorFromXml) ExtractBase64TlvQr() (string, error) {
	return "", errors.New("Not implemented")
}
