import { Box, Button, Container, Input, Text } from '@chakra-ui/react';
import { useRouter } from 'next/router';
import { useCallback } from 'react';
import { useForm } from 'react-hook-form';

import { api } from 'src/lib/fira-api';

type FormValues = {
  token: string;
};

export default function VerifyEmail() {
  const router = useRouter();
  const { register, handleSubmit } = useForm<FormValues>();

  const handleVerify = useCallback(
    async (data: FormValues) => {
      const response = await api.firaServiceVerifyAccount({
        // @ts-expect-error type is required
        type: 1,
        token: data.token,
      });
      console.info('response', response);
      router.push('/');
    },
    [router]
  );

  return (
    <Container maxW="container.xl">
      <Text>
        This is a placeholder page. Copy your verification token from the console and
        enter here to verify your new account
      </Text>
      <Box as="form" onSubmit={handleSubmit(handleVerify)} mt={2}>
        <Input
          placeholder="verification token"
          {...register('token', {
            required: 'Token is required',
          })}
        />
        <Button type="submit" mt={2}>
          Verify Email
        </Button>
      </Box>
    </Container>
  );
}
