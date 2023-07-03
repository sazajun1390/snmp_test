FROM golang:1.20.5-alpine3.18

RUN apk update && apk add git

RUN mkdir -p /go/src/github.com/snmptest/cmd/hello

WORKDIR /go/src/github.com/snmptest

ADD . /go/src/github.com/snmptest

RUN cd cmd/hello

RUN go mod init snmptest

RUN go mod tidy

RUN go get -u github.com/gosnmp/gosnmp 

RUN go get -u github.com/sleepinggenius2/gosmi