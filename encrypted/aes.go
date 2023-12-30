
Skip to content


Create account

4
Jump to Comments
17
Save


Bouchaala Reda
Posted on 20 янв. • Updated on 21 янв.


9
Secret Key Encryption with Go using AES
#
go
#
security
#
cryptography
#
encryption
I was working on a personal Go project when I needed the ability to encrypt & decrypt arbitrary pieces of data (strings or JSON payloads or whatever) using a secret key. This article is the result of all the learning I had in the last couple of days while trying to achieve that.

TL;DR
Secret key (or Symmetric) encryption requires the use of a block cipher such as AES. AES by itself can only encrypt/decrypt 16 byte long data (which is its block size), hence the need to use a block cipher mode.

GCM block cipher mode is one of many modes. It is a method that uses AES in a way that enables us to encrypt/decrypt arbitrary sized data. It also adds message authentication (integrity).

The complete working code example is here.

A Cryptography primer
Before diving in, I'd like to define some Cryptography terminology.

A secret key is basically an agreed upon value between communicating parties (machines, you and your friends ... or any combination of that).
Plaintext is the text that you want to encrypt so that no one else can read it unless that someone has the secret key.
Ciphertext is the encryption result of the plaintext.
A cipher is the algorithm that does the actual work of encrypting and decrypting our data.
What we are talking about here is called Symmetric encryption (also called Secret key encryption) whereby we use the same key to encrypt our plaintext to ciphertext and to revert the ciphertext back to plaintext.

Cipher types
Now let's talk a bit about ciphers. Symmetric ciphers belong to two main categories:

Stream ciphers: operate by encrypting each bit (or byte) of the plaintext at a time, producing the ciphertext.
and Block ciphers: operate by encrypting fixed-sized blocks of plaintext.
Stream and Block ciphers are the building blocks (no pun intended) of more complicated and siphisticated cryptographic utilities such as MACs, hash functions, symmetric-key digital signature schemes and much more.

Now let's take one example of a Block cipher and dive more into it: AES, Advanded Encryption Standard.

AES is a block cipher that takes a fixed-size key and fixed-size plaintext, and returns fixed-size ciphertext. AES has three variants that are selected based on the secret key length, all of which use a fixed-sized block of 16 bytes (or 128 bits).

Secret Key Length	AES Variant	Block Size
16 bytes (128 bits)	AES-128	16 bytes (128 bits)
24 bytes (192 bits)	AES-192	16 bytes (128 bits)
32 bytes (25 6bits)	AES-256	16 bytes (128 bits)
Because the block size of AES is set to 16 bytes, the plaintext must be at least 16 bytes long. Which causes a problem for us since we want to be able to encrypt/decrypt arbitrary sized data.

Let's put that aside for now and let's have a look at an example of encrypting some data using AES.

Let's give it a try.
package main

import (
    "crypto/aes"
    "fmt"
)

var (
    // We're using a 32 byte long secret key
    secretKey string = "N1PCdw3M2B1TfJhoaY2mL736p2vCUc47"
)

func encrypt(plaintext string) (string) {
    aes, err := aes.NewCipher([]byte(secretKey))
    if err != nil {
        panic(err)
    }

    // Make a buffer the same length as plaintext
    ciphertext := make([]byte, len(plaintext))
    aes.Encrypt(ciphertext, []byte(plaintext))

    return string(ciphertext)
}

func main() {
    // This will successfully encrypt.
    ciphertext := encrypt("This is some sensitive information")
    fmt.Printf("Ciphertext: %x \n", ciphertext)

    // This will cause an error since the
    // plaintext is less than 16 bytes.
    ciphertext = encrypt("Hello")
}
Block Cipher Modes
Now back to our problem: using AES alone will not be enough to get what we want which is to: Encrypt arbitrary sized messages.

This is where Block Cipher modes (or Block modes of Operation) come in. A Block Cipher mode is a method that uses the Block cipher to solve a particular problem. There's a lot of Block Cipher modes available and the majority of them solve the problem of "How do I encrypt an arbitrary sized message". That's exactly what we want!

I'm by no means a Cryptography expert so I can't and I won't compare block cipher modes or dive into their details. But here's what we need to undertstand:

Block ciphers by themselves have two limitations:

They can only encrypt/decrypt fixed-sized data
They only provide confidentiality which means one looking at ciphertext cannot possibly revert it back to plaintext without knowing the secret key
Block cipher modes were created to solve those two limtations:

All block cipher modes solve the size limitation.
Some modes only solve the size limitation and nothing else, like ECB, CBC & CTR.
Other modes were created to also add authentication or message integrity (these are called combined modes) like: CCM, GCM & TKW.
By using one of the modes that also provide authentication, we are effectively using Authenticated Encryption, as opposed to Block encryption. Here's a recap to get the full picture:

