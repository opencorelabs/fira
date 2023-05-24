import {
  Box,
  Button,
  ButtonGroup,
  chakra,
  Heading,
  Icon,
  VStack,
} from '@chakra-ui/react';
import NextImage from 'next/image';
import Link from 'next/link';
import { FaDiscord, FaStar } from 'react-icons/fa';

const LandingImage = chakra(NextImage, {
  baseStyle: {
    borderRadius: 'lg',
  },
  shouldForwardProp: (prop) => ['width', 'height', 'src', 'alt'].includes(prop),
});

export default function Index() {
  return (
    <VStack gap={8} mb={10}>
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
          colorScheme="github"
          variant="outline"
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
        <LandingImage
          src="/images/fira-screenshot.png"
          alt="Fira Screenshot"
          width={700}
          height={555}
          w="auto"
          h="auto"
          objectFit="cover"
        />
      </Box>
      <ButtonGroup>
        <Button
          as={Link}
          href="/auth/login"
          colorScheme="primary"
          variant="outline"
          px={8}
        >
          Login
        </Button>
        <Button
          as={Link}
          href="/auth/register"
          colorScheme="primary"
          variant="solid"
          px={8}
        >
          Register
        </Button>
      </ButtonGroup>
    </VStack>
  );
}

Index.landing = true;
