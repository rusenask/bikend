FROM golang:1.5

MAINTAINER karolis.rusenas@gmail.com

ADD . /go/src/github.com/rusenask/bikend

RUN go get gopkg.in/mgo.v2
RUN go get github.com/Sirupsen/logrus
RUN go get github.com/codegangsta/negroni
RUN go get  github.com/go-zoo/bone
RUN go get github.com/meatballhat/negroni-logrus
RUN go get github.com/unrolled/render

ENV MongoURI mongo

#ENV GO15VENDOREXPERIMENT 1

RUN go install github.com/rusenask/bikend

ENTRYPOINT /go/bin/bikend

EXPOSE 80
