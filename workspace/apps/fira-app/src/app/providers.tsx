'use client';

import { CacheProvider } from '@chakra-ui/next-js';
import { ChakraProvider, theme } from '@chakra-ui/react';
import { SessionProvider } from 'next-auth/react';

import { GlobalStyle } from 'src/theme/GlobalStyle';

export function Providers({ children, session }) {
  return (
    <CacheProvider>
      <ChakraProvider theme={theme}>
        <GlobalStyle />
        <SessionProvider session={session}>{children}</SessionProvider>
      </ChakraProvider>
    </CacheProvider>
  );
}
