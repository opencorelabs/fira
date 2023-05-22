import { Box, Button } from '@chakra-ui/react';
import Link from 'next/link';

export default function Index() {
  return (
    <Box>
      <Button as={Link} href="/auth/login">
        Login
      </Button>
      <Button as={Link} href="/auth/register">
        Signup
      </Button>
    </Box>
  );
}

Index.auth = false;
