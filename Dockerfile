# syntax=docker/dockerfile:1

FROM golang:1.17-alpine

WORKDIR /app

#COPY go.mod ./
#COPY go.sum ./
#RUN go mod download
#
#COPY *.go ./

COPY go_web ./

RUN go build -o /go_web
EXPOSE 8080

CMD ["/go_web"]