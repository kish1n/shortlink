FROM golang:1.20-alpine as buildbase

RUN apk add git build-base

WORKDIR /go/src/github.com/kish1n/shortlink
COPY vendor .
COPY . .

RUN GOOS=linux go build  -o /usr/local/bin/shortlink /go/src/github.com/kish1n/shortlink


FROM alpine:3.9

COPY --from=buildbase /usr/local/bin/shortlink /usr/local/bin/shortlink
RUN apk add --no-cache ca-certificates

ENTRYPOINT ["shortlink"]
