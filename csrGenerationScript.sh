#!/bin/sh

# Read environment variable for container name
CONTAINER_NAME=${CONTAINER_NAME}

# Hardcoded paths for root CA certificate and key
ROOT_CA="/prod/root.pem"
ROOT_CA_KEY="/prod/root-key.pem"

# Create csr.json from template
cat << EOF > csr.json
{
  "hosts": [
    "${CONTAINER_NAME}"
  ],
  "CN": "${CONTAINER_NAME}",
  "key": {
    "algo": "rsa",
    "size": 2048
  },
  "names": []
}
EOF
# Print the content of csr.json (uncommented line)
cat csr.json
# Generate CSR and private key
cfssl genkey csr.json | cfssljson -bare EndServer

cp /prod/root.pem /etc/ssl/certs
update-ca-certificates

# Sign CSR with root CA
cfssl sign -ca=${ROOT_CA} -ca-key=${ROOT_CA_KEY} -config=/prod/cfssl.json -profile=EndServer EndServer.csr | cfssljson -bare EndServer

echo "CSR and private key generated and signed successfully."
