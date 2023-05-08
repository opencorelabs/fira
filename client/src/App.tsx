import { AppProviders } from 'src/AppProviders';
import { Layout } from 'src/components/layout/Layout';

export function App({ Component, pageProps }) {
  return (
    <AppProviders>
      <Layout>
        <Component {...pageProps} />
      </Layout>
    </AppProviders>
  );
}
