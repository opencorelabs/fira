'use client';

import { Box, Button, Container, Input, Text } from '@chakra-ui/react';
import { V1AccountNamespace } from '@fira/api-sdk';
import { useRouter } from 'next/navigation';
import { useCallback } from 'react';
import { useForm } from 'react-hook-form';

import { api } from 'src/lib/fira-api';

type FormValues = {
  token: string;
};

export function VerifyEmail() {
  const router = useRouter();
  const { register, handleSubmit } = useForm<FormValues>();

  const handleVerify = useCallback(
    async (data: FormValues) => {
      const response = await api.firaServiceVerifyAccount({
        token: data.token,
        // @ts-expect-error bad types in api-sdk
        type: 1,
        namespace: V1AccountNamespace.ACCOUNT_NAMESPACE_CONSUMER,
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
