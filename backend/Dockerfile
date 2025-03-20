# Stage de build
FROM golang:1.21-alpine AS builder

WORKDIR /app

# Installation des dépendances nécessaires
RUN apk add --no-cache gcc musl-dev

# Copie des fichiers go.mod et go.sum
COPY go.mod go.sum ./
RUN go mod download

# Copie du reste du code source
COPY . .

# Build de l'application
RUN CGO_ENABLED=1 GOOS=linux go build -o main .

# Stage de production
FROM alpine:latest

WORKDIR /app

# Copie du binaire depuis le stage de build
COPY --from=builder /app/main .
COPY --from=builder /app/.env .

# Exposition du port
EXPOSE 8080

# Commande de démarrage
CMD ["./main"] 