AES is a Block Cipher that gives us confidentiality. This is *Block Encryption *.
When using AES alone, you can only encrypt/decrypt data that is 16 bytes long (which the size of an AES block).
Using AES with CBC mode for example, aleviates the size limitation. This is also Block Encryption.
Using AES with GCM (a combined mode) aleviates the size limitation but also gives us message authentication (integrity). This is Authenticated Encryption.
Alright, we'll be using GCM mode bacause it's one of the most widely adopted symmetric block cipher modes. GCM requires an IV (initialization vector) that should ALWAYS be randomly generated (the term used here is nonce, which is pretty much the same). We're just using a random string in our example.

Using AES with GCM
Let's see how we can do that in Go:
package main

import (
    "crypto/aes"
    "crypto/cipher"
    "crypto/rand"
    "fmt"
)

var (
    // We're using a 32 byte long secret key.
    // This is probably something you generate first
    // then put into and environment variable.
    secretKey string = "N1PCdw3M2B1TfJhoaY2mL736p2vCUc47"
)

func encrypt(plaintext string) string {
    aes, err := aes.NewCipher([]byte(secretKey))
    if err != nil {
        panic(err)
    }

    gcm, err := cipher.NewGCM(aes)
    if err != nil {
        panic(err)
    }

    // We need a 12-byte nonce for GCM (modifiable if you use cipher.NewGCMWithNonceSize())
    // A nonce should always be randomly generated for every encryption.
    nonce := make([]byte, gcm.NonceSize())
    _, err = rand.Read(nonce)
    if err != nil {
        panic(err)
    }

    // ciphertext here is actually nonce+ciphertext
    // So that when we decrypt, just knowing the nonce size
    // is enough to separate it from the ciphertext.
    ciphertext := gcm.Seal(nonce, nonce, []byte(plaintext), nil)

    return string(ciphertext)
}

func decrypt(ciphertext string) string {
    aes, err := aes.NewCipher([]byte(secretKey))
    if err != nil {
        panic(err)
    }

    gcm, err := cipher.NewGCM(aes)
    if err != nil {
        panic(err)
    }

    // Since we know the ciphertext is actually nonce+ciphertext
    // And len(nonce) == NonceSize(). We can separate the two.
    nonceSize := gcm.NonceSize()
    nonce, ciphertext := ciphertext[:nonceSize], ciphertext[nonceSize:]

    plaintext, err := gcm.Open(nil, []byte(nonce), []byte(ciphertext), nil)
    if err != nil {
        panic(err)
    }

    return string(plaintext)
}

func main() {
    // This will successfully encrypt & decrypt
    ciphertext1 := encrypt("This is some sensitive information")
    fmt.Printf("Encrypted ciphertext 1: %x \n", ciphertext1)

    plaintext1 := decrypt(ciphertext1)
    fmt.Printf("Decrypted plaintext 1: %s \n", plaintext1)

    // This will successfully encrypt & decrypt as well.
    ciphertext2 := encrypt("Hello")
    fmt.Printf("Encrypted ciphertext 2: %x \n", ciphertext2)

    plaintext2 := decrypt(ciphertext2)
    fmt.Printf("Decrypted plaintext 2: %s \n", plaintext2)
}
As you can see, we can now encrypt and decrypt any arbitrary sized data! If you want to store the ciphertext somewhere, you'd have to encode it with encoding/base64 or encoding/hex or something similar.

This is just an example and shouldn't be used as is in anything real.

In a real world project, one should probably use a better library to encrypt/decrypt sensitive information (like libsodium), but by figuring out how to do that with Go's standard library and by using symmetric cryptography primitives such as block ciphers and block modes, you understand more about they they relate to each other.

Again, I am by no means a security expert. This article is just me sharing my learnings so please let me know in the comments if there's anything I missed, anything you didn't understand or any other general feedback!

Thanks for reading and have a lovely day!

Top comments (4)
Subscribe

Add to the discussion
 
 

Rida F'kih
•
21 янв. • Edited

You should be programmatically generating the nonce (not nounce) for each encryption. It’s vital that the nonce is cryptographically random generated, and hard-coding it in your example might get the wrong message across.

It’s always best to use well-known solutions unless you are an actual cryptography expert, libsodium is great for a lot of cryptographic utilities including what you’re trying to do in this post and argon2 is great for salting & hashing.


3
Like
 
 

Bouchaala Reda
•
21 янв. • Edited

Absolutely. For the sake of keeping the code example simple I opted to just hardcode the nonce, but you're right the nonce is critical for the encryption/decryption and should never be hardcoded. I updated the code example to randomly generate it.

Some libraries don't even give you the ability to pass a nonce when encrypting, they're generated internally.

Thanks for taking the time to write your feedback! Appreciated.


3
Like
 
 

birowo
•
22 янв.

just for practice, I have made the aes-gcm encryption code, encrypt using golang and decrypt using javascript
gitlab.com/birowo/aes-gcm


