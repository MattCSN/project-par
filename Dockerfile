# Utiliser une image de base Go
FROM golang:1.18-alpine

# Définir le répertoire de travail à l'intérieur du conteneur
WORKDIR /app

# Copier les fichiers go.mod et go.sum pour la gestion des dépendances
COPY go.mod go.sum ./

# Télécharger les dépendances
RUN go mod download

# Copier le reste de l'application source
COPY . .

# Construire l'application
RUN go build -o main .

# Exposer le port sur lequel l'application s'exécute
EXPOSE 8080

# Définir la commande par défaut pour exécuter l'application
CMD ["./main"]