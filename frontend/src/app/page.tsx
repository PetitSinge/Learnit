'use client';

import { useRouter } from 'next/navigation';
import { useEffect, useState } from 'react';
import Link from 'next/link';
import { 
  AcademicCapIcon, 
  BeakerIcon, 
  ClipboardDocumentCheckIcon,
  TrophyIcon,
  ChartBarIcon,
  UserGroupIcon
} from '@heroicons/react/24/outline';

const features = [
  {
    name: 'Cours',
    description: 'Accédez à des cours structurés par chapitres avec des ressources PDF.',
    href: '/cours',
    icon: AcademicCapIcon,
  },
  {
    name: 'Exercices Pratiques',
    description: 'Pratiquez dans un environnement Docker Linux avec des exercices préenregistrés.',
    href: '/exercices',
    icon: BeakerIcon,
  },
  {
    name: 'Quiz',
    description: 'Testez vos connaissances avec des QCM et obtenez des explications détaillées.',
    href: '/quiz',
    icon: ClipboardDocumentCheckIcon,
  },
  {
    name: 'Tests Réels',
    description: 'Préparez-vous aux examens avec des simulations de tests réels.',
    href: '/tests',
    icon: TrophyIcon,
  },
  {
    name: 'Tableau de Bord',
    description: 'Suivez votre progression et analysez vos performances.',
    href: '/dashboard',
    icon: ChartBarIcon,
  },
  {
    name: 'Classement',
    description: 'Comparez vos résultats et motivez-vous à progresser.',
    href: '/classement',
    icon: UserGroupIcon,
  },
];

