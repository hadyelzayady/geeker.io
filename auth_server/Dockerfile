FROM golang:1.16.5-buster

WORKDIR /go/src/app

COPY . .

RUN go get -d -v ./...
RUN go install -v ./...


CMD ["authserver.xyz"]
