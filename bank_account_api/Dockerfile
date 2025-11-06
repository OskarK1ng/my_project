FROM golang:1.24 AS builder

WORKDIR /app

COPY go.mod ./
RUN go mod download

COPY .env .

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -o /app/main ./cmd/app

FROM alpine:latest

WORKDIR /root

COPY --from=builder /app/main .

EXPOSE 8080

CMD ["./main"]