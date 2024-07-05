FROM golang:1.22.1-alpine

RUN apk update && \
    apk add --no-cache git

RUN mkdir /app

WORKDIR /app

COPY . /app

RUN go build -o binaryFile .

RUN chmod +x /app/binaryFile
CMD ["/app/binaryFile"]