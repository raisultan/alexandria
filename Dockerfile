FROM golang:1.14.7-alpine

RUN apk update && apk upgrade && \
    apk add --no-cache bash git openssh

LABEL maintainer="raisultan karimov <ki.xbozz@gmail.com>"

WORKDIR /app

COPY ./app/go.mod /app
COPY ./app/go.sum /app

RUN go mod download

COPY ./app /app

RUN go build -o main .

CMD ["./main"]
