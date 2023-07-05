package zatcaqr

import (
	"encoding/base64"
	"encoding/binary"
	"time"

	"github.com/pauloavelar/go-tlv/tlv"
)

type ZatcaQr struct {
	SellerName   string
	VatRegNo     string
	Timestamp    time.Time
	InvoiceTotal string
	VatTotal     string
	XmlHash      string
	XmlHashSig   string
	PublicKey    string
	CaStampSig   string
}

func DecodeQrCode(b64TlvQr string) (*ZatcaQr, error) {

	decoder, err := tlv.CreateDecoder(1, 1, binary.BigEndian)
	if err != nil {
		return nil, err
	}

	tlvQrBytes, err := base64.StdEncoding.DecodeString(b64TlvQr)
	if err != nil {
		return nil, err
	}

	tlvNodes, err := decoder.DecodeBytes(tlvQrBytes)
	if err != nil {
		return nil, err
	}

	zatcaQr := &ZatcaQr{}

	for _, node := range tlvNodes {
		switch node.Tag {
		case 1:
			zatcaQr.SellerName = string(node.Value)
		case 2:
			zatcaQr.VatRegNo = string(node.Value)
		case 3:
			iso8601Time := string(node.Value)
			parsedTime, err := time.Parse("2006-01-02T15:04:05Z", iso8601Time)
			if err != nil {
				return nil, err
			}
			zatcaQr.Timestamp = parsedTime
		case 4:
			zatcaQr.InvoiceTotal = string(node.Value)
		case 5:
			zatcaQr.VatTotal = string(node.Value)
		case 6:
			zatcaQr.XmlHash = string(node.Value)
		case 7:
			zatcaQr.XmlHashSig = string(node.Value)
		case 8:
			zatcaQr.PublicKey = base64.StdEncoding.EncodeToString(node.Value)
		case 9:
			zatcaQr.CaStampSig = string(node.Value)
		}
	}

	return zatcaQr, nil
}
