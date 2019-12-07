openssl req -x509 -newkey rsa:4096 -keyout key.pem -out tls.crt -days 365
openssl rsa -in key.pem -out tls.key