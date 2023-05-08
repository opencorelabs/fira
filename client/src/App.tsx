import { AppProviders } from 'src/AppProviders';
import { Layout } from 'src/components/layout/Layout';

import { GlobalStyle } from './theme/GlobalStyle';

export function App({ Component, pageProps }) {
  return (
    <AppProviders>
      <GlobalStyle />
      <Layout>
        <Component {...pageProps} />
      </Layout>
    </AppProviders>
  );
}
