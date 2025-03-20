'use client';

import { useState, useEffect } from 'react';
import { useRouter } from 'next/navigation';

interface Course {
  id: number;
  title: string;
  description: string;
  image_url: string;
}

export default function CoursesPage() {
  const router = useRouter();
  const [courses, setCourses] = useState<Course[]>([]);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState('');

  useEffect(() => {
    const fetchCourses = async () => {
      try {
        const token = localStorage.getItem('token');
        if (!token) {
          router.push('/auth/login');
          return;
        }

        const response = await fetch('http://localhost:8080/api/v1/courses', {
          headers: {
            'Authorization': `Bearer ${token}`,
          },
        });

        if (!response.ok) {
          if (response.status === 401) {
            router.push('/auth/login');
            return;
          }
          throw new Error('Erreur lors de la récupération des cours');
        }

        const data = await response.json();
        setCourses(data.courses || []);
      } catch (err: any) {
        setError(err.message);
      } finally {
        setLoading(false);
      }
    };

    fetchCourses();
  }, [router]);

  if (loading) {
    return (
      <div className="min-h-screen flex items-center justify-center">
        <div className="animate-spin rounded-full h-12 w-12 border-b-2 border-indigo-600"></div>
      </div>
    );
  }

  return (
    <div className="min-h-screen bg-gray-100 py-12 px-4 sm:px-6 lg:px-8">
      <div className="max-w-7xl mx-auto">
        <div className="text-center">
          <h2 className="text-3xl font-extrabold text-gray-900 sm:text-4xl">
            Nos Cours
          </h2>
          <p className="mt-3 max-w-2xl mx-auto text-xl text-gray-500 sm:mt-4">
            Découvrez notre sélection de cours pour développer vos compétences
          </p>
        </div>

        {error && (
          <div className="mt-8 rounded-md bg-red-50 p-4">
            <div className="text-sm text-red-700">{error}</div>
          </div>
        )}

        <div className="mt-12 grid gap-8 sm:grid-cols-2 lg:grid-cols-3">
          {courses.map((course) => (
            <div
              key={course.id}
              className="flex flex-col rounded-lg shadow-lg overflow-hidden bg-white hover:shadow-xl transition-shadow duration-300 cursor-pointer"
              onClick={() => router.push(`/cours/${course.id}`)}
            >
              <div className="flex-shrink-0">
                <img
                  className="h-48 w-full object-cover"
                  src={course.image_url || '/placeholder-course.jpg'}
                  alt={course.title}
                />
              </div>
              <div className="flex-1 p-6 flex flex-col justify-between">
                <div className="flex-1">
                  <h3 className="text-xl font-semibold text-gray-900">
                    {course.title}
                  </h3>
                  <p className="mt-3 text-base text-gray-500">
                    {course.description}
                  </p>
                </div>
                <div className="mt-6">
                  <div className="text-base font-semibold text-indigo-600 hover:text-indigo-500">
                    Voir le cours →
                  </div>
                </div>
              </div>
            </div>
          ))}

          {courses.length === 0 && !error && (
            <div className="col-span-full text-center text-gray-500">
              Aucun cours disponible pour le moment.
            </div>
          )}
        </div>
      </div>
    </div>
  );
} 