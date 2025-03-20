-- Quiz
INSERT INTO quizzes (title, description, category, is_test, time_limit, created_at, updated_at)
VALUES 
  ('Introduction à Go', 'Testez vos connaissances sur les bases de Go', 'Programmation', false, 15, NOW(), NOW()),
  ('Test de certification Go', 'Évaluez vos compétences en Go', 'Programmation', true, 60, NOW(), NOW()),
  ('Docker pour débutants', 'Quiz sur les concepts de base de Docker', 'DevOps', false, 20, NOW(), NOW());

-- Questions pour le quiz "Introduction à Go"
INSERT INTO questions (quiz_id, text, explanation, points, created_at, updated_at)
VALUES 
  (1, 'Quelle est la commande pour créer un nouveau module Go ?', 'go mod init est la commande de base pour initialiser un nouveau module', 2, NOW(), NOW()),
  (1, 'Comment déclarer une variable en Go ?', 'Il existe plusieurs façons de déclarer des variables en Go', 1, NOW(), NOW());

-- Options pour les questions
INSERT INTO options (question_id, text, is_correct, created_at, updated_at)
VALUES 
  (1, 'go mod init', true, NOW(), NOW()),
  (1, 'go new module', false, NOW(), NOW()),
  (1, 'go init', false, NOW(), NOW()),
  (2, 'var x int', true, NOW(), NOW()),
  (2, 'int x;', false, NOW(), NOW()),
  (2, 'x := 5', true, NOW(), NOW());

-- Exercices
INSERT INTO exercises (title, description, category, difficulty, docker_image, instructions, created_at, updated_at)
VALUES 
  ('Hello World en Go', 'Créez votre premier programme Go', 'Programmation', 'easy', 'golang:1.21-alpine', 'Écrivez un programme qui affiche "Hello, World!"', NOW(), NOW()),
  ('API REST simple', 'Créez une API REST basique avec Fiber', 'Web', 'medium', 'golang:1.21-alpine', 'Implémentez une API avec les opérations CRUD', NOW(), NOW()),
  ('Conteneurisation', 'Conteneurisez une application Go', 'DevOps', 'hard', 'docker:latest', 'Créez un Dockerfile pour une application Go', NOW(), NOW());

-- Fichiers initiaux pour les exercices
INSERT INTO files (exercise_id, path, content, is_readonly, created_at, updated_at)
VALUES 
  (1, 'main.go', 'package main\n\nfunc main() {\n  // Votre code ici\n}', false, NOW(), NOW()),
  (2, 'main.go', 'package main\n\nimport "github.com/gofiber/fiber/v2"\n\nfunc main() {\n  // Votre code ici\n}', false, NOW(), NOW()),
  (3, 'Dockerfile', '# Écrivez votre Dockerfile ici', false, NOW(), NOW());

-- Tests pour les exercices
INSERT INTO test_cases (exercise_id, name, command, expected, points, created_at, updated_at)
VALUES 
  (1, 'Test Hello World', 'go run main.go', 'Hello, World!', 1, NOW(), NOW()),
  (2, 'Test GET /', 'curl http://localhost:3000', '{"message":"Hello, World!"}', 2, NOW(), NOW()),
  (3, 'Test Docker Build', 'docker build -t test .', '', 3, NOW(), NOW());

-- Progrès utilisateur initial
INSERT INTO user_progresses (user_id, courses_started, courses_completed, quizzes_taken, quiz_avg_score, exercises_done, total_points, study_time, last_active, created_at, updated_at)
VALUES 
  (1, 2, 1, 3, 85.5, 2, 150, 120, NOW(), NOW(), NOW());

-- Classement initial
INSERT INTO rankings (user_id, score, level, position, category, created_at, updated_at)
VALUES 
  (1, 150, 2, 1, 'global', NOW(), NOW());

-- Achievements
INSERT INTO achievements (title, description, icon, condition, points, created_at, updated_at)
VALUES 
  ('Premier Pas', 'Complétez votre premier cours', '🎯', 'courses_completed >= 1', 10, NOW(), NOW()),
  ('Quiz Master', 'Obtenez 100% à 5 quiz', '🏆', 'perfect_quizzes >= 5', 50, NOW(), NOW()),
  ('Code Warrior', 'Complétez 10 exercices', '⚔️', 'exercises_completed >= 10', 100, NOW(), NOW());

-- Badges
INSERT INTO badges (name, description, icon, min_level, created_at, updated_at)
VALUES 
  ('Débutant', 'Niveau débutant atteint', '🌱', 1, NOW(), NOW()),
  ('Intermédiaire', 'Niveau intermédiaire atteint', '🌿', 5, NOW(), NOW()),
  ('Expert', 'Niveau expert atteint', '🌳', 10, NOW(), NOW()); 