# syntax=docker/dockerfile:

FROM golang:1.21-alpine

WORKDIR /app

COPY go.mod ./

RUN go mod download

COPY *.go ./

RUN go build -o /octvis

EXPOSE 8090

CMD [ "/octvis" ]
