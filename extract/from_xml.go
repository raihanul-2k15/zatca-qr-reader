package extract

import (
	"encoding/xml"
	"errors"
)

type extractorFromXml struct {
	xmlBytes []byte
}

func FromXml(xmlBytes []byte) *extractorFromXml {
	return &extractorFromXml{xmlBytes: xmlBytes}
}

func (e *extractorFromXml) ExtractBase64TlvQr() (string, error) {
	var invoice struct {
		XMLName    xml.Name `xml:"Invoice"`
		AddDocRefs []struct {
			ID         string `xml:"ID"`
			Attachment struct {
				EmDocBinObj string `xml:"EmbeddedDocumentBinaryObject"`
			} `xml:"Attachment"`
		} `xml:"AdditionalDocumentReference"`
	}

	err := xml.Unmarshal(e.xmlBytes, &invoice)
	if err != nil {
		return "", err
	}

	for _, docRef := range invoice.AddDocRefs {
		if docRef.ID == "QR" {
			return docRef.Attachment.EmDocBinObj, nil
		}
	}

	return "", errors.New("QR not found in XML")
}
