FROM alpine:3.1
MAINTAINER Joey Grasso <me@joeygrasso.com>
ADD resum-api /usr/bin/resum-api
ENTRYPOINT ["resum-api"]
