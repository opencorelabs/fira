import { Box, Button, ButtonGroup, Container, Flex, IconButton } from '@chakra-ui/react';
import Link from 'next/link';
import { FaDiscord, FaGithub } from 'react-icons/fa';

import { WordMark } from 'src/components/WordMark';

export function Header() {
  return (
    <header>
      <Container maxW="container.xl" py={6}>
        <Flex justify="space-between" align="center">
          <Box>
            <Link href="/">
              <WordMark fill="white" />
            </Link>
          </Box>
          <ButtonGroup>
            {/* TODO: Enable this with build/env var */}
            {process.env.NODE_ENV === 'development' && (
              <>
                <Button
                  as={Link}
                  href="/auth/login"
                  colorScheme="primary"
                  variant="outline"
                  px={8}
                  size="sm"
                >
                  Login
                </Button>
                <Button
                  as={Link}
                  href="/auth/register"
                  colorScheme="primary"
                  variant="outline"
                  px={8}
                  size="sm"
                >
                  Register
                </Button>
              </>
            )}
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
          </ButtonGroup>
        </Flex>
      </Container>
    </header>
  );
}
