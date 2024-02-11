FROM golang:1.21.2

WORKDIR /go/src/app

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

ENV GO111MODULE=on

RUN go build -o main

EXPOSE 3001

CMD ["./main", "-addr=:3001"]
