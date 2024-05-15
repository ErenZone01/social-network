# Utiliser l'image de base Golang
FROM golang:alpine AS builder

# Définir le répertoire de travail dans le conteneur
WORKDIR /app

# Copier le contenu du répertoire backend dans le répertoire de travail du conteneur
COPY . .

# Installe les dépendances
RUN apk add --no-cache sqlite-dev gcc musl-dev

# Build de l'application
RUN CGO_ENABLED=1 GOOS=linux go build -o main .

# Crée une nouvelle image pour le runtime
# FROM alpine:latest

# Définir le répertoire de travail dans le conteneur
# WORKDIR /app

# Copier l'exécutable de l'étape de construction précédente dans le répertoire de travail du conteneur
# COPY --from=builder /app/main .

# Exposer le port sur lequel le serveur Go s'exécute
EXPOSE 8080

# Commande pour démarrer le serveur Go
CMD ["./main"]
