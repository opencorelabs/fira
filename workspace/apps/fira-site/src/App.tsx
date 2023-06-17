import { ChakraProvider } from '@chakra-ui/react';

import { Layout } from 'src/components/layout/Layout';
import { theme } from 'src/theme';
import { GlobalStyle } from 'src/theme/GlobalStyle';

export function App({ Component, pageProps }) {
  return (
    <ChakraProvider theme={theme}>
      <GlobalStyle />
      <Layout>
        <Component {...pageProps} />
      </Layout>
    </ChakraProvider>
  );
}
