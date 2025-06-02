# Etapa de build
FROM golang:1.24.1-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -o audit-app ./cmd/server/main.go

# Imagen final
FROM alpine:latest

WORKDIR /root/
COPY --from=builder /app/audit-app .
EXPOSE 8080

ENV SMTP_HOST=${SMTP_HOST}
ENV SMTP_PORT=${SMTP_PORT}
ENV SMTP_USER=${SMTP_USER}
ENV SMTP_PASS=${SMTP_PASS}
ENV EMAIL_FROM=${EMAIL_FROM}

CMD ["./audit-app"]
