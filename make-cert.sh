openssl req -x509 -newkey rsa:4096 -sha256 -days 3650 \
  -nodes -keyout mkn.edu.key -out mkn.edu.crt -subj "//CN=mkn.edu" \
  -addext "subjectAltName=DNS:mkn.edu,DNS:*.mkn.edu,IP:10.0.0.1"