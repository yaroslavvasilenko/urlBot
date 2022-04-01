FROM golang:1.17.1

WORKDIR /app

ADD . /app/

WORKDIR /app

RUN go build -o main /app/src/main

CMD ["/app/main"]