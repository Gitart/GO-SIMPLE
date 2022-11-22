package main     

import (                                                                                                                                                                                                    
    "encoding/base64"                                                                                                                                                                                       
    "crypto/aes"                                                                                                                                                                                            
    "crypto/cipher"                                                                                                                                                                                         
    "fmt"                                                                                                                                                                                                   
)

func main() {
     key        := "123456789012345678901234"
     plaintext1 := "255.255.00.00:4444"
     plaintext2 := "hello"
     foo        := Encrypt(key, plaintext1)

     fmt.Println(foo)
     fmt.Println(Decrypt(key, foo))

     bar := Encrypt(key, plaintext2)
     fmt.Println(bar)
     fmt.Println(Decrypt(key, bar))
}                                                                                                                                                                                                           
                                                                                                                                                                                                            
var iv = []byte{35, 46, 57, 24, 85, 35, 24, 74, 87, 35, 88, 98, 66, 32, 14, 05}                                                                                                                             
                                                                                                                                                                                                            
func encodeBase64(b []byte) string {                                                                                                                                                                        
    return base64.StdEncoding.EncodeToString(b)                                                                                                                                                             
}                                                                                                                                                                                                           
                                                                                                                                                                                                            
func decodeBase64(s string) []byte {                                                                                                                                                                        
    data, err := base64.StdEncoding.DecodeString(s)                                                                                                                                                         
    if err != nil { panic(err) }                                                                                                                                                                            
    return data                                                                                                                                                                                             
}                                                                                                                                                                                                           
                                                                                                                                                                                                            
func Encrypt(key, text string) string {                                                                                                                                                                     
    fmt.Println(text)                                                                                                                                                                                       

    block, err := aes.NewCipher([]byte(key))  
                                                                                                                                                             
    if err != nil { panic(err) }                                                                                                                                                                            
    plaintext := []byte(text)                                                                                                                                                                               
    cfb := cipher.NewCFBEncrypter(block, iv)                                                                                                                                                                
    ciphertext := make([]byte, len(plaintext))                                                                                                                                                              
    cfb.XORKeyStream(ciphertext, plaintext)                                                                                                                                                                 
    return encodeBase64(ciphertext)                                                                                                                                                                         
}                                                                                                                                                                                                           
                                                                                                                                                                                                            
func Decrypt(key, text string) string {                                                                                                                                                                     
    fmt.Println(text)                                                                                                                                                                                       
    block, err := aes.NewCipher([]byte(key))                                                                                                                                                                
    if err != nil { panic(err) }                                                                                                                                                                            
    ciphertext := decodeBase64(text)                                                                                                                                                                        
    cfb := cipher.NewCFBEncrypter(block, iv)                                                                                                                                                                
    plaintext := make([]byte, len(ciphertext))                                                                                                                                                              
    cfb.XORKeyStream(plaintext, ciphertext)                                                                                                                                                                 
    return string(plaintext)                                                                                                                                                                                
}                                                                                                                                                                                                           
