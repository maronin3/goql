# SSL Certificate

### Generate SSL certificate
```bash
1. server.key
openssl genrsa -des3 -out server.key 2048

2. server.csr
openssl req -new -key server.key -out server.csr

3. openssl x509 -req -days 3650 -in server.csr -signkey server.key -out server.crt
```

### RSA private key and public key generation
```bash
1. private key (private.pem 1024bit)
openssl genrsa -out private.pem 1024

2. public key (private.pem -> public.pem)
openssl rsa -in private.pem -out public.pem -outform PEM -pubout
```