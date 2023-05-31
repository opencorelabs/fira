import { Container, DarkMode, Flex } from '@chakra-ui/react';

import { Footer } from './Footer';
import { Header } from './Header';

type AuthLayoutProps = {
  children: React.ReactNode;
};

export function LoggedOutLayout({ children }: AuthLayoutProps) {
  return (
    <DarkMode>
      <Flex minH="100vh" direction="column" bg="gray.900">
        <Header />
        <Flex flex={1} align="center" justify="center">
          <Container maxW="container.xl">
            <Flex justify="center">{children}</Flex>
          </Container>
        </Flex>
        <Footer />
      </Flex>
    </DarkMode>
  );
}
