import { Container, Flex } from '@chakra-ui/react';

import { WordMark } from '../WordMark';

type AuthLayoutProps = {
  children: React.ReactNode;
};

export function AuthLayout({ children }: AuthLayoutProps) {
  return (
    <Flex minH="100vh" direction="column">
      <header>
        <WordMark />
      </header>
      <Flex flex={1} align="center" justify="center">
        <Container maxW="container.xl">
          <Flex justify="center">{children}</Flex>
        </Container>
      </Flex>
      <footer>Footer</footer>
    </Flex>
  );
}
