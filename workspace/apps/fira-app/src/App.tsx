import { ChakraProvider } from '@chakra-ui/react';

import { AuthLayout } from './components/layout/AuthLayout';
import { Layout as DashboardLayout } from './components/layout/Layout';
import { ModalProvider } from './context/ModalContext';
import { theme } from './theme';
import { GlobalStyle } from './theme/GlobalStyle';

export function App({ Component, pageProps }) {
  const Layout = Component.auth ? DashboardLayout : AuthLayout;
  return (
    <ChakraProvider theme={theme}>
      <GlobalStyle />
      <ModalProvider>
        <Layout>
          <Component {...pageProps} />
        </Layout>
      </ModalProvider>
    </ChakraProvider>
  );
}
