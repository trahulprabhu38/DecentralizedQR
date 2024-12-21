package main

import (
	"fmt"

	"github.com/skip2/go-qrcode"
)

func main() {
	// Content to encode in the QR code
	content := "https://www.linkedin.com/in/trahulprabhu38/"

	// Generate QR code and save as PNG file
	err := qrcode.WriteFile(content, qrcode.Medium, 256, "linkedin.png")
	if err != nil {
		fmt.Println("Error generating QR code:", err)
		return
	}

	fmt.Println("QR code generated and saved as qrcode.png")
}
