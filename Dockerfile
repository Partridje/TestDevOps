# Используем официальный образ Go в качестве базового
FROM golang:latest

# Устанавливаем рабочую директорию внутри контейнера
WORKDIR /app

# Копируем файлы модулей и загружаем зависимости
COPY go.mod ./
COPY go.sum ./
RUN go mod download

# Копируем исходный код из поддиректории cv-devops в рабочую директорию контейнера
COPY cv-devops/*.go ./
COPY cv-devops/resume.html ./

# Собираем приложение
RUN go build -o cv-devops

# Указываем порт, на котором будет работать приложение
EXPOSE 8080

# Команда для запуска приложения
CMD ["./cv-devops"]
