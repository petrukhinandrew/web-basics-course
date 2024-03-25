# /usr/bin/bash

openssl genrsa -out mkn.edu.key 2048

openssl req -new -sha256 \
    -key mkn.edu.key \
    -subj "/C=RU/ST=Spb/O=MKN/CN=mkn.edu" \
    -reqexts SAN \
    -config <(cat /etc/ssl/openssl.cnf \
        <(printf "\n[SAN]\nsubjectAltName=DNS:mkn.edu,DNS:mkn.edu")) \
    -out mkn.edu.csr
openssl req -in mkn.edu.csr -noout -text

openssl x509 -req -extfile <(printf "subjectAltName=DNS:mkn.edu") -in mkn.edu.csr -CA rootCA.crt -CAkey rootCA.key -CAcreateserial -out mkn.edu.crt -days 500 -sha256 -copy_extensions copyall
openssl x509 -in mkn.edu.crt -text -noout

