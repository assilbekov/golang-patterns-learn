FROM golang:1.23.3

WORKDIR /app

COPY . .

RUN go mod download

RUN go build -o app tracker/main.go

EXPOSE 8080

RUN chmod +x app

CMD ["/app/app"]
