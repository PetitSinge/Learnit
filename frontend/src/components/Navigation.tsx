'use client';

import { Fragment } from 'react';
import { Disclosure, Menu, Transition } from '@headlessui/react';
import { Bars3Icon, XMarkIcon } from '@heroicons/react/24/outline';
import Link from 'next/link';
import { usePathname } from 'next/navigation';

const navigation = [
  { name: 'Accueil', href: '/' },
  { name: 'Cours', href: '/cours' },
  { name: 'Exercices', href: '/exercices' },
  { name: 'Quiz', href: '/quiz' },
  { name: 'Tests', href: '/tests' },
  { name: 'Classement', href: '/classement' },
];

function classNames(...classes: string[]) {
  return classes.filter(Boolean).join(' ');
}

const Navigation = () => {
  const pathname = usePathname();

  const isActive = (path: string) => pathname?.startsWith(path);

  const navItems = [
    { path: '/cours', label: 'Cours', icon: '📚' },
    { path: '/quiz', label: 'Quiz', icon: '✍️' },
    { path: '/exercices', label: 'Exercices', icon: '💻' },
    { path: '/dashboard', label: 'Tableau de bord', icon: '📊' },
    { path: '/classement', label: 'Classement', icon: '👑' },
  ];

  return (
    <nav className="bg-white shadow-lg">
      <div className="max-w-7xl mx-auto px-4">
        <div className="flex justify-between h-16">
          <div className="flex space-x-8">
            {navItems.map((item) => (
              <Link
                key={item.path}
                href={item.path}
                className={`inline-flex items-center px-3 py-2 text-sm font-medium ${
                  isActive(item.path)
                    ? 'text-blue-600 border-b-2 border-blue-600'
                    : 'text-gray-500 hover:text-blue-600 hover:border-b-2 hover:border-blue-300'
                }`}
              >
                <span className="mr-2">{item.icon}</span>
                {item.label}
              </Link>
            ))}
          </div>
        </div>
      </div>
    </nav>
  );
};

export default Navigation; 