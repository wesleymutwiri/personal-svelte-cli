FROM golang:1.15-alpine3.12

RUN apk add git

RUN mkdir /app

ADD . /app

WORKDIR /app

ENV GO111MODULE=on

RUN go mod download

RUN go build -o main .

CMD "/app/main"