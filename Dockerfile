#Build Stage
FROM golang:1.22.1-alpine as builder

ENV GOPATH=/app
ENV GOBIN=/app/bin

RUN apk update && \
    apk add --no-cache git dos2unix

WORKDIR /app

COPY go.mod ./
COPY go.sum ./

RUN go mod download

RUN go install github.com/cloudflare/cfssl/cmd/cfssl@latest && \
    go install github.com/cloudflare/cfssl/cmd/cfssljson@latest

COPY . ./

RUN dos2unix /app/csrGenerationScript.sh
RUN chmod +x /app/csrGenerationScript.sh

RUN CGO_ENABLED=0 GOOS=linux go build -v -o binaryFile .

# Production Stage
FROM alpine:latest as production
RUN apk --no-cache add ca-certificates

RUN mkdir /prod
WORKDIR /prod
RUN mkdir /prod/bin

COPY --from=builder /app/binaryFile /prod/
COPY --from=builder /app/csrGenerationScript.sh /prod/
COPY ./root-key.pem /prod/
COPY ./root.pem /prod/
COPY ./cfssl.json /prod/
COPY --from=builder /app/bin /bin

RUN ls /bin

ENV GRPC_GO_LOG_VERBOSITY_LEVEL=99
ENV GRPC_GO_LOG_SEVERITY_LEVEL=info
ENV JWT_PRIVATE_KEY=h/M3hOr9mTkeZgnYtLxOUIfhK9kQXAU+hgW7pR84xAQ=

RUN chmod +x /prod/binaryFile

CMD ["/prod/binaryFile"]