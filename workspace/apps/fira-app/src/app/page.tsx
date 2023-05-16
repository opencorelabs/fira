'use client';
import { Heading } from '@chakra-ui/react';

export default function IndexPage() {
  return (
    <div>
      <Heading>Logged Out</Heading>
      <pre>{JSON.stringify({}, null, 2)}</pre>
    </div>
  );
}
