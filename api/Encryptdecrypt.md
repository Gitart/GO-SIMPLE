# Encrypt and decrypt data with AES crypto

In this tutorial we will learn how to encrypt data with Golang's AES crypto package.  
AES or Advanced Encryption Standard is encryption algorithm based on the Rijndael cipher   
developed by the Belgian cryptographers, Joan Daemen and Vincent Rijmen. AES was adopted   
for encryption by the United States government and is now used worldwide.   

Below is the source code for encrypting and decrypting a 16 bytes string message :


```golang
 package main

 import (
   "fmt"
   "crypto/aes"
   "crypto/cipher"
 )

 func main() {
    //The key argument should be the AES key, either 16, 24, or 32 bytes to select AES-128, AES-192, or AES-256.
    key := "opensesame123456" // 16 bytes!

    block,err := aes.NewCipher([]byte(key))

    if err != nil {
      panic(err)
    }

    fmt.Printf("%d bytes NewCipher key with block size of %d bytes\n", len(key), block.BlockSize)

    str := []byte("Hello World!")

    // 16 bytes for AES-128, 24 bytes for AES-192, 32 bytes for AES-256
    ciphertext := []byte("abcdef1234567890") 
    iv := ciphertext[:aes.BlockSize] // const BlockSize = 16


    // encrypt

    encrypter := cipher.NewCFBEncrypter(block, iv)

    encrypted := make([]byte, len(str))
    encrypter.XORKeyStream(encrypted, str)

    fmt.Printf("%s encrypted to %v\n", str, encrypted)

    // decrypt

    decrypter := cipher.NewCFBDecrypter(block, iv) // simple!

    decrypted := make([]byte, len(str))
    decrypter.XORKeyStream(decrypted, encrypted)

    fmt.Printf("%v decrypt to %s\n", encrypted, decrypted)


 }
 ```
Executing the above code will produce the following output :

```
16 bytes NewCipher key with block size of 9360 bytes
Hello World! encrypted to [246 121 236 39 139 97 102 181 16 102 237 145]
[246 121 236 39 139 97 102 181 16 102 237 145] decrypt to Hello World!
Hope this tutorial is useful for those learning Go and AES crypto.
```
