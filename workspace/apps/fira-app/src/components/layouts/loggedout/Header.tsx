import { Container } from '@chakra-ui/react';
import Link from 'next/link';

import { WordMark } from 'src/components/WordMark';

export function Header() {
  return (
    <header>
      <Container maxW="container.xl" py={4}>
      <Link href="/">
        <WordMark />
      </Link>
      </Container>
    </header>
  );
}
