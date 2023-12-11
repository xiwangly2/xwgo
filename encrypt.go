package main

import (
	"fmt"
	"github.com/ProtonMail/gopenpgp/v2/helper"
	"os"
)

func main() {
	// Replace with the path to the public key file
	publicKeyPath := "path/to/public-key.asc"

	// Replace with the path to the file you want to encrypt.exe
	inputFilePath := "path/to/plain-text-file.txt"

	// Replace with the path where you want to save the encrypted file
	outputFilePath := "path/to/plain-text-file.txt.gpg"

	// Read public key
	publicKey, err := os.ReadFile(publicKeyPath)
	if err != nil {
		fmt.Println("Error reading public key:", err)
		return
	}

	// Read file content
	fileContent, err := os.ReadFile(inputFilePath)
	if err != nil {
		fmt.Println("Error reading input file:", err)
		return
	}

	// Encrypt the file content with the public key
	encryptedContent, err := helper.EncryptMessageArmored(string(publicKey), string(fileContent))
	if err != nil {
		fmt.Println("Error encrypting file:", err)
		return
	}

	// Save the encrypted content to the output file
	err = os.WriteFile(outputFilePath, []byte(encryptedContent), 0644)
	if err != nil {
		fmt.Println("Error writing encrypted content to file:", err)
		return
	}

	fmt.Println("File encrypted successfully.")
}
