FROM golang:latest

MAINTAINER Joey Grasso <me@joeygrasso.com>

WORKDIR /opt/

ADD ./resum-api.go /

RUN go build -o /opt/resum-api /resum-api.go

CMD ["/opt/resum-api"]
