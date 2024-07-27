FROM golang:1.22.1-alpine

ENV JWT_PRIVATE_KEY=h/M3hOr9mTkeZgnYtLxOUIfhK9kQXAU+hgW7pR84xAQ=

RUN apk update && \
    apk add --no-cache git dos2unix

RUN mkdir /app

WORKDIR /app

RUN go install github.com/cloudflare/cfssl/cmd/cfssl@latest && \
    go install github.com/cloudflare/cfssl/cmd/cfssljson@latest


COPY . /app
COPY csrGenerationScript.sh /app
RUN dos2unix /app/csrGenerationScript.sh
RUN chmod +x /app/csrGenerationScript.sh
RUN go build -o binaryFile .


ENV GRPC_GO_LOG_VERBOSITY_LEVEL=99
ENV GRPC_GO_LOG_SEVERITY_LEVEL=info


RUN chmod +x /app/binaryFile
CMD ["/app/binaryFile"]