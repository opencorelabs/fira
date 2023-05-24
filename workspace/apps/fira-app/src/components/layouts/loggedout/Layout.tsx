import { Container, Flex } from '@chakra-ui/react';

import { WordMark } from 'src/components/WordMark';

import { Footer } from '../Footer';

type AuthLayoutProps = {
  children: React.ReactNode;
};

export function LoggedOutLayout({ children }: AuthLayoutProps) {
  return (
    <Flex minH="100vh" direction="column">
      <header>
        <WordMark fill="black" />
      </header>
      <Flex flex={1} align="center" justify="center">
        <Container maxW="container.xl">
          <Flex justify="center">{children}</Flex>
        </Container>
      </Flex>
      <Footer />
    </Flex>
  );
}
