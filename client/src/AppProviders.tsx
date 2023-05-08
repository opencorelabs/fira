import { ChakraProvider, theme } from '@chakra-ui/react';

export function AppProviders({ children }) {
  return <ChakraProvider theme={theme}>{children}</ChakraProvider>;
}
