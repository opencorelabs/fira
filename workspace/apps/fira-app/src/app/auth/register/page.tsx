import { redirect } from 'next/navigation';
import { getServerSession } from 'next-auth/next';
import { getCsrfToken } from 'next-auth/react';

import { authOptions } from 'src/pages/api/auth/[...nextauth]';

import { Register } from './Register';

export default async function RegisterPage() {
  const session = await getServerSession(authOptions);
  if (session) {
    redirect('/networth');
  }
  const csrfToken = await getCsrfToken();
  return <Register csrfToken={csrfToken} />;
}
