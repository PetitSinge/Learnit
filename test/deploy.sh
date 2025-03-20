#!/bin/bash

# Fonction pour vérifier les erreurs
check_error() {
    if [ $? -ne 0 ]; then
        echo "❌ Erreur: $1"
        exit 1
    fi
}

# Se connecter à OpenShift (à adapter selon votre configuration)
# oc login --token=<your-token> --server=<your-server>

# Créer un nouveau projet si nécessaire
oc new-project learnit

echo "🚀 Déploiement de l'application LearnIT..."

# Appliquer les secrets
echo "1/9 Application des secrets..."
oc apply -f secrets/
check_error "Échec de l'application des secrets"

# Créer le PVC
echo "2/9 Création du PVC..."
oc apply -f pvc/
check_error "Échec de la création du PVC"

# Créer les BuildConfigs et ImageStreams
echo "3/9 Création des configurations de build..."
oc apply -f buildconfigs/
check_error "Échec de la création des BuildConfigs"

# Lancer les builds
echo "4/9 Construction des images..."
echo "Construction du frontend..."
oc start-build frontend-build --follow
check_error "Échec de la construction du frontend"

echo "Construction du backend..."
oc start-build backend-build --follow
check_error "Échec de la construction du backend"

# Déployer PostgreSQL
echo "5/9 Déploiement de PostgreSQL..."
oc apply -f services/postgres-service.yaml
oc apply -f deployments/postgres-deployment.yaml
check_error "Échec du déploiement de PostgreSQL"

# Attendre que PostgreSQL soit prêt
echo "Attente du démarrage de PostgreSQL..."
for i in {1..30}; do
    if oc wait --for=condition=available deployment/postgres --timeout=10s; then
        break
    fi
    echo "Tentative $i/30..."
    if [ $i -eq 30 ]; then
        echo "❌ Timeout en attendant PostgreSQL"
        exit 1
    fi
done

# Déployer le backend
echo "6/9 Déploiement du backend..."
oc apply -f services/backend-service.yaml
oc apply -f deployments/backend-deployment.yaml
check_error "Échec du déploiement du backend"

# Attendre que le backend soit prêt
echo "Attente du démarrage du backend..."
for i in {1..30}; do
    if oc wait --for=condition=available deployment/backend --timeout=10s; then
        break
    fi
    echo "Tentative $i/30..."
    if [ $i -eq 30 ]; then
        echo "❌ Timeout en attendant le backend"
        exit 1
    fi
done

# Déployer le frontend
echo "7/9 Déploiement du frontend..."
oc apply -f services/frontend-service.yaml
oc apply -f deployments/frontend-deployment.yaml
check_error "Échec du déploiement du frontend"

# Attendre que le frontend soit prêt
echo "Attente du démarrage du frontend..."
for i in {1..30}; do
    if oc wait --for=condition=available deployment/frontend --timeout=10s; then
        break
    fi
    echo "Tentative $i/30..."
    if [ $i -eq 30 ]; then
        echo "❌ Timeout en attendant le frontend"
        exit 1
    fi
done

# Créer la route
echo "8/9 Création de la route..."
oc apply -f routes/
check_error "Échec de la création de la route"

echo "✅ 9/9 Déploiement terminé !"
echo "🌐 URL de l'application :"
oc get route learnit-route -o jsonpath='{.spec.host}'

# Afficher les logs en cas d'erreur
echo -e "\n📋 Vérification des logs :"
echo "PostgreSQL logs :"
oc logs deployment/postgres
echo -e "\nBackend logs :"
oc logs deployment/backend
echo -e "\nFrontend logs :"
oc logs deployment/frontend 