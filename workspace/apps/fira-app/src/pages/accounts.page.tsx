import { getServerSession } from 'next-auth';
import { useSession } from 'next-auth/react';

import { authOptions } from './api/auth/[...nextauth].api';

export default function Accounts() {
  console.info('useSession()', useSession());
  return (
    <div>
      <h1>Accounts</h1>
    </div>
  );
}

Accounts.auth = true;

export const getServerSideProps = async (context) => {
  const session = await getServerSession(context.req, context.res, authOptions);
  console.info('session', session);
  return {
    props: {},
  };
};
