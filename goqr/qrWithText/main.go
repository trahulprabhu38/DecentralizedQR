package main

import (
	"encoding/base64"
	"fmt"
	"log"
	"os"

	"github.com/skip2/go-qrcode"
)

// saveQRCode saves the generated QR code to a file
func saveQRCode(content string, filePath string, size int) {
	err := qrcode.WriteFile(content, qrcode.Medium, size, filePath)
	if err != nil {
		log.Fatalf("Error generating QR code: %v", err)
	}
	fmt.Printf("QR code successfully saved to %s\n", filePath)
}

// encodeSecretKey encodes the secret key in Base64
func encodeSecretKey(secret string) string {
	encoded := base64.StdEncoding.EncodeToString([]byte(secret))
	fmt.Printf("Encoded secret key: %s\n", encoded)
	return encoded
}

// logSuccessMessage prints a success message
func logSuccessMessage(message string) {
	fmt.Println("----------------------------------------------------")
	fmt.Println("SUCCESS:", message)
	fmt.Println("----------------------------------------------------")
}

func main() {
	// Step 1: Define the secret key (or a piece of it)
	secretKeyPiece := "my-secret-key-piece-123"

	// Step 2: Encode the secret key for added security
	fmt.Println("Encoding the secret key...")
	encodedKey := encodeSecretKey(secretKeyPiece)

	// Step 3: Specify the file path and size for the QR code
	filePath := "secret-key-qr.png"
	qrSize := 256

	// Step 4: Generate the QR code and save it to the file
	fmt.Println("Generating QR code...")
	saveQRCode(encodedKey, filePath, qrSize)

	// Step 5: Check if the file exists
	if _, err := os.Stat(filePath); err == nil {
		logSuccessMessage("The QR code has been generated and saved successfully.")
	} else {
		fmt.Println("Error: File was not created successfully.")
	}

	// Step 6: Notify the user
	fmt.Println("You can scan the QR code to retrieve the encoded key.")
	fmt.Println("Ensure the QR code is stored securely to prevent unauthorized access.")
}
