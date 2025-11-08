# -------- Stage 1: Build --------
FROM golang:1.24 AS builder

# Устанавливаем рабочую директорию
WORKDIR /app

# Копируем go.mod и go.sum для кеширования зависимостей
COPY go.mod ./

# Загружаем зависимости
RUN go mod download

# Копируем всё остальное
COPY . .

# Собираем бинарник
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -o /app/main ./cmd/app

# -------- Stage 2: Run --------
FROM alpine:latest

# Создаем рабочую директорию
WORKDIR /root

# Копируем бинарник из предыдущего stage
COPY --from=builder /app/main .

# Копируем .env, если нужен в контейнере
COPY .env .env

# Открываем порт (опционально)
EXPOSE 8081

# Указываем команду запуска
CMD ["./main"]