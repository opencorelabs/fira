import { Layout } from 'src/components/layout/Layout';

export default function SettingsAccount() {
  return (
    <div>
      <h1>Settings Account</h1>
    </div>
  );
}

SettingsAccount.getLayout = function getLayout(page: React.ReactNode) {
  return <Layout>{page}</Layout>;
};
