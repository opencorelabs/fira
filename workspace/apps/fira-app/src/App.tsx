import { NextPage } from 'next';
import { SessionProvider } from 'next-auth/react';

import { AppProviders } from 'src/AppProviders';

import { GlobalStyle } from './theme/GlobalStyle';

export function App({ Component, pageProps, session }) {
  const getLayout = Component.getLayout || ((page: NextPage) => page);
  return (
    <AppProviders>
      <SessionProvider session={session}>
        <GlobalStyle />
        {getLayout(<Component {...pageProps} />)}
      </SessionProvider>
    </AppProviders>
  );
}
