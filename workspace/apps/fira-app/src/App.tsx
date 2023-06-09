import { ChakraProvider } from '@chakra-ui/react';

import { DashboardLayout } from 'src/components/layouts/dashboard/DashboardLayout';
import { LoggedOutLayout } from 'src/components/layouts/loggedout/Layout';
import { ModalProvider } from 'src/context/ModalContext';
import { theme } from 'src/theme';

import { GlobalStyle } from './theme/GlobalStyle';

function getLayout(Component) {
  if (Component.authenticated) {
    return DashboardLayout;
  }
  return LoggedOutLayout;
}

export function App({ Component, pageProps }) {
  const Layout = getLayout(Component);
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
