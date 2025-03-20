'use client';

import { useState, useEffect } from 'react';
import Link from 'next/link';

interface Quiz {
  id: number;
  title: string;
  description: string;
  category: string;
  is_test: boolean;
  time_limit: number;
}

export default function QuizPage() {
  const [quizzes, setQuizzes] = useState<Quiz[]>([]);
  const [loading, setLoading] = useState(true);

  useEffect(() => {
    const fetchQuizzes = async () => {
      try {
        const response = await fetch('http://localhost:8080/api/v1/quiz', {
          credentials: 'include',
        });
        if (response.ok) {
          const data = await response.json();
          setQuizzes(data);
        }
      } catch (error) {
        console.error('Erreur lors de la récupération des quiz:', error);
      } finally {
        setLoading(false);
      }
    };

    fetchQuizzes();
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
      <h1 className="text-3xl font-bold mb-8">Quiz disponibles</h1>
      <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6">
        {quizzes.map((quiz) => (
          <Link
            key={quiz.id}
            href={`/quiz/${quiz.id}`}
            className="block bg-white rounded-lg shadow-md hover:shadow-lg transition-shadow p-6"
          >
            <div className="flex items-center justify-between mb-4">
              <h2 className="text-xl font-semibold">{quiz.title}</h2>
              {quiz.is_test && (
                <span className="bg-yellow-100 text-yellow-800 text-xs font-medium px-2.5 py-0.5 rounded">
                  Test
                </span>
              )}
            </div>
            <p className="text-gray-600 mb-4">{quiz.description}</p>
            <div className="flex justify-between items-center text-sm text-gray-500">
              <span>{quiz.category}</span>
              {quiz.time_limit > 0 && (
                <span>Temps: {quiz.time_limit} min</span>
              )}
            </div>
          </Link>
        ))}
      </div>
    </div>
  );
} 