import { Container, Flex } from '@chakra-ui/react';

import { Footer } from '../Footer';
import { LandingHeader } from './Header';

type LandingLayoutProps = {
  children: React.ReactNode;
};

export function LandingLayout({ children }: LandingLayoutProps) {
  return (
    <Flex minH="100vh" direction="column" bg="black">
      <LandingHeader />
      <Flex flex={1} align="center" justify="center">
        <Container maxW="container.xl">
          <Flex justify="center">{children}</Flex>
        </Container>
      </Flex>
      <Footer />
    </Flex>
  );
}
