package extract

import (
	"bytes"
	"image"
	_ "image/jpeg"
	_ "image/png"

	"github.com/makiuchi-d/gozxing"
	"github.com/makiuchi-d/gozxing/qrcode"
)

type extractorFromImage struct {
	imageBytes []byte
}

func FromImage(imageBytes []byte) *extractorFromImage {
	return &extractorFromImage{imageBytes: imageBytes}
}

func (e *extractorFromImage) ExtractBase64TlvQr() (string, error) {
	img, _, err := image.Decode(bytes.NewReader(e.imageBytes))
	if err != nil {
		return "", err
	}

	bmp, err := gozxing.NewBinaryBitmapFromImage(img)
	if err != nil {
		return "", err
	}

	qrReader := qrcode.NewQRCodeReader()
	result, err := qrReader.Decode(bmp, nil)
	if err != nil {
		return "", err
	}

	return result.GetText(), nil
}
