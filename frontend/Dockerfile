# Utiliser l'image de base Node.js pour le développement
FROM node:18 AS development

# Définir le répertoire de travail dans le conteneur
WORKDIR /app

# Copier les fichiers de configuration et de dépendances du frontend
COPY package.json package-lock.json ./

# Installer les dépendances
RUN npm install

# Copier le reste des fichiers du frontend 
COPY . .

# Exposer le port 5170 pour le serveur de développement
EXPOSE 5173

# Commande pour démarrer le serveur de développement
CMD ["npm", "run", "dev"]
