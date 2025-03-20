#!/bin/bash

# Configuration
DOCKER_USERNAME="petitsinge"  # Remplacez par votre nom d'utilisateur Docker Hub
VERSION="latest"

echo "Construction et push des images Docker..."

# Frontend
echo "Construction de l'image frontend..."
docker build -t $DOCKER_USERNAME/learnit-frontend:$VERSION frontend/

echo "Push de l'image frontend vers Docker Hub..."
docker push $DOCKER_USERNAME/learnit-frontend:$VERSION

# Backend
echo "Construction de l'image backend..."
docker build -t $DOCKER_USERNAME/learnit-backend:$VERSION backend/

echo "Push de l'image backend vers Docker Hub..."
docker push $DOCKER_USERNAME/learnit-backend:$VERSION

echo "Images construites et poussées avec succès !" 