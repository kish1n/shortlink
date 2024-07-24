FROM golang:1.22-alpine as buildbase

RUN apk add git build-base

WORKDIR /go/src/github.com/kish1n/shortlink
COPY vendor .
COPY . .

RUN GOOS=linux go build  -o /usr/local/bin/shortlink /go/src/github.com/kish1n/shortlink


FROM alpine:3.9

COPY --from=buildbase /usr/local/bin/shortlink /usr/local/bin/shortlink
COPY config.yaml /usr/local/bin/config.yaml
RUN apk add --no-cache ca-certificates

ENTRYPOINT ["shortlink"]
CMD ["run", "service"]

#trpdjke@Trpdjke:~/go/src/github.com/kish1n/shortlink$ export POSTGRES_DB=postgres
#export POSTGRES_USER=postgres
#export POSTGRES_PASSWORD=1029
#export DB_HOST=localhost
#export DB_PORT=5432
#export KV_VIPER_FILE=./config.yaml
#trpdjke@Trpdjke:~/go/src/github.com/kish1n/shortlink$ go run main.go run service

#curl -X POST -H "Content-Type: application/json" -d '{"original": "https://www.example.com"}' http://localhost:8080/integrations/shortlink/add
