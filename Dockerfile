FROM golang:1.14-alpine as builder
WORKDIR /jfrog-cli-go
COPY . /jfrog-cli-go
RUN apk add --update git && sh build.sh
FROM alpine:3.7
ENV CI true
RUN apk add --no-cache bash tzdata ca-certificates curl
COPY --from=builder /jfrog-cli-go/jfrog /usr/local/bin/jfrog
RUN chmod +x /usr/local/bin/jfrog
