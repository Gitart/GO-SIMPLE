# Hash\-based Message Authentication Code (HMAC)

Hash\-based MACs (HMACs) takes a long message as the input and produce a fixed\-length output. In this scheme, the sender signs a message using the MAC and the receiver verifies it using the shared key. It hashes the key with the message using either of the two methods known as a secret prefix (key comes first and the message comes afterwards) or the secret suffix (key comes after the message).

---

## High\-level design of HMAC

Message Authentication Code (MAC) is a small part of information or a small algorithm, basically used to authenticate a message and to maintain integrity and authenticity assurances on the message. Hash\-based Message Authentication Code is a message authentication code derived from a cryptographic hash function such as MD5 and SHA\-1. The basic idea behind HMAC is to add a layer using a secret key in the existing message digest algorithms. Even if an attacker got the database of hashed passwords with the salts, they would still have a difficult time cracking them without the secret key. As algorithms such as MD5 and SHA\-1 do not rely on the secret key, HMAC has been selected as mandatory\-to\-implement MAC for IP security. HMAC can work with any existing message digest algorithms (hash functions). It considers the message digest produced by the embedded hash function as a black box. It then uses the shared symmetric key to encrypt the message digest, thus, producing the final output, that is, MAC. HMAC is a calculation of a MAC through the use of a cryptographic hash function such as MD5 or SHA\-1.

When we use SHA\-1, then corresponding MAC would be known as HMAC\-SHA1, or if SHA\-2 is being used then we would say HMAC\-SHA256. It is a good practice to store secret key in a separate location such as an environment variable rather than in the database with hashed passwords and salts.

---

## HMAC in Go

This example has a limited use as it is. Use it as a reference for your own applications.

package main

import (
	"crypto/hmac"
	"crypto/rand"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/base64"
	"fmt"
	"io"
)

var secretKey = "4234kxzjcjj3nxnxbcvsjfj"

// Generate a salt string with 16 bytes of crypto/rand data.
func generateSalt() string {
	randomBytes := make(\[\]byte, 16)
	\_, err := rand.Read(randomBytes)
	if err != nil {
		return ""
	}
	return base64.URLEncoding.EncodeToString(randomBytes)
}

func main() {
	message := "Today web engineering has modern apps adhere to what is known as a single\-page app (SPA) model."
	salt := generateSalt()
	fmt.Println("Message: " + message)
	fmt.Println("\\nSalt: " + salt)

	hash := hmac.New(sha256.New, \[\]byte(secretKey))
	io.WriteString(hash, message+salt)
	fmt.Printf("\\nHMAC\-Sha256: %x", hash.Sum(nil))

	hash = hmac.New(sha512.New, \[\]byte(secretKey))
	io.WriteString(hash, message+salt)
	fmt.Printf("\\n\\nHMAC\-sha512: %x", hash.Sum(nil))
}

#### Output of above program

C:\\golang\\example\>go run hashing1.go
Message: Today web engineering has modern apps adhere to what is known as a sing
le\-page app (SPA) model.

Salt: iWk9q\-tQgWQTnqDgdoxaXQ==

HMAC\-Sha256: b158c5a1bbcdac3cf87fe761030828cb5811b0a6fdfa6366c7bdfddba6391728

HMAC\-sha512: e350ca7f0349c2b16a410f224b1ad0c8fc9319708b1dd2be9e83a53b3d4b93d9dd1
f0637ea27641edcfac3d3196795d9889778bd4894ad332ba643d0735aa089
C:\\golang\\example\>
