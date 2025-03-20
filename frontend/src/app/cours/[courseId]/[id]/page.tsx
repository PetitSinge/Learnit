'use client';

import { useState, useEffect } from 'react';
import { useRouter } from 'next/navigation';

interface Chapter {
  id: number;
  title: string;
  description: string;
  order: number;
}

interface Course {
  id: number;
  title: string;
  description: string;
  image_url: string;
  chapters: Chapter[];
}

export default function CourseDetailPage({ params }: { params: { courseId: string } }) {
  const router = useRouter();
  const [course, setCourse] = useState<Course | null>(null);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState('');

  useEffect(() => {
    const fetchCourse = async () => {
      try {
        const token = localStorage.getItem('token');
        if (!token) {
          router.push('/auth/login');
          return;
        }

        const response = await fetch(`http://localhost:8080/api/v1/courses/${params.courseId}`, {
          headers: {
            'Authorization': `Bearer ${token}`,
          },
        });

        if (!response.ok) {
          if (response.status === 401) {
            router.push('/auth/login');
            return;
          }
          throw new Error('Erreur lors de la récupération du cours');
        }

        const data = await response.json();
        setCourse(data.course);
      } catch (err: any) {
        setError(err.message);
      } finally {
        setLoading(false);
      }
    };

    fetchCourse();
  }, [params.courseId, router]);

  if (loading) {
    return (
      <div className="min-h-screen flex items-center justify-center">
        <div className="animate-spin rounded-full h-12 w-12 border-b-2 border-indigo-600"></div>
      </div>
    );
  }

  if (error || !course) {
    return (
      <div className="min-h-screen flex items-center justify-center">
        <div className="text-center">
          <h2 className="text-2xl font-bold text-gray-900 mb-4">
            {error || 'Cours non trouvé'}
          </h2>
          <button
            onClick={() => router.push('/cours')}
            className="text-indigo-600 hover:text-indigo-500"
          >
            Retour aux cours
          </button>
        </div>
      </div>
    );
  }

  return (
    <div className="min-h-screen bg-gray-100">
      {/* En-tête du cours */}
      <div className="bg-white shadow">
        <div className="max-w-7xl mx-auto py-16 px-4 sm:px-6 lg:px-8">
          <div className="lg:grid lg:grid-cols-2 lg:gap-8">
            <div>
              <h1 className="text-4xl font-extrabold text-gray-900">
                {course.title}
              </h1>
              <p className="mt-4 text-lg text-gray-500">
                {course.description}
              </p>
            </div>
            <div className="mt-8 lg:mt-0">
              <img
                className="rounded-lg shadow-lg object-cover"
                src={course.image_url || '/placeholder-course.jpg'}
                alt={course.title}
              />
            </div>
          </div>
        </div>
      </div>

      {/* Liste des chapitres */}
      <div className="max-w-7xl mx-auto py-12 px-4 sm:px-6 lg:px-8">
        <h2 className="text-3xl font-extrabold text-gray-900 mb-8">
          Chapitres du cours
        </h2>
        <div className="space-y-4">
          {course.chapters?.map((chapter) => (
            <div
              key={chapter.id}
              className="bg-white rounded-lg shadow p-6 hover:shadow-lg transition-shadow duration-300 cursor-pointer"
              onClick={() => router.push(`/cours/${course.id}/chapitres/${chapter.id}`)}
            >
              <div className="flex items-center justify-between">
                <div>
                  <h3 className="text-xl font-semibold text-gray-900">
                    {chapter.title}
                  </h3>
                  <p className="mt-2 text-gray-500">{chapter.description}</p>
                </div>
                <div className="text-indigo-600">
                  <svg
                    className="h-6 w-6"
                    fill="none"
                    viewBox="0 0 24 24"
                    stroke="currentColor"
                  >
                    <path
                      strokeLinecap="round"
                      strokeLinejoin="round"
                      strokeWidth={2}
                      d="M9 5l7 7-7 7"
                    />
                  </svg>
                </div>
              </div>
            </div>
          ))}

          {(!course.chapters || course.chapters.length === 0) && (
            <div className="text-center text-gray-500">
              Aucun chapitre disponible pour le moment.
            </div>
          )}
        </div>
      </div>
    </div>
  );
} 