FROM golang:1.15

WORKDIR /go/src/primeNumCalculator/client

COPY . /go/src/primeNumCalculator/client

RUN apt-get update --yes
RUN apt-get install --yes netcat
RUN go mod download

EXPOSE 50051/tcp
