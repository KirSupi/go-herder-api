FROM golang:1.19

RUN mkdir /app
WORKDIR /
ADD . /app
WORKDIR /app/cmd/api
RUN go build -o ./../../main
EXPOSE 80
WORKDIR /app
CMD ["/app/main"]
