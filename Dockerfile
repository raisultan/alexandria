FROM golang:1.14.7-alpine

LABEL maintainer="raisultan"

ENV ALEXANDRIA_YAML "/app/examples/ws.yaml"

WORKDIR /app

COPY ./app/go.mod /app
COPY ./app/go.sum /app

RUN go mod download

COPY ./app /app

RUN go build -o main .

EXPOSE 8080

CMD ["./main"]
