import { redirect } from 'next/navigation';
import { getServerSession } from 'next-auth/next';

import { authOptions } from 'src/pages/api/auth/[...nextauth]';

import { VerifyEmail } from './VerifyEmail';

export default async function VerifyEmailPage() {
  const session = await getServerSession(authOptions);
  if (session) {
    redirect('/networth');
  }
  return <VerifyEmail />;
}
