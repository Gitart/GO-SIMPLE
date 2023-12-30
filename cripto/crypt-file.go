package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"io"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	os.Remove("assets/ciphertext.bin")
	os.Remove("assets/outputs/plaintext.txt")

	encryptFile()
	decryptFile()
}

func encryptFile() {
	// Reading plaintext file
	plainText, err := ioutil.ReadFile("assets/plaintext.txt")
	if err != nil {
		log.Fatalf("read file err: %v", err.Error())
	}

	// Reading key
	key, err := ioutil.ReadFile("assets/key.txt")
	if err != nil {
		log.Fatalf("read file err: %v", err.Error())
	}

	// Creating block of algorithm
	block, err := aes.NewCipher(key)
	if err != nil {
		log.Fatalf("cipher err: %v", err.Error())
	}

	// Creating GCM mode
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		log.Fatalf("cipher GCM err: %v", err.Error())
	}
	
	// Generating random nonce
	nonce := make([]byte, gcm.NonceSize())
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		log.Fatalf("nonce  err: %v", err.Error())
	}

	// Decrypt file
	cipherText := gcm.Seal(nonce, nonce, plainText, nil)
	
	// Writing ciphertext file
	err = ioutil.WriteFile("assets/ciphertext.bin", cipherText, 0777)
	if err != nil {
		log.Fatalf("write file err: %v", err.Error())
	}

}

func decryptFile() {
	// Reading ciphertext file
	cipherText, err := ioutil.ReadFile("assets/ciphertext.bin")
	if err != nil {
		log.Fatal(err)
	}

	// Reading key
	key, err := ioutil.ReadFile("assets/key.txt")
	if err != nil {
		log.Fatalf("read file err: %v", err.Error())
	}

	// Creating block of algorithm
	block, err := aes.NewCipher(key)
	if err != nil {
		log.Fatalf("cipher err: %v", err.Error())
	}

	// Creating GCM mode
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		log.Fatalf("cipher GCM err: %v", err.Error())
	}
	
	// Deattached nonce and decrypt
	nonce := cipherText[:gcm.NonceSize()]
	cipherText = cipherText[gcm.NonceSize():]
	plainText, err := gcm.Open(nil, nonce, cipherText, nil)
	if err != nil {
		log.Fatalf("decrypt file err: %v", err.Error())
	}
	
	// Writing decryption content
	err = ioutil.WriteFile("assets/outputs/plaintext.txt", plainText, 0777)
	if err != nil {
		log.Fatalf("write file err: %v", err.Error())
	}
}