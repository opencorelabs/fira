import { Box, Button, ButtonGroup, Heading, Icon, VStack } from '@chakra-ui/react';
import Image from 'next/image';
import Link from 'next/link';
import { FaDiscord, FaStar } from 'react-icons/fa';

export default function Index() {
  return (
    <VStack gap={8}>
      <Box textAlign="center">
        <Heading size="2xl" color="white" maxW={768}>
          The completely open personal financial management tool
        </Heading>
        <Heading size="lg" color="gray.500" maxW={986} mt={4}>
          Finally, an open source way to manage your money
        </Heading>
      </Box>
      <ButtonGroup>
        <Button
          as={Link}
          href="https://github.com/opencorelabs/fira"
          rightIcon={<Icon as={FaStar} color="yellow.500" />}
        >
          Star us on Github
        </Button>
        <Button
          as={Link}
          href="https://discord.gg/uGFwMGDGku"
          rightIcon={<FaDiscord />}
          colorScheme="discord"
        >
          Join Discord
        </Button>
      </ButtonGroup>
      <Box>
        <Image
          src="/images/fira-screenshot.png"
          alt="Fira Screenshot"
          width={700}
          height={555}
        />
      </Box>
    </VStack>
  );
}

Index.landing = true;
