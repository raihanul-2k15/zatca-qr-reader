package extract

import (
	"bytes"
	"errors"
	"io/ioutil"

	"github.com/pdfcpu/pdfcpu/pkg/api"
	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu/model"
)

type extractorFromPdf struct {
	pdfBytes []byte
}

func FromPdf(pdfBytes []byte) *extractorFromPdf {
	return &extractorFromPdf{pdfBytes: pdfBytes}
}

func (e *extractorFromPdf) ExtractBase64TlvQr() (string, error) {
	pageImagesMap, err := api.ExtractImagesRaw(bytes.NewReader(e.pdfBytes), []string{}, nil)
	if err != nil {
		return "", err
	}

	images := make([]model.Image, 0)
	for _, pageImages := range pageImagesMap {
		for _, img := range pageImages {
			images = append(images, img)
		}
	}

	for i := range images {
		imgBytes, err := ioutil.ReadAll(images[i])
		if err != nil {
			continue
		}

		possibleQr, err := FromImage(imgBytes).ExtractBase64TlvQr()
		if err != nil {
			return "", err
		}

		return possibleQr, nil
	}

	return "", errors.New("No decodable QR code images found in PDF")
}
