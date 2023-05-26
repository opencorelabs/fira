import { Box, Container, Flex, Text } from '@chakra-ui/react';

import { WordMark } from 'src/components/WordMark';

export function Footer() {
  return (
    <Flex as="footer" p={6} bg="black">
      <Container maxW="container.xl">
        <Flex
          align="center"
          color="white"
          direction={{ base: 'column', md: 'row' }}
          justify="space-between"
        >
          <Box>
            <Text>
              <WordMark fill="white" width="75px" />
            </Text>
          </Box>
          <Box>
            <Text>&copy; 2023 - Fira</Text>
          </Box>
          <Box />
        </Flex>
      </Container>
    </Flex>
  );
}
