## Key Generation KEY PEM


#### Installing certificates to the Intel® Edison board

From your computer, run the following commands:

```
scp -r cert.pem USERNAME@xxx.xxx.x.xxx:/home/root/.ssh
scp -r publicKey.pem USERNAME@xxx.xxx.x.xxx:/home/root/.ssh
scp -r privateKey.pem USERNAME@xxx.xxx.x.xxx:/home/root/.ssh
```

where USERNAME@xxx.xxx.x.xxx is the username and IP address you set for your board.


```golang
package main

 import (
         "crypto/dsa"
         "crypto/md5"
         "crypto/rand"
         "encoding/asn1"
         "encoding/gob"
         "encoding/pem"
         "fmt"
         "hash"
         "io"
         "math/big"
         "os"
 )

 func main() {

         params := new(dsa.Parameters)

         // see http://golang.org/pkg/crypto/dsa/#ParameterSizes
         if err := dsa.GenerateParameters(params, rand.Reader, dsa.L1024N160); err != nil {
                 fmt.Println(err)
                 os.Exit(1)
         }

         privatekey := new(dsa.PrivateKey)
         privatekey.PublicKey.Parameters = *params
         dsa.GenerateKey(privatekey, rand.Reader) // this generates a public & private key pair

         var pubkey dsa.PublicKey
         pubkey = privatekey.PublicKey

         fmt.Println("Private Key :")
         fmt.Printf("%x \n", privatekey)

         fmt.Println("Public Key :")
         fmt.Printf("%x \n", pubkey)

         // save private and public key separately
         privatekeyfile, err := os.Create("DSAprivate.key")
         if err != nil {
                 fmt.Println(err)
                 os.Exit(1)
         }
         privatekeyencoder := gob.NewEncoder(privatekeyfile)
         privatekeyencoder.Encode(privatekey)
         privatekeyfile.Close()

         publickeyfile, err := os.Create("DSApublic.key")
         if err != nil {
                 fmt.Println(err)
                 os.Exit(1)
         }

         publickeyencoder := gob.NewEncoder(publickeyfile)
         publickeyencoder.Encode(pubkey)
         publickeyfile.Close()

         // save DSA public key to PEM encoded file
         pemfile, err := os.Create("DSApublickey.pem")

         if err != nil {
                 fmt.Println(err)
                 os.Exit(1)
         }

         // http://golang.org/pkg/encoding/pem/#Block
         // The decoded bytes of the contents. Typically a DER encoded ASN.1 structure.

         asn1Bytes, err := asn1.Marshal(pubkey)

         if err != nil {
                 fmt.Println(err)
                 os.Exit(1)
         }

         var pemkey = &pem.Block{
                 Type:  "PUBLIC KEY",
                 Bytes: asn1Bytes}

         err = pem.Encode(pemfile, pemkey)

         if err != nil {
                 fmt.Println(err)
                 os.Exit(1)
         }

         pemfile.Close()

         // ------------------------------
         // below here is bonus

         // Sign
         var h hash.Hash
         h = md5.New()
         r := big.NewInt(0)
         s := big.NewInt(0)

         io.WriteString(h, "This is the message to be signed and verified!")
         signhash := h.Sum(nil)

         r, s, err = dsa.Sign(rand.Reader, privatekey, signhash)
         if err != nil {
                 fmt.Println(err)
         }

         signature := r.Bytes()
         signature = append(signature, s.Bytes()...)

         fmt.Printf("Signature : %x\n", signature)

         // Verify
         verifystatus := dsa.Verify(&pubkey, signhash, r, s)
         fmt.Println(verifystatus) // should be true

         // we add additional data to change the signhash
         io.WriteString(h, "This message is NOT to be signed and verified!")
         signhash = h.Sum(nil)

         verifystatus = dsa.Verify(&pubkey, signhash, r, s)
         fmt.Println(verifystatus) // should be false
 }
 ```
