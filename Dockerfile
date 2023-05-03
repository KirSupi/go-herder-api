FROM golang:1.19

RUN mkdir /app
WORKDIR /
ADD . /app
WORKDIR /app/cmd/api
RUN go build -o ./../../main

WORKDIR /app
CMD ["/app/main"]
