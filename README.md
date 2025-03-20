# LearnIT

Application web d'apprentissage et de préparation aux certifications.

## Structure du Projet

```
.
├── frontend/          # Application Next.js
└── backend/          # Serveur Go
```

## Fonctionnalités

- 📚 Apprentissage : Cours et ressources PDF par chapitre
- 💻 Environnement d'exercices : Pratique sur environnement Docker
- ✍️ Quiz : Questions à choix multiples avec explications
- 🎯 Tests réels : Simulation d'examens
- 📊 Tableau de bord : Suivi des performances
- 👑 Classement : Progression et statistiques
- 🔐 Interface administrateur : Gestion du contenu
- 🚪 Authentification : Système de connexion sécurisé

## Prérequis

- Go 1.21+
- Node.js 18+
- PostgreSQL 15+
- Docker

## Installation

### Backend (Go)

```bash
cd backend
go mod download
go run main.go
```

### Frontend (Next.js)

```bash
cd frontend
npm install
npm run dev
```

## Développement Local

Le projet est configuré pour fonctionner en local avec :
- Frontend : http://localhost:3000
- Backend : http://localhost:8080
- Base de données : PostgreSQL sur le port 5432

## Technologies Utilisées

### Frontend
- Next.js (React)
- TypeScript
- Tailwind CSS
- Framer Motion

### Backend
- Go
- Fiber
- PostgreSQL
- JWT 