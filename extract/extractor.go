package extract

type Extractor interface {
	ExtractBase64TlvQr() (string, error)
}
