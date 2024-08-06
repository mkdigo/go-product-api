FROM golang:1.22

WORKDIR /go/src/app

COPY . .

EXPOSE 80

RUN go build -o main cmd/main.go

CMD ["./main"]