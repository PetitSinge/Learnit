#!/bin/zsh

# Couleurs pour les logs
RED='\033[0;31m'
GREEN='\033[0;32m'
NC='\033[0m'

# Fonction de nettoyage
cleanup() {
    echo "Arrêt des serveurs..."
    pkill -f "go run main.go"
    pkill -f "next"
    docker-compose down
    exit 0
}

trap cleanup INT

# Vérifier si Docker est installé
if ! command -v docker &> /dev/null; then
    echo -e "${RED}Docker n'est pas installé. Veuillez l'installer d'abord.${NC}"
    exit 1
fi

# Configuration de Node.js
echo "Configuration de Node.js..."
source ~/.zshrc
nvm use 20

# Démarrer PostgreSQL avec Docker
echo "Démarrage de PostgreSQL..."
docker-compose down --remove-orphans &> /dev/null
docker-compose up -d

# Attendre que PostgreSQL soit prêt
echo "Attente du démarrage de PostgreSQL..."
until docker exec learnit_db pg_isready -h localhost -p 5432 -U learnit; do
    sleep 1
done

# Démarrer le backend
echo "Démarrage du backend..."
cd backend
go run main.go &

# Attendre que le backend soit prêt
echo "Attente du démarrage du backend..."
until curl -s http://localhost:8080/api/v1/health &> /dev/null; do
    sleep 1
done

# Démarrer le frontend
echo "Démarrage du frontend..."
cd ../frontend
npm run dev &

# Attendre que le frontend soit prêt
echo "Attente du démarrage du frontend..."
until curl -s http://localhost:3000 &> /dev/null; do
    sleep 1
done

echo -e "${GREEN}Application démarrée avec succès !${NC}"
echo "Frontend: http://localhost:3000"
echo "Backend API: http://localhost:8080"
echo "Admin credentials: admin/admin"

# Attendre que l'utilisateur arrête le script
wait