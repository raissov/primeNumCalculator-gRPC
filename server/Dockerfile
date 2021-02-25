FROM golang:1.15

WORKDIR /go/src/primeNumCalculator/server

COPY . /go/src/primeNumCalculator/server

RUN go mod download

EXPOSE 50051/tcp

CMD go run server.go
