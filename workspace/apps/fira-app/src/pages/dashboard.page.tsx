import type { GetServerSidePropsContext } from 'next';

import { NetworthCard } from 'src/components/NetworthCard/NetworthCard';
import { withSessionSsr } from 'src/lib/session/session';

export default function Dashboard() {
  return <NetworthCard />;
}

Dashboard.auth = true;

export const getServerSideProps = withSessionSsr(async function getServerSideProps(
  context: GetServerSidePropsContext
) {
  // if (!context.req.session?.user?.verified) {
  //   return {
  //     redirect: {
  //       destination: '/auth/verify-email',
  //       permanent: false,
  //     },
  //   };
  // }
  console.info('dashboard context.req.session', context.req.session);
  return {
    props: {},
  };
});
