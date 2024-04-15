rm certs/*.pem
cd certs

openssl req -x509 -newkey rsa:4096 -nodes -days 365 -keyout ca-key.pem -out ca-cert.pem

echo "CA's self-signed certificate"
openssl x509 -in ca-cert.pem -noout -text

echo "server CSR for certificate"
openssl req -newkey rsa:4096 -nodes -keyout server-key.pem -out server-req.pem


openssl x509 -req -in server-req.pem -days 60 -CA ca-cert.pem -CAkey ca-key.pem -CAcreateserial -out server-cert.pem -extfile server-ext.cnf
echo "Server's signed certificate"
openssl x509 -in server-cert.pem -noout -text