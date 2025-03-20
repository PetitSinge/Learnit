'use client';

import { useState, useEffect } from 'react';
import { useRouter } from 'next/navigation';

interface Chapter {
  id: number;
  title: string;
  description: string;
  content: string;
  order: number;
  course_id: number;
}

export default function ChapterDetailPage({
  params,
}: {
  params: { courseId: string; chapterId: string };
}) {
  const router = useRouter();
  const [chapter, setChapter] = useState<Chapter | null>(null);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState('');

  useEffect(() => {
    const fetchChapter = async () => {
      try {
        const token = localStorage.getItem('token');
        if (!token) {
          router.push('/auth/login');
          return;
        }

        const response = await fetch(
          `http://localhost:8080/api/v1/courses/${params.courseId}/chapters/${params.chapterId}`,
          {
            headers: {
              'Authorization': `Bearer ${token}`,
            },
          }
        );

        if (!response.ok) {
          if (response.status === 401) {
            router.push('/auth/login');
            return;
          }
          throw new Error('Erreur lors de la récupération du chapitre');
        }

        const data = await response.json();
        setChapter(data.chapter);
      } catch (err: any) {
        setError(err.message);
      } finally {
        setLoading(false);
      }
    };

    fetchChapter();
  }, [params.courseId, params.chapterId, router]);

  if (loading) {
    return (
      <div className="min-h-screen flex items-center justify-center">
        <div className="animate-spin rounded-full h-12 w-12 border-b-2 border-indigo-600"></div>
      </div>
    );
  }

  if (error || !chapter) {
    return (
      <div className="min-h-screen flex items-center justify-center">
        <div className="text-center">
          <h2 className="text-2xl font-bold text-gray-900 mb-4">
            {error || 'Chapitre non trouvé'}
          </h2>
          <button
            onClick={() => router.push(`/cours/${params.courseId}`)}
            className="text-indigo-600 hover:text-indigo-500"
          >
            Retour au cours
          </button>
        </div>
      </div>
    );
  }

  return (
    <div className="min-h-screen bg-gray-100">
      <div className="max-w-4xl mx-auto py-12 px-4 sm:px-6 lg:px-8">
        {/* Navigation */}
        <div className="mb-8">
          <button
            onClick={() => router.push(`/cours/${params.courseId}`)}
            className="flex items-center text-indigo-600 hover:text-indigo-500"
          >
            <svg
              className="h-5 w-5 mr-2"
              fill="none"
              viewBox="0 0 24 24"
              stroke="currentColor"
            >
              <path
                strokeLinecap="round"
                strokeLinejoin="round"
                strokeWidth={2}
                d="M15 19l-7-7 7-7"
              />
            </svg>
            Retour au cours
          </button>
        </div>

        {/* Contenu du chapitre */}
        <div className="bg-white rounded-lg shadow-lg overflow-hidden">
          <div className="p-8">
            <h1 className="text-3xl font-bold text-gray-900 mb-4">
              {chapter.title}
            </h1>
            <p className="text-lg text-gray-600 mb-8">{chapter.description}</p>
            
            <div className="prose max-w-none">
              {/* Ici, vous pourriez utiliser un composant Markdown si le contenu est au format Markdown */}
              <div dangerouslySetInnerHTML={{ __html: chapter.content }} />
            </div>
          </div>
        </div>

        {/* Navigation entre chapitres */}
        <div className="mt-8 flex justify-between">
          <button
            onClick={() => {
              // Implémenter la navigation vers le chapitre précédent
            }}
            className="text-indigo-600 hover:text-indigo-500 disabled:opacity-50 disabled:cursor-not-allowed"
            disabled={chapter.order === 1}
          >
            Chapitre précédent
          </button>
          <button
            onClick={() => {
              // Implémenter la navigation vers le chapitre suivant
            }}
            className="text-indigo-600 hover:text-indigo-500"
          >
            Chapitre suivant
          </button>
        </div>
      </div>
    </div>
  );
} 