# Étape 1 : builder le binaire
FROM golang:1.21 AS builder

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . ./
RUN go build -o myapp .

# Étape 2 : image minimale
FROM debian:bookworm-slim

WORKDIR /app
COPY --from=builder /app/myapp .

EXPOSE 3000

CMD ["./myapp"]