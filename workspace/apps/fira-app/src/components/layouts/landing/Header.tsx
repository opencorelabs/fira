import { Box, Container, Flex, HStack, IconButton } from '@chakra-ui/react';
import Link from 'next/link';
import { FaDiscord, FaGithub } from 'react-icons/fa';

import { WordMark } from 'src/components/WordMark';

export function LandingHeader() {
  return (
    <header>
      <Container maxW="container.xl" py={6}>
        <Flex justify="space-between" align="center">
          <Box>
            <Link href="/">
              <WordMark fill="white" />
            </Link>
          </Box>
          <HStack gap={0}>
            <IconButton
              as={Link}
              href="https://discord.gg/uGFwMGDGku"
              target="_blank"
              colorScheme="transparent"
              color="white"
              variant="outline"
              size="sm"
              aria-label="Discord"
              icon={<FaDiscord />}
              _hover={{ transform: 'scale(1.05)' }}
            />
            <IconButton
              as={Link}
              href="https://github.com/opencorelabs/fira"
              target="_blank"
              colorScheme="transparent"
              variant="outline"
              color="white"
              size="sm"
              aria-label="Github"
              icon={<FaGithub />}
              _hover={{ transform: 'scale(1.05)' }}
            />
          </HStack>
        </Flex>
      </Container>
    </header>
  );
}
