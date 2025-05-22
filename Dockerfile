# Этап 1: Сборка приложения
FROM golang:1.23-alpine AS builder

# Устанавливаем зависимости, необходимые для сборки
RUN apk add --no-cache git ca-certificates

# Создаем рабочую директорию
WORKDIR /app

# Копируем файлы модулей (для эффективного кэширования)
COPY go.mod go.sum ./

# Скачиваем зависимости
RUN go mod download

# Копируем исходный код
COPY . .

# Собираем приложение
# Убедитесь, что ваш main package находится в ./cmd/bot/
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -o /go/bin/bot ./cmd/bot

# Этап 2: Создаем минимальный образ для запуска
FROM alpine:latest

# Устанавливаем зависимости времени выполнения
RUN apk add --no-cache ca-certificates tzdata

# Копируем бинарник из этапа сборки
COPY --from=builder /go/bin/bot /bot

# Копируем статические файлы (если есть)
# COPY --from=builder /app/static /static

COPY .env .

# Указываем точку входа
CMD ["/bot"]