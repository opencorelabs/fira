import { Container, Flex } from '@chakra-ui/react';

import { Footer } from './Footer';
import { Header } from './Header';

type LayoutProps = {
  children: React.ReactNode;
};

export function Layout({ children }: LayoutProps) {
  return (
    <Flex minH="100vh" direction="column" bg="black">
      <Header />
      <Flex as="main" flex={1} align="center" justify="center">
        <Container maxW="container.xl">{children}</Container>
      </Flex>
      <Footer />
    </Flex>
  );
}
