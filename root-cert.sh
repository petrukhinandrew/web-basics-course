# /usr/bin/bash

openssl genrsa -out rootCA.key 4096
openssl rsa -noout -text -in rootCA.key
openssl req -new -x509 -days 365 -key rootCA.key -out rootCA.crt
openssl x509 -noout -text -in rootCA.crt