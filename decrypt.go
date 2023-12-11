package main

import (
	"fmt"
	"github.com/ProtonMail/gopenpgp/v2/helper"
	"os"
)

func main() {
	// Replace with the path to the private key file
	privateKeyPath := "path/to/private-key.asc"

	// Replace with the passphrase for the private key
	passphrase := []byte("test")

	// Replace with the path to the encrypted file
	encryptedFilePath := "path/to/plain-text-file.txt.gpg"

	// Replace with the path where you want to save the decrypted file
	decryptedFilePath := "path/to/decrypted/plain-text-file.txt"

	// Read private key
	privateKey, err := os.ReadFile(privateKeyPath)
	if err != nil {
		fmt.Println("Error reading private key:", err)
		return
	}

	// Read encrypted content
	encryptedContent, err := os.ReadFile(encryptedFilePath)
	if err != nil {
		fmt.Println("Error reading encrypted file:", err)
		return
	}

	// Decrypt the encrypted content with the private key and passphrase
	decryptedContent, err := helper.DecryptMessageArmored(string(privateKey), passphrase, string(encryptedContent))
	if err != nil {
		fmt.Println("Error decrypting file:", err)
		return
	}

	// Save the decrypted content to the output file
	err = os.WriteFile(decryptedFilePath, []byte(decryptedContent), 0644)
	if err != nil {
		fmt.Println("Error writing decrypted content to file:", err)
		return
	}

	fmt.Println("File decrypted successfully.")
}
