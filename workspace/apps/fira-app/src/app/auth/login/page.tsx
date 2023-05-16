import { getServerSession } from 'next-auth/next';
import { getCsrfToken } from 'next-auth/react';

import { authOptions } from 'src/pages/api/auth/[...nextauth]';

import { Login } from './Login';

export default async function LoginPage() {
  const session = await getServerSession(authOptions);
  if (session) {
    // how to redirect to home page?
    return null;
  }
  const csrfToken = await getCsrfToken();
  return <Login csrfToken={csrfToken} />;
}
