import { Layout } from 'src/components/layout/Layout';

export default function Accounts() {
  return (
    <div>
      <h1>Accounts</h1>
    </div>
  );
}

Accounts.getLayout = function getLayout(page: React.ReactNode) {
  return <Layout>{page}</Layout>;
};
