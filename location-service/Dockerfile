FROM golang:1.25.4-alpine

WORKDIR /app

# Go mod yükleme
COPY go.mod go.sum ./
RUN go mod tidy

# Tüm proje dosyalarını kopyala
COPY . .

# Binary oluştur (cmd/main.go’yu build ediyoruz)
RUN go build -o location-service ./cmd/main.go

# Port expose (env’den okuyor)
EXPOSE 8081

# Container çalıştırıldığında
CMD ["./location-service"]
