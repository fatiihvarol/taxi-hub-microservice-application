# api-gateway/Dockerfile
FROM golang:1.25.4-alpine


WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o api-gateway .

EXPOSE 8000

CMD ["./api-gateway"]
