# Используем официальный образ Go для сборки
FROM golang:1.23 as builder

# Устанавливаем рабочую директорию в контейнере
WORKDIR /server

# Копируем файлы проекта в контейнер
COPY go.mod ./
#COPY go.sum ./
RUN go mod download

COPY . .

# Сборка приложения
RUN go build -o ginol-server ./src/main.go

# Используем минимальный базовый образ
FROM gcr.io/distroless/base-debian11
WORKDIR /server

# Копируем собранное приложение из предыдущего контейнера
COPY --from=builder /server/ginol-server /server/ginol-server

# Определяем порт, который будет использоваться
EXPOSE 8080

# Запускаем приложение
CMD ["/server/ginol-server"]
