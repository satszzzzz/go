FROM golang:1.22.12

WORKDIR /app
COPY . .
RUN go test ./...
RUN go build -o main .

CMD ["./main"]