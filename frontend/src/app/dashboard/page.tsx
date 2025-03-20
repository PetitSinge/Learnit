'use client';

import { useState, useEffect } from 'react';

interface Progress {
  courses_started: number;
  courses_completed: number;
  quizzes_taken: number;
  quiz_avg_score: number;
  exercises_done: number;
  total_points: number;
  study_time: number;
}

interface Achievement {
  id: number;
  title: string;
  description: string;
  icon: string;
  unlocked: boolean;
}

export default function DashboardPage() {
  const [progress, setProgress] = useState<Progress | null>(null);
  const [achievements, setAchievements] = useState<Achievement[]>([]);
  const [loading, setLoading] = useState(true);

  useEffect(() => {
    const fetchData = async () => {
      try {
        const [progressRes, achievementsRes] = await Promise.all([
          fetch('http://localhost:8080/api/v1/progress', {
            credentials: 'include',
          }),
          fetch('http://localhost:8080/api/v1/progress/achievements', {
            credentials: 'include',
          }),
        ]);

        if (progressRes.ok && achievementsRes.ok) {
          const [progressData, achievementsData] = await Promise.all([
            progressRes.json(),
            achievementsRes.json(),
          ]);
          setProgress(progressData);
          setAchievements(achievementsData);
        }
      } catch (error) {
        console.error('Erreur lors de la récupération des données:', error);
      } finally {
        setLoading(false);
      }
    };

    fetchData();
  }, []);

  if (loading) {
    return (
      <div className="flex justify-center items-center min-h-screen">
        <div className="animate-spin rounded-full h-12 w-12 border-t-2 border-b-2 border-blue-500"></div>
      </div>
    );
  }

  return (
    <div className="container mx-auto px-4 py-8">
      <h1 className="text-3xl font-bold mb-8">Tableau de bord</h1>
      
      {/* Statistiques */}
      <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 gap-6 mb-8">
        <div className="bg-white rounded-lg shadow p-6">
          <h3 className="text-lg font-semibold mb-2">Cours</h3>
          <div className="text-3xl font-bold text-blue-600">
            {progress?.courses_completed}/{progress?.courses_started}
          </div>
          <p className="text-gray-500">complétés</p>
        </div>

        <div className="bg-white rounded-lg shadow p-6">
          <h3 className="text-lg font-semibold mb-2">Quiz</h3>
          <div className="text-3xl font-bold text-green-600">
            {progress?.quiz_avg_score.toFixed(1)}%
          </div>
          <p className="text-gray-500">{progress?.quizzes_taken} quiz passés</p>
        </div>

        <div className="bg-white rounded-lg shadow p-6">
          <h3 className="text-lg font-semibold mb-2">Exercices</h3>
          <div className="text-3xl font-bold text-yellow-600">
            {progress?.exercises_done}
          </div>
          <p className="text-gray-500">exercices réalisés</p>
        </div>

        <div className="bg-white rounded-lg shadow p-6">
          <h3 className="text-lg font-semibold mb-2">Points</h3>
          <div className="text-3xl font-bold text-purple-600">
            {progress?.total_points}
          </div>
          <p className="text-gray-500">points gagnés</p>
        </div>
      </div>

      {/* Achievements */}
      <h2 className="text-2xl font-bold mb-4">Achievements</h2>
      <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6">
        {achievements.map((achievement) => (
          <div
            key={achievement.id}
            className={`bg-white rounded-lg shadow p-6 ${
              achievement.unlocked ? 'border-2 border-green-500' : 'opacity-50'
            }`}
          >
            <div className="flex items-center mb-4">
              <span className="text-2xl mr-3">{achievement.icon}</span>
              <h3 className="text-lg font-semibold">{achievement.title}</h3>
            </div>
            <p className="text-gray-600">{achievement.description}</p>
            {achievement.unlocked && (
              <div className="mt-4 text-green-600 text-sm font-medium">
                Débloqué ✓
              </div>
            )}
          </div>
        ))}
      </div>
    </div>
  );
} 