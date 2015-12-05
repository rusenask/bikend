FROM golang:1.5-onbuild

MAINTAINER karolis.rusenas@gmail.com

ADD . /go/src/github.com/rusenask/bikend

ENV MongoURI=mongo

ENV GO15VENDOREXPERIMENT 1

RUN go install github.com/rusenask/bikend

ENTRYPOINT /go/bin/bikend

EXPOSE 80