2
Like
 
 

Bouchaala Reda
•
23 янв.

That's a good example, thanks for sharing.


2
Like
Code of Conduct • Report abuse
DEV Community

Trending in Go
The Go community is exploring cloud backends, full-stack development with Docker, and efficient APIs using Go's performance strengths.

 
Building a cloud backend in Go using REST and PostgreSQL
Marcus Kohlberg for Encore ・ Dec 18
#tutorial #go #cloud #webdev
 
Go + TypeScript full stack web app, with nextjs, PostgreSQL and Docker
Francesco Ciulla ・ Dec 26
#go #webdev #beginners #devops
 
How to use GoLang in Flutter Application - Golang FFI
Jhin Lee ・ Dec 20
#flutter #go #ffi #dart
 
Docker setup for Go APIs
arcade ・ Dec 21
#docker #go #api #backend
 
How to host your own CDN for free in less than 10 minutes
Kevin Nielsen ・ Dec 16
#cdn #go #programming #beginners
Read next

6 🔥 Awesome Golang packages (web devs)
Kevin Naidoo - Dec 4


Building an event-driven system in Go using Pub/Sub
Marcus Kohlberg - Nov 30


How to reverse proxy the WebSocket protocol
Lorain - Dec 4


Which is More Secure: Linux or Windows?
Susheel Thapa - Dec 3


Bouchaala Reda
Follow
Backend Engineer with a love for anything DevOpsy
LOCATION
Algeria
JOINED
3 сент. 2021 г.
More from Bouchaala Reda
How To Safely Verify MACs With Go And PHP Examples
#go #php #cryptography #security
Dynamic PostgreSQL credentials using HashiCorp Vault (with PHP Symfony & Go examples)
#go #php #devops #postgres
DEV Community

🌚 Pro Tip
You can set dark mode, and make other customizations, by creating a DEV account.
Sign up now

package main

import (
    "crypto/aes"
    "crypto/cipher"
    "crypto/rand"
    "fmt"
)

var (
    // We're using a 32 byte long secret key.
    // This is probably something you generate first
    // then put into and environment variable.
    secretKey string = "N1PCdw3M2B1TfJhoaY2mL736p2vCUc47"
)

func encrypt(plaintext string) string {
    aes, err := aes.NewCipher([]byte(secretKey))
    if err != nil {
        panic(err)
    }

    gcm, err := cipher.NewGCM(aes)
    if err != nil {
        panic(err)
    }

    // We need a 12-byte nonce for GCM (modifiable if you use cipher.NewGCMWithNonceSize())
    // A nonce should always be randomly generated for every encryption.
    nonce := make([]byte, gcm.NonceSize())
    _, err = rand.Read(nonce)
    if err != nil {
        panic(err)
    }

    // ciphertext here is actually nonce+ciphertext
    // So that when we decrypt, just knowing the nonce size
    // is enough to separate it from the ciphertext.
    ciphertext := gcm.Seal(nonce, nonce, []byte(plaintext), nil)

    return string(ciphertext)
}

func decrypt(ciphertext string) string {
    aes, err := aes.NewCipher([]byte(secretKey))
    if err != nil {
        panic(err)
    }

    gcm, err := cipher.NewGCM(aes)
    if err != nil {
        panic(err)
    }

    // Since we know the ciphertext is actually nonce+ciphertext
    // And len(nonce) == NonceSize(). We can separate the two.
    nonceSize := gcm.NonceSize()
    nonce, ciphertext := ciphertext[:nonceSize], ciphertext[nonceSize:]

    plaintext, err := gcm.Open(nil, []byte(nonce), []byte(ciphertext), nil)
    if err != nil {
        panic(err)
    }

    return string(plaintext)
}

func main() {
    // This will successfully encrypt & decrypt
    ciphertext1 := encrypt("This is some sensitive information")
    fmt.Printf("Encrypted ciphertext 1: %x \n", ciphertext1)

    plaintext1 := decrypt(ciphertext1)
    fmt.Printf("Decrypted plaintext 1: %s \n", plaintext1)

    // This will successfully encrypt & decrypt as well.
    ciphertext2 := encrypt("Hello")
    fmt.Printf("Encrypted ciphertext 2: %x \n", ciphertext2)

    plaintext2 := decrypt(ciphertext2)
    fmt.Printf("Decrypted plaintext 2: %s \n", plaintext2)
}
DEV Community — A constructive and inclusive social network for software developers. With you every step of your journey.

Home
Podcasts
Videos
Tags
FAQ
Forem Shop
Advertise on DEV
About
Contact
Guides
Software comparisons
Code of Conduct
Privacy Policy
Terms of use
Built on Forem — the open source software that powers DEV and other inclusive communities.

Made with love and Ruby on Rails. DEV Community © 2016 - 2023.