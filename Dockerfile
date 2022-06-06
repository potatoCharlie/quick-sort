# syntax=docker/dockerfile:1

FROM golang:1.18-alpine

WORKDIR /quick-sort

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY *.go ./
COPY lib ./lib

RUN go build -o /quick-sort

EXPOSE 8080

CMD [ "./quick-sort" ]