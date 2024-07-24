# Utiliser une image de base Go
FROM golang:1.22.5

# Définir le répertoire de travail à l'intérieur du conteneur
WORKDIR /app

# Désactiver la vérification SSL pour les commandes Go
ENV GONOSUMDB=*
ENV GOPRIVATE=*
ENV GOSUMDB=off

# Copier les fichiers go.mod et go.sum pour la gestion des dépendances
COPY go.mod go.sum ./

# Télécharger les dépendances
RUN go get -d ./...

# Copier le reste de l'application source
COPY . .

# Construire l'application
RUN go build -o main .

# Exposer le port sur lequel l'application s'exécute
EXPOSE 8080

# Définir la commande par défaut pour exécuter l'application
CMD ["./main"]