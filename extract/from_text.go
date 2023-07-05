package extract

import (
	"encoding/base64"
	"errors"
	"strings"
)

type extractorFromText struct {
	text string
}

func FromText(text string) *extractorFromText {
	return &extractorFromText{text: text}
}

func (e *extractorFromText) ExtractBase64TlvQr() (string, error) {
	e.text = strings.TrimSpace(e.text)

	_, err := base64.StdEncoding.DecodeString(e.text)
	if err != nil {
		return "", errors.New("Not a valid base64 encoded string")
	}

	return e.text, nil
}
