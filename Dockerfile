FROM alpine:latest

MAINTAINER John McKenzie<jmckind@gmail.com>

COPY entrypoint.sh /usr/bin/canary
RUN chmod a+x /usr/bin/canary

RUN apk add --no-cache curl
RUN curl -L https://github.com/sequenceiq/docker-alpine-dig/releases/download/v9.10.2/dig.tgz|tar -xzv -C /usr/bin/

ENTRYPOINT ["/usr/bin/canary"]
