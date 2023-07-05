package extract

import (
	"errors"
	"io/ioutil"
	"os"

	"github.com/gabriel-vasile/mimetype"
)

type extractorFromFrom struct {
	filePath string
}

func FromFile(filePath string) *extractorFromFrom {
	return &extractorFromFrom{filePath: filePath}
}

func (e *extractorFromFrom) ExtractBase64TlvQr() (string, error) {
	mimetype.SetLimit(0)
	mtype, err := mimetype.DetectFile(e.filePath)
	if err != nil {
		return "", err
	}

	file, err := os.Open(e.filePath)
	defer file.Close()
	if err != nil {
		return "", err
	}

	fileBytes, err := ioutil.ReadAll(file)
	file.Close()
	if err != nil {
		return "", err
	}

	if mtype.Is("application/pdf") {
		return FromPdf(fileBytes).ExtractBase64TlvQr()
	}

	if mtype.Is("text/xml") {
		return FromXml(fileBytes).ExtractBase64TlvQr()
	}

	if mtype.Is("text/plain") {
		return FromText(string(fileBytes)).ExtractBase64TlvQr()
	}

	if mtype.Is("image/png") || mtype.Is("image/jpeg") {
		return FromImage(fileBytes).ExtractBase64TlvQr()
	}

	return "", errors.New("Cannot extract QR Code from file of mimetype " + mtype.String())
}
