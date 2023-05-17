'use client';

import { CacheProvider } from '@chakra-ui/next-js';
import { ChakraProvider } from '@chakra-ui/react';
import { SessionProvider } from 'next-auth/react';

import { theme } from 'src/theme';
import { GlobalStyle } from 'src/theme/GlobalStyle';

export function Providers({ children, session }) {
  return (
    <SessionProvider session={session}>
      <CacheProvider>
        <ChakraProvider theme={theme}>
          <GlobalStyle />
          {children}
        </ChakraProvider>
      </CacheProvider>
    </SessionProvider>
  );
}
