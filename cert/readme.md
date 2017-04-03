## Note
You can buy SSL certificates from many sources. However, I will recommend comodo
security provider (https://ssl.comodo.com/). Comodo also provides free SSL certificate, learn
more about it by visiting this link ( https://ssl.comodo.com/free-ssl-certificate.php)

### You need to generate a key and a certificate file and send it to the security provider (if you are
buying it) for verification. Here is a sample command to generate a key and a certificate file using openssl.

#### Generate a key as follows:
```
openssl genrsa -out key.pem 2048
```
#### Now, generate a certificate file as follows:
```
openssl req -new -x509 -key key.pem -out cert.pem -days 365
```

You need to answer a few questions in the setup wizard, but make sure that the common name matches the domain name of your server.
