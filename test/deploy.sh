#!/bin/bash

# Se connecter à OpenShift (à adapter selon votre configuration)
# oc login --token=<your-token> --server=<your-server>

# Créer un nouveau projet si nécessaire
oc new-project learnit

echo "Déploiement de l'application LearnIT..."

# Appliquer les secrets
echo "1/9 Application des secrets..."
oc apply -f secrets/

# Créer le PVC
echo "2/9 Création du PVC..."
oc apply -f pvc/

# Créer les BuildConfigs et ImageStreams
echo "3/9 Création des configurations de build..."
oc apply -f buildconfigs/

# Lancer les builds
echo "4/9 Construction des images..."
oc start-build frontend-build --follow
oc start-build backend-build --follow

# Déployer PostgreSQL
echo "5/9 Déploiement de PostgreSQL..."
oc apply -f services/postgres-service.yaml
oc apply -f deployments/postgres-deployment.yaml

# Attendre que PostgreSQL soit prêt
echo "Attente du démarrage de PostgreSQL..."
oc wait --for=condition=available deployment/postgres --timeout=300s

# Déployer le backend
echo "6/9 Déploiement du backend..."
oc apply -f services/backend-service.yaml
oc apply -f deployments/backend-deployment.yaml

# Attendre que le backend soit prêt
echo "Attente du démarrage du backend..."
oc wait --for=condition=available deployment/backend --timeout=300s

# Déployer le frontend
echo "7/9 Déploiement du frontend..."
oc apply -f services/frontend-service.yaml
oc apply -f deployments/frontend-deployment.yaml

# Attendre que le frontend soit prêt
echo "Attente du démarrage du frontend..."
oc wait --for=condition=available deployment/frontend --timeout=300s

# Créer la route
echo "8/9 Création de la route..."
oc apply -f routes/

echo "9/9 Déploiement terminé !"
echo "URL de l'application :"
oc get route learnit-route -o jsonpath='{.spec.host}' 