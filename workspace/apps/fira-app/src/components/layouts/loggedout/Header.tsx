import { Container } from '@chakra-ui/react';

import { WordMark } from 'src/components/WordMark';

export function Header() {
  return (
    <header>
      <Container maxW="container.xl" py={4}>
        <WordMark />
      </Container>
    </header>
  );
}
