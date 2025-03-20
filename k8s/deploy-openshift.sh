#!/bin/bash

# Se connecter à OpenShift (à adapter selon votre configuration)
# oc login --token=<your-token> --server=<your-server>

# Créer un nouveau projet si nécessaire
oc new-project learnit

# Appliquer les secrets
echo "Application des secrets..."
oc apply -f k8s/secrets.yaml

# Déployer PostgreSQL
echo "Déploiement de PostgreSQL..."
oc apply -f k8s/postgres-deployment.yaml

# Attendre que PostgreSQL soit prêt
echo "Attente du démarrage de PostgreSQL..."
oc wait --for=condition=available deployment/postgres --timeout=300s

# Déployer le backend
echo "Déploiement du backend..."
oc apply -f k8s/backend-deployment.yaml

# Attendre que le backend soit prêt
echo "Attente du démarrage du backend..."
oc wait --for=condition=available deployment/backend --timeout=300s

# Déployer le frontend
echo "Déploiement du frontend..."
oc apply -f k8s/frontend-deployment.yaml

# Attendre que le frontend soit prêt
echo "Attente du démarrage du frontend..."
oc wait --for=condition=available deployment/frontend --timeout=300s

# Créer la route
echo "Création de la route..."
oc apply -f k8s/route.yaml

# Afficher l'URL de l'application
echo "Déploiement terminé !"
echo "URL de l'application :"
oc get route learnit-route -o jsonpath='{.spec.host}' 