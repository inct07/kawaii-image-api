FROM golang:1.10.1-stretch

WORKDIR /go/src/kawaii-image-api
COPY . .

RUN go build -o /usr/local/bin/kawaii-image-api

CMD ["kawaii-image-api"]