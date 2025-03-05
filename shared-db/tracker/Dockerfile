FROM golang:1.23.3

WORKDIR /app

COPY . .

RUN go mod download

RUN go build -o tracker ./main.go

EXPOSE 8080

CMD ["./tracker"]
