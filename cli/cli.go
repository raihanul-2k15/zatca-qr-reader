package cli

import (
	"fmt"
	"os"

	"github.com/raihanul-2k15/zatca-qr-reader/extract"
	"github.com/raihanul-2k15/zatca-qr-reader/zatcaqr"
)

func Run() {

	if len(os.Args) < 2 {
		fmt.Println("Please provide a file path or a base64 encoded tlv")
		os.Exit(1)
	}

	filePath := os.Args[1]

	// first assume arg itself is the base64 encoded tlv
	qr, err := extract.FromText(filePath).ExtractBase64TlvQr()

	if err != nil {
		// if not, assume arg is a file path
		qr, err = extractQrFromFile(filePath)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}

	details, err := zatcaqr.DecodeQrCode(qr)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Printf("QR Code decoded successfully\n\nDetails:\n")

	fmt.Printf("SellerName:   %s\n", details.SellerName)
	fmt.Printf("VatRegNo:     %s\n", details.VatRegNo)
	fmt.Printf("Timestamp:    %s\n", details.Timestamp)
	fmt.Printf("InvoiceTotal: %s\n", details.InvoiceTotal)
	fmt.Printf("VatTotal:     %s\n", details.VatTotal)
	fmt.Printf("XmlHash:      %s\n", details.XmlHash)
	fmt.Printf("XmlHashSig:   %s\n", details.XmlHashSig)
	fmt.Printf("PublicKey:    %s\n", details.PublicKey)
	fmt.Printf("CaStampSig:   %s\n", details.CaStampSig)
}

func extractQrFromFile(filePath string) (string, error) {

	exists, _ := fileExists(filePath)
	if !exists {
		fmt.Printf("File %s does not exist", filePath)
		os.Exit(1)
	}

	return extract.FromFile(filePath).ExtractBase64TlvQr()
}

func fileExists(filename string) (bool, error) {
	_, err := os.Stat(filename)
	if err == nil {
		// File exists
		return true, nil
	} else if os.IsNotExist(err) {
		// File does not exist
		return false, nil
	} else {
		// An error occurred while checking the file
		return false, err
	}
}
