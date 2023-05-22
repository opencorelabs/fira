import { GetServerSidePropsContext } from 'next';

import { withSessionSsr } from 'src/lib/session';

export default function Accounts() {
  return (
    <div>
      <h1>Accounts</h1>
    </div>
  );
}

Accounts.auth = true;
export const getServerSideProps = withSessionSsr(async function getServerSideProps(
  context: GetServerSidePropsContext
) {
  console.info('context', context.req.session);
  return {
    props: {},
  };
});
