import { ChakraProvider } from '@chakra-ui/react';
import { SessionProvider } from 'next-auth/react';

import { AuthLayout } from './components/auth/Layout';
import { Layout as DashboardLayout } from './components/layout/Layout';
import { theme } from './theme';
import { GlobalStyle } from './theme/GlobalStyle';

export function App({ Component, pageProps }) {
  const Layout = Component.auth ? DashboardLayout : AuthLayout;
  return (
    <ChakraProvider theme={theme}>
      <SessionProvider session={pageProps.session}>
        <GlobalStyle />
        <Layout>
          <Component {...pageProps} />
        </Layout>
      </SessionProvider>
    </ChakraProvider>
  );
}
