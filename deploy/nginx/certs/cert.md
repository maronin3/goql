# SSL Certificate for NGINX

The certificate which was generated here follows these configurations;

```
 key; 210730
  country code; 
  state; 
  city; 
  company; 
  division; 
  commonname; 

  day; 385000

  pass:210730
```

Were made using these commands;

Assuming we're doing this inside a blank folder--

```bash
$ openssl genrsa -aes256 -passout pass:210730 -out obs.pass.key 4096
$ openssl rsa -passin pass:210730 -in obs.pass.key -out obs.key
$ openssl req -new -key obs.key -out obs.csr  
$ openssl x509 -req -sha256 -days 385000 -in obs.csr -signkey obs.key -out obs.crt
```
