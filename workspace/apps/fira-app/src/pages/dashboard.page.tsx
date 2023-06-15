import { NetworthCard } from 'src/components/NetworthCard/NetworthCard';
import { withSessionSsr } from 'src/lib/session/session';

export default function Dashboard() {
  return <NetworthCard />;
}

Dashboard.authenticated = true;

export const getServerSideProps = withSessionSsr(async function getServerSideProps() {
  return {
    props: {},
  };
});
