package zatcaqr

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestDecodeStandard(t *testing.T) {
	qr := "ARNBY21lIFdpZGdldOKAmXMgTFREAg8zMTExMTExMTExMDExMTMDFDIwMjItMDktMDdUMTI6MjE6MjhaBAQ0LjYwBQMwLjYGLFZkNUtNYXQyWkp0WnFyRHlvZ2hibFNsdGVvdUlTL3IzL09STnc3T1ovMnM9B2BNRVFDSURHQlJIaVBvNnloWElROWRmNnBNRWt1ZmNHbm9xWWFTK084Sm4weGFnQmlBaUJ0b3hwYnJ3ZkVKSGhVR1FIVHF6RDFPUlg1K1ovdHVtTTB3TGZaNGN1WVJnPT0IWDBWMBAGByqGSM49AgEGBSuBBAAKA0IABGGDDKDmhWAITDv7LXqLX2cmr6+qddUkpcLCvWs5rC2O29W/hS4ajAK4Qdnahym6MaijX75Cg3j4aao7ouYXJ9E="
	zatcaQr, err := DecodeQrCode(qr)

	expectedTs, _ := time.Parse("2006-01-02T15:04:05Z", "2022-09-07T12:21:28Z")

	assert.Nil(t, err)
	assert.Equal(t, "Acme Widget’s LTD", zatcaQr.SellerName)
	assert.Equal(t, "311111111101113", zatcaQr.VatRegNo)
	assert.Equal(t, expectedTs, zatcaQr.Timestamp)
	assert.Equal(t, "4.60", zatcaQr.InvoiceTotal)
	assert.Equal(t, "0.6", zatcaQr.VatTotal)
}
