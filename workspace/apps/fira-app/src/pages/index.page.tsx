import { Layout } from 'src/components/layout/Layout';

export default function Index() {
  return (
    <div>
      <h1>Home</h1>
    </div>
  );
}

Index.getLayout = function getLayout(page: React.ReactNode) {
  return <Layout>{page}</Layout>;
};