export default function Home() {
  const router = useRouter();
  const [isAuthenticated, setIsAuthenticated] = useState(false);

  useEffect(() => {
    const token = localStorage.getItem('token');
    setIsAuthenticated(!!token);
  }, []);

  return (
    <div className="bg-white">
      {/* Hero section */}
      <div className="relative isolate overflow-hidden bg-gradient-to-b from-indigo-100/20">
        <div className="mx-auto max-w-7xl pb-24 pt-10 sm:pb-32 lg:grid lg:grid-cols-2 lg:gap-x-8 lg:px-8 lg:py-40">
          <div className="px-6 lg:px-0 lg:pt-4">
            <div className="mx-auto max-w-2xl">
              <div className="max-w-lg">
                <h1 className="mt-10 text-4xl font-bold tracking-tight text-gray-900 sm:text-6xl">
                  Apprenez à votre rythme
                </h1>
                <p className="mt-6 text-lg leading-8 text-gray-600">
                  Découvrez notre plateforme d'apprentissage en ligne avec des cours interactifs,
                  des exercices pratiques et un suivi personnalisé de votre progression.
                </p>
                <div className="mt-10 flex items-center gap-x-6">
                  {isAuthenticated ? (
                    <button
                      onClick={() => router.push('/cours')}
                      className="rounded-md bg-indigo-600 px-3.5 py-2.5 text-sm font-semibold text-white shadow-sm hover:bg-indigo-500 focus-visible:outline focus-visible:outline-2 focus-visible:outline-offset-2 focus-visible:outline-indigo-600"
                    >
                      Voir les cours
                    </button>
                  ) : (
                    <>
                      <button
                        onClick={() => router.push('/auth/login')}
                        className="rounded-md bg-indigo-600 px-3.5 py-2.5 text-sm font-semibold text-white shadow-sm hover:bg-indigo-500 focus-visible:outline focus-visible:outline-2 focus-visible:outline-offset-2 focus-visible:outline-indigo-600"
                      >
                        Commencer
                      </button>
                      <button
                        onClick={() => router.push('/auth/register')}
                        className="text-sm font-semibold leading-6 text-gray-900"
                      >
                        Créer un compte <span aria-hidden="true">→</span>
                      </button>
                    </>
                  )}
                </div>
              </div>
            </div>
          </div>
          <div className="mt-20 sm:mt-24 md:mx-auto md:max-w-2xl lg:mx-0 lg:mt-0 lg:w-screen">
            <div className="absolute inset-y-0 right-1/2 -z-10 -mr-10 w-[200%] skew-x-[-30deg] bg-white shadow-xl shadow-indigo-600/10 ring-1 ring-indigo-50 md:-mr-20 lg:-mr-36" />
            <div className="shadow-lg md:rounded-3xl">
              <div className="bg-indigo-500 [clip-path:inset(0)] md:[clip-path:inset(0_round_theme(borderRadius.3xl))]">
                <div className="absolute -inset-y-px left-1/2 -z-10 ml-10 w-[200%] skew-x-[-30deg] bg-indigo-100 opacity-20 ring-1 ring-inset ring-white md:ml-20 lg:ml-36" />
                <div className="relative px-6 pt-8 sm:pt-16 md:pl-16 md:pr-0">
                  <div className="mx-auto max-w-2xl md:mx-0 md:max-w-none">
                    <div className="w-screen overflow-hidden rounded-tl-xl bg-gray-900">
                      <div className="flex bg-gray-800/40 ring-1 ring-white/5">
                        <div className="-mb-px flex text-sm font-medium leading-6 text-gray-400">
                          <div className="border-b border-r border-b-white/20 border-r-white/10 bg-white/5 px-4 py-2 text-white">
                            LearnIT
                          </div>
                          <div className="border-r border-gray-600/10 px-4 py-2">Cours</div>
                        </div>
                      </div>
                      <div className="px-6 pt-6 pb-14">
                        {/* Placeholder pour une image ou une animation */}
                        <div className="text-white">
                          Bienvenue sur LearnIT
                        </div>
                      </div>
                    </div>
                  </div>
                </div>
              </div>
            </div>
          </div>
        </div>
      </div>

      {/* Features section */}
      <div className="mx-auto mt-32 max-w-7xl px-6 sm:mt-56 lg:px-8">
        <div className="mx-auto max-w-2xl lg:text-center">
          <h2 className="text-base font-semibold leading-7 text-indigo-600">
            Apprenez plus rapidement
          </h2>
          <p className="mt-2 text-3xl font-bold tracking-tight text-gray-900 sm:text-4xl">
            Tout ce dont vous avez besoin pour progresser
          </p>
          <p className="mt-6 text-lg leading-8 text-gray-600">
            Notre plateforme offre une expérience d'apprentissage complète et interactive,
            conçue pour vous aider à atteindre vos objectifs.
          </p>
        </div>
        <div className="mx-auto mt-16 max-w-2xl sm:mt-20 lg:mt-24 lg:max-w-none">
          <dl className="grid max-w-xl grid-cols-1 gap-x-8 gap-y-16 lg:max-w-none lg:grid-cols-3">
            <div className="flex flex-col">
              <dt className="text-base font-semibold leading-7 text-gray-900">
                Cours interactifs
              </dt>
              <dd className="mt-2 flex flex-auto flex-col text-base leading-7 text-gray-600">
                <p className="flex-auto">
                  Des cours structurés et interactifs pour un apprentissage efficace.
                </p>
              </dd>
            </div>
            <div className="flex flex-col">
              <dt className="text-base font-semibold leading-7 text-gray-900">
                Exercices pratiques
              </dt>
              <dd className="mt-2 flex flex-auto flex-col text-base leading-7 text-gray-600">
                <p className="flex-auto">
                  Mettez en pratique vos connaissances avec des exercices concrets.
                </p>
              </dd>
            </div>
            <div className="flex flex-col">
              <dt className="text-base font-semibold leading-7 text-gray-900">
                Suivi personnalisé
              </dt>
              <dd className="mt-2 flex flex-auto flex-col text-base leading-7 text-gray-600">
                <p className="flex-auto">
                  Suivez votre progression et recevez des recommandations personnalisées.
                </p>
              </dd>
            </div>
          </dl>
        </div>
      </div>
    </div>
  );
}
