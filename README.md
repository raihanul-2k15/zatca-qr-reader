# ZATCA QR Reader

This is a Go based ZATCA E-Invoicing Phase 2 QR Code extraction and decoding tool that can extract the base64 encoded TLV (Tag Length Value) QR code and decode it to display information contained in it.

ZATCA ([Zakat, Tax and Customs Authority](https://zatca.gov.sa/en/Pages/default.aspx)) is a government organization of Saudi Arabia and has mandated the issuance of electronic invoices with `UBL2.1` XML and `base64` encoded `TLV` QR Codes.

This tool can be used by developers to quickly verify their own implementation of Phase 2 QR Code generation or simply extract and view QR Code from any ZATCA E-Invoice.

This tool (with the desktop app) can also be used by non technical people to extract and view contents of QR code from a variety of inputs.

## Features

### CLI Program

| Feature                            | Implemented        |
| ---------------------------------- | ------------------ |
| Decode from base64 string          | :heavy_check_mark: |
| Decode from a PNG/JPG image file   | :heavy_check_mark: |
| Extract and decode from a PDF file | :heavy_check_mark: |
| Extract and decode from a XML file | :heavy_check_mark: |

### Desktop App

Coming Soon :smiley:
