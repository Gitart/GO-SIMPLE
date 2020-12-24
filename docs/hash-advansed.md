# Advanced Encryption Standard

The national institute of standards and technology (NIST) announced a call for a project for creating a new cipher in January 1997. Many groups had proposed various ciphers. Various ciphers were examined on speed and security parameters and after several rounds of studies and examinations, NIST has finally chosen an algorithm known as Rijndael. Rijndael selected as the best algorithm in terms of security, cost, resilience, integrity and surveillance of the algorithm, hence NIST selected Rijndael as advanced encryption standard (AES) in October 2000.

On 26 November 2001, AES became a FIPS (Federal Information Processing Standards) standard. AES specifies a FIPS\-approved cryptographic algorithm used to secure electronic data. The U.S. government (NSA) in June 2003 accepted and announced that AES was secure enough to safeguard highly classified information up to the supersecret level.

Rijndael was named because it was developed by two Belgian cryptographers Dr Joan Daemen and Dr Vincent Rijmen at the Electrical Engineering Department of Katholieke University in Leuven. Rijndael or AES is patent free, and the creators have given out various reference implementations as public domain.

---

---

## AES Encryption and Decryption in Go

Below sample program will encrypt a text message and decrypt a file using a key, which is basically a 16\-byte (128\-bit) password. This program will create two files **aes.enc** which contain encrypted data and **aes.key** which contains AES key.

This example has a limited use as it is. Use it as a reference for your own applications.

package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/pem"
	"fmt"
	"io/ioutil"
	"log"
)

const (
	keyFile       = "aes.key"
	encryptedFile = "aes.enc"
)

var IV = \[\]byte("1234567812345678")

func readKey(filename string) (\[\]byte, error) {
	key, err := ioutil.ReadFile(filename)
	if err != nil {
		return key, err
	}
	block, \_ := pem.Decode(key)
	return block.Bytes, nil
}

func createKey() \[\]byte {
	genkey := make(\[\]byte, 16)
	\_, err := rand.Read(genkey)
	if err != nil {
		log.Fatalf("Failed to read new random key: %s", err)
	}
	return genkey
}

func saveKey(filename string, key \[\]byte) {
	block := &pem.Block{
		Type:  "AES KEY",
		Bytes: key,
	}
	err := ioutil.WriteFile(filename, pem.EncodeToMemory(block), 0644)
	if err != nil {
		log.Fatalf("Failed in saving key to %s: %s", filename, err)
	}
}

func aesKey() \[\]byte {
	file := fmt.Sprintf(keyFile)
	key, err := readKey(file)
	if err != nil {
		log.Println("Creating a new AES key")
		key = createKey()
		saveKey(file, key)
	}
	return key
}

func createCipher() cipher.Block {
	c, err := aes.NewCipher(aesKey())
	if err != nil {
		log.Fatalf("Failed to create the AES cipher: %s", err)
	}
	return c
}

func encryption(plainText string) {
	bytes := \[\]byte(plainText)
	blockCipher := createCipher()
	stream := cipher.NewCTR(blockCipher, IV)
	stream.XORKeyStream(bytes, bytes)
	err := ioutil.WriteFile(fmt.Sprintf(encryptedFile), bytes, 0644)
	if err != nil {
		log.Fatalf("Writing encryption file: %s", err)
	} else {
		fmt.Printf("Message encrypted in file: %s\\n\\n", encryptedFile)
	}
}

func decryption() \[\]byte {
	bytes, err := ioutil.ReadFile(fmt.Sprintf(encryptedFile))
	if err != nil {
		log.Fatalf("Reading encrypted file: %s", err)
	}
	blockCipher := createCipher()
	stream := cipher.NewCTR(blockCipher, IV)
	stream.XORKeyStream(bytes, bytes)
	return bytes
}

func main() {

	var plainText = "AES is now being used worldwide for encrypting digital information, including financial, and government data."
	encryption(plainText)

	fmt.Printf("Decrypted Message: %s", decryption())
}

#### Output of above

C:\\golang\\test\>go fmt aes.go

C:\\golang\\test\>go run aes.go
Message encrypted in file: aes.enc

Decrypted Message: AES is now being used worldwide for encrypting digital inform
ation, including financial, and government data.
C:\\golang\\test\>

---

## High\-level design of AES

AES have arithmetic operations are based on Galois Filed which have GF(2N) structure where N = 8. AES is a symmetric cipher which uses the same key for both encryption and decryption process. This symmetric cipher encrypts a 128\-bit block of plaintext using a 128\-bit key value to produce a 128\-bit ciphertext at a time. AES needs a large 128\-bit key size to implement encryption and decryption process.

AES 128\-bit cipher uses 10 rounds(a substitution and permutation network design with a single collection of steps) of operation for performing encryption and decryption processes. Depending on the types of keys and number of rounds operations, the three versions are AES\-128 uses 10 rounds, AES\-192 uses 12 rounds and AES\-256 uses 14 rounds of operations are in used available.

AES entire data block is being processed in an identical way during each round. In AES, a plaintext has to travel through *Nr* number of rounds before producing the cipher. Again, each round comprises four different operations. One operation is permutation and the other three are substitutions. They are SubBytes, ShiftRows, MixColumns, and AddRoundKey.

In AES, all the transformations that are being used in the encryption process will have the inverse transformations that are being used in the decryption process. Each round of the decryption process in AES uses the inverse transformations InvSubBytes(), InvShiftRows() and InvMixColumns().

---

## Strong encryption with AES

As AES was produced after DES, all identified attacks on DES have been demonstrated on AES, and all the final results were valid. AES is more confident to brute\-force attack than DES because of its larger variable key size and block size. AES is not susceptible to statistical attacks, and it has been tested that it is not achievable with common techniques to do the statistical analysis of ciphertext in AES. As of today, there is no differential, linear and successful attack on AES has been detected. The best part of AES is that the algorithms used in it are so basic that they can be quickly implemented using cheap processors and a minimum amount of memory. AES needs higher processing and more rounds of transmission than DES, and we can comparatively tell this is AES's disadvantage.
