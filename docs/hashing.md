# What is Hashing?

The process of taking plaintext and transforms it into a digest of the plaintext information, in such a way that it is not intended to be decrypt is called Hashing. The output of Hashing is known as a hash, hash value, or message digest. Hashing is a intriguing area of cryptography and is different from encryption algorithms. Hashing creates a scrambled output that cannot be reversed easily. Technically, a hashing generates a fixed\-length value that is relatively easy to compute in one direction, but nearly impossible to reverse.

---

## Hashing Basics

A hash, hash value, or message digest is a value which is an output of plaintext or ciphertext being given into a hashing algorithm. No matter what is input into the hashing algorithm, the hash is of a fixed length and will always be of a certain length. The resulting hash has its length fixed by the design of the algorithm itself. We also refer, a hash as a summary of a file or message, often in numeric format. Hashes are being used in digital signatures, in the file and message authentication, and to protect the integrity of sensitive data.

A hash can take place into the category of a one\-way function. This shows hash can compute when generated but difficult (or impossible) to compute in reverse. With a hash, initial computation is relatively easy to produce a condensed version of the message,  but the original message should not be re\-created from the hash.

---

## Cryptographic Hash Functions

The simplest approach to hash a message is to slice it into chunks and process each chunk successively using a similar algorithm. This approach is called iterative hashing. Iterative hashing use a compression function that converts an input to a smaller output; and converts an input to an output of the same size, such that any two different inputs give two different outputs. Cryptographic hash functions are those hash functions which have a base on block ciphers. Examples of cryptographic hash functions include message digest (MD) algorithm based such as MD2, MD4 and MD5 and secure hash algorithm based (SHA) such as SHA\-1, SHA\-224, SHA\-256, SHA\-384 and SHA\-512.

---

### Message\-Digest algorithm 5 (MD5)

This cryptographic hash algorithm developed by Ron Rivest in 1991. This algorithm takes a variable\-length message as input and generates a fixed\-length message digest of 128 bits. This algorithm uses Big\-endian scheme where the least significant byte of a 32\-bit word will be stored in the low\-address byte position. This algorithm undergoes four rounds, each having 16 iterations, they use thus total 64 iterations. Hence it requires a 128\-bit buffer. This is less secure but faster in operation as compared to SHA\-1. This algorithm takes 2128 operations for detecting the original message from the given message digest and 264 operations to detect two messages generating the same message digest.

### Secure Hash Algorithm (SHA\-1)

This algorithm takes a variable\-length message as input and generates a fixed\-length message digest of 160 bits. This algorithm uses Little\-endian scheme to interpret the message as a sequence of 32\-bit words. In this algorithm, the most significant byte of a 32\-bit word is stored in the low\-address byte position. This algorithm undergoes four rounds, each having 20 iterations, they use thus total 80 iterations. Hence it requires a 160\-bit buffer. This is more secure but slower in operation as compared to MD5. This algorithm takes 2160 operations for detecting the original message from the given message digest and 280 operations to detect two messages generating the same message digest.

#### Short example to illustrate use of Cryptographic Hash Functions

package main

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"fmt"
)

func main() {
	fmt.Println("\\n\-\-\-\-\-\-\-\-\-\-\-\-\-\-\-\-Small Message\-\-\-\-\-\-\-\-\-\-\-\-\-\-\-\-\\n")
	message := \[\]byte("Today web engineering has modern apps adhere to what is known as a single\-page app (SPA) model.")

	fmt.Printf("Md5: %x\\n\\n", md5.Sum(message))
	fmt.Printf("Sha1: %x\\n\\n", sha1.Sum(message))
	fmt.Printf("Sha256: %x\\n\\n", sha256.Sum256(message))
	fmt.Printf("Sha512: %x\\n\\n", sha512.Sum512(message))

	fmt.Println("\\n\\n\-\-\-\-\-\-\-\-\-\-\-\-\-\-\-\-Large Message\-\-\-\-\-\-\-\-\-\-\-\-\-\-\-\-\\n")
	message = \[\]byte("Today web engineering has modern apps adhere to what is known as a single\-page app (SPA) model. This model gives you an experience in which you never navigate to particular pages or even reload a page.  It loads and unloads the various views of our app into the same page itself. If you've ever run popular web apps like Gmail, Facebook, Instagram, or Twitter, you've used a single\-page app. In all those apps, the content gets dynamically displayed without requiring you to refresh or navigate to a different page. React gives you a powerful subjective model to work with and supports you to build user interfaces in a declarative and component\-driven way.")

	fmt.Printf("Md5: %x\\n\\n", md5.Sum(message))
	fmt.Printf("Sha1: %x\\n\\n", sha1.Sum(message))
	fmt.Printf("Sha256: %x\\n\\n", sha256.Sum256(message))
	fmt.Printf("Sha512: %x\\n\\n", sha512.Sum512(message))
}

#### Output of above program

C:\\golang\\example\>go run hashing.go

\-\-\-\-\-\-\-\-\-\-\-\-\-\-\-\-Small Message\-\-\-\-\-\-\-\-\-\-\-\-\-\-\-\-

Md5: d53c7002ebeeaa872f02efdda82f76f0

Sha1: b56fd6c3eabc1ea6e880ffb7762d31a4b39bbcc9

Sha256: 0e81731d26a2ffd898be00c40bf0ab3c24ec82a0837311701193dd861efdb944

Sha512: 97be88309470c98e8ff840a74567c3948263251f3ac6f232a6e228cebc6daa6d95569a96
6a5f2ba5d694dc644d02c144849e7a181a159b583a2060e2ce3ad375

\-\-\-\-\-\-\-\-\-\-\-\-\-\-\-\-Large Message\-\-\-\-\-\-\-\-\-\-\-\-\-\-\-\-

Md5: f52c27fb5fa09afbb717e5ded9ef5707

Sha1: 86b1102f10518ef02089742862695e132ca6b5fe

Sha256: 88c9eca72b7b635458945710d12f6709af44dd02c0a3bb9b16c7e352b6d13f84

Sha512: 5400f78f1642d5b7df7a0c60e22e95ba04a94d16d33d008fc78a5002958f853ecc469bb8
786632c0be47a0c5b44d5f78aa9878abe7f00b688d8712c921aa81b8

C:\\golang\\example\>

Due to vulnerabilities in MD5 and SHA\-1, some agencies has started using SHA\-2 as early as 2011. SHA\-2 is more secure, and it has a 256\-bit and 512\-bit block sizes. A common hash used on the Internet is SHA\-2, SHA\-1 is now deprecated and should be replaced if it's in use.

---

We commonly using this hash functions in public\-key encryption, message authentication, key agreement protocols, digital signatures, integrity verification, identification protection and many other cryptographic contacts. Always there is a hash function somewhere under the hood whether we're encrypting an email, sending a message on your mobile phone, connecting to a HTTPS website, or connecting to a remote machine through IPSec or SSH.

#### List of real world examples where we are using Hash functions:

*   Cloud storage systems use hash functions to analyze identical files and to find changed files.
*   Git revision control system uses hash functions to detect files in a repository.
*   Bitcoin using a hash function in its proof\-of\-work systems.
*   Forensic analysts use hash values to verify that digital artifacts hasn't altered.
*   NIDS use hashes to analyze known\-malicious data going through a network.
