#!/bin/bash

# source https://gist.github.com/fntlnz/cf14feb5a46b2eda428e000157447309

DOMAIN=mydomain.com
DAYS=3650
COUNTRY=CH
STATE=ZH
ORG="MyOrg, Inc."

# Create Root Key
openssl genrsa -out rootCA.key 4096

# Create and self sign the Root Certificate
openssl req -x509 -new -nodes -key rootCA.key -subj "/C=${COUNTRY}/ST=${STATE}/O=${ORG}/CN=${DOMAIN}" -sha256 -days ${DAYS} -out rootCA.crt

# Create the certificate key
openssl genrsa -out ${DOMAIN}.key 2048

# Create the signing
openssl req -new -sha256 -key ${DOMAIN}.key -subj "/C=${COUNTRY}/ST=${STATE}/O=${ORG}/CN=${DOMAIN}" -out ${DOMAIN}.csr

# Verify the csr's content
openssl req -in ${DOMAIN}.csr -noout -text

# Generate the certificate
openssl x509 -req -in ${DOMAIN}.csr -CA rootCA.crt -CAkey rootCA.key -CAcreateserial -out ${DOMAIN}.crt -days ${DAYS} -sha256

# Verify the certificate's content
openssl x509 -in ${DOMAIN}.crt -text -noout