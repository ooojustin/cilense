FROM golang:alpine

ENV GIN_MODE=release

RUN apk add build-base

RUN mkdir /app
WORKDIR /app

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .
RUN go build -o ./cilense ./main.go

EXPOSE 8080

ENTRYPOINT ["./cilense"]
