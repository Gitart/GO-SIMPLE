# Explain the term Cryptography in brief

The term "cryptography" is evolved from two Greek words, namely crypto and graphy. As per Greek language, crypto means secret and graphy means writing. The term crypto has become more popular with the introduction of all crypto currencies like  Bitcoin, Ethereum, and Litecoin.

---

## What is Cryptography?

In simple terms, the process of altering messages in a way that their meaning is hidden from an enemy or opponent who might seize them, is known as Cryptography. Cryptography is the science of secret writing that brings numerous techniques to safeguard information that is present in an unreadable format. Only the designated recipients can be converted this unreadable format into the readable format.

In secure electronic transactions, cryptographic techniques are adopted to secure E\-mail messages, credit card details, audio/video broadcasting, storage media and other sensitive information. By using cryptographic systems, the sender can first encrypt a message and then pass on it through the network. The receiver on the other hand can decrypt the message and restore its original content.

---

## Components of Cryptography

**Plaintext:** Plaintext can be text, binary code, or an image that needs to be converted into a format that is unreadable by anybody except those who carry the secret to unlocking it. It refers to the original unencrypted or unadulterated message that the sender wishes to send.

**Ciphertext:** During the process of encryption *plaintext* get converted into a rushed format, the resulting format is called the ciphertext. It relates to the encrypted message, the receiver receives that. However, ciphertext is like the plaintext that has been operated on by the encryption process to reproduce a final output. This final output contains the original message though in a format, that is not retrievable unless official knows the correct means or can crack the code.

**Encryption:** Encryption, receives information and transforms it to an unreadable format that can be reversed. It is the process of encrypting the plaintext so it can provide the ciphertext. Encryption needs an algorithm called a cipher and a secret key. No one can decrypt the vital information on the encrypted message without knowing the secret key. Plaintext gets transformed into ciphertext using the encryption cipher.

**Decryption:** This is the reverse of the encryption process, in which it transforms the ciphertext back into the plaintext using a decryption algorithm and a secret key. In symmetric encryption, the key used to decrypt is the same as the key used to encrypt. On other hand, in asymmetric encryption or public\-key encryption the key used to decrypt differs from the key used to encrypt.

**Ciphers:** The encryption and decryption algorithms are together known as ciphers. Perhaps the trickiest, interesting and most curious part in the encryption process is the algorithm or cipher. The algorithm or cipher is nothing more than a formula that comprises various steps that illustrate how the encryption/decryption process is being implemented on an information. A basic cipher takes bits and returns bits and it doesn't care whether bits represents textual information, an image, or a video.

**Key:** A key is generally a number or a set of numbers on which the cipher operates. In technical terms, a key is a discrete piece of information that is used to control the output (ciphertext and plaintext) of a given cryptographic algorithm. Encryption and decryption algorithms needs this key to encrypt or decrypt messages, respectively. Sender uses the encryption algorithm and the secret key to convert the plaintext into the ciphertext. On other hand receiver uses same decryption algorithm and the secret key to convert ciphertext back into the plaintext. The longer the secret key is, the harder it is for an attacker to decrypt the message.

![Components of Cryptography](https://www.golangprograms.com/media/wysiwyg/components.jpg)

---

## Example of Cryptography (Classical Cipher)

Below is very basic example, we have created a simple cipher to encrypt and decrypt a plaintext into ciphertext and vice versa. The algorithm cipherAlgorithm() is same for encryption and decryption. The key, we have used is 01, 10 and 15 to encrypt and decrypt the message. The output of encryption is different each time when the key is different. This cipher shifts the letter based on key value, key plays an important role in cryptography.

package main

import (
    "fmt"
    "unicode"
)

// Cipher encrypts and decrypts a string.
type Cipher interface {
    Encryption(string) string
    Decryption(string) string
}

// Cipher holds the key used to encrypts and decrypts messages.
type cipher \[\]int

// cipherAlgorithm encodes a letter based on some function.
func (c cipher) cipherAlgorithm(letters string, shift func(int, int) int) string {
    shiftedText := ""
    for \_, letter := range letters {
        if !unicode.IsLetter(letter) {
            continue
        }
        shiftDist := c\[len(shiftedText)%len(c)\]
        s := shift(int(unicode.ToLower(letter)), shiftDist)
        switch {
        case s < 'a':
            s += 'z' \- 'a' + 1
        case 'z' < s:
            s \-= 'z' \- 'a' + 1
        }
        shiftedText += string(s)
    }
    return shiftedText
}

// Encryption encrypts a message.
func (c \*cipher) Encryption(plainText string) string {
    return c.cipherAlgorithm(plainText, func(a, b int) int { return a + b })
}

// Decryption decrypts a message.
func (c \*cipher) Decryption(cipherText string) string {
    return c.cipherAlgorithm(cipherText, func(a, b int) int { return a \- b })
}

// NewCaesar creates a new Caesar shift cipher.
func NewCaesar(key int) Cipher {
    return NewShift(key)
}

// NewShift creates a new Shift cipher.
func NewShift(shift int) Cipher {
    if shift < \-25 || 25 < shift || shift == 0 {
        return nil
    }
    c := cipher(\[\]int{shift})
    return &c
}

func main() {
    c := NewCaesar(1)
    fmt.Println("Encrypt Key(01) abcd =>", c.Encryption("abcd"))
    fmt.Println("Decrypt Key(01) bcde =>", c.Decryption("bcde"))
    fmt.Println()

    c = NewCaesar(10)
    fmt.Println("Encrypt Key(10) abcd =>", c.Encryption("abcd"))
    fmt.Println("Decrypt Key(10) klmn =>", c.Decryption("klmn"))
    fmt.Println()

    c = NewCaesar(15)
    fmt.Println("Encrypt Key(15) abcd =>", c.Encryption("abcd"))
    fmt.Println("Decrypt Key(15) pqrs =>", c.Decryption("pqrs"))
}

#### Output of above

C:\\golang\\example\>go run test8.go
Encrypt Key(01) abcd =\> bcde
Decrypt Key(01) bcde =\> abcd

Encrypt Key(10) abcd =\> klmn
Decrypt Key(10) klmn =\> abcd

Encrypt Key(15) abcd =\> pqrs
Decrypt Key(15) pqrs =\> abcd

C:\\golang\\example\>

Go provides extensive options for cryptography, such as **encryption**, **hashing**. Go have packages to support symmetric encryption algorithms: **base64**, **AES** and **DES**, this we will cover in coming tutorials.
