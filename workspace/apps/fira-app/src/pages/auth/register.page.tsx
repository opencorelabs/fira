import {
  Box,
  Button,
  FormControl,
  FormErrorMessage,
  Heading,
  Input,
  Text,
  VStack,
} from '@chakra-ui/react';
import { V1AccountCredentialType, V1AccountNamespace } from '@fira/api-sdk';
import type { GetServerSidePropsContext, InferGetServerSidePropsType } from 'next';
import { useRouter } from 'next/router';
import { getCsrfToken, signIn } from 'next-auth/react';
import React, { useCallback, useState } from 'react';
import { useForm } from 'react-hook-form';

import { api } from 'src/lib/fira-api';

type FormValues = {
  email: string;
  password: string;
};

type RegisterResponse = {
  error: string | null;
  ok: boolean;
  status: number;
  url: null | string;
};

export default function Register({
  csrfToken,
}: InferGetServerSidePropsType<typeof getServerSideProps>) {
  const [response, setResponse] = useState<RegisterResponse>(null);
  const router = useRouter();
  const {
    handleSubmit,
    register,
    formState: { errors },
  } = useForm<FormValues>();

  const onSubmit = useCallback(
    async (values: FormValues) => {
      try {
        setResponse(null);

        const response = await api.firaServiceCreateAccount({
          namespace: V1AccountNamespace.ACCOUNT_NAMESPACE_CONSUMER,
          credential: {
            credentialType: V1AccountCredentialType.ACCOUNT_CREDENTIAL_TYPE_EMAIL,
            emailCredential: {
              email: values.email,
              password: values.password,
            },
          },
        });
        console.log('response', response);

        if (!response.ok && response.error) {
          throw response.error;
        }

        // Signin to create the session
        const signinRepsonse = await signIn('credentials', {
          ...values,
          redirect: false,
        });
        console.log('signinRepsonse', signinRepsonse);
        if (!signinRepsonse.ok && signinRepsonse.error) {
          throw response.error;
        }
        router.push('/');
      } catch (error) {
        console.error(error);
      }
    },
    [router]
  );

  const onError = console.error;

  return (
    <VStack h="full" align="center" justify="center">
      <Heading>Register</Heading>
      <Box w="24rem">
        <VStack as="form" onSubmit={handleSubmit(onSubmit, onError)}>
          <input name="csrfToken" type="hidden" defaultValue={csrfToken} />
          <FormControl isInvalid={Boolean(errors.email)}>
            <Input
              {...register('email', { required: 'required' })}
              placeholder="Email"
              type="email"
            />
            <FormErrorMessage>{errors.email?.message}</FormErrorMessage>
          </FormControl>
          <FormControl isInvalid={Boolean(errors.password)}>
            <Input
              {...register('password', { required: 'required' })}
              placeholder="Password"
              type="password"
            />
            <FormErrorMessage>{errors.password?.message}</FormErrorMessage>
          </FormControl>
          <Button size="sm" w="full" type="submit" colorScheme="blue">
            Login
          </Button>
          {!!response?.error && (
            <Text color="red">
              {response.error} - Status {response.status}
            </Text>
          )}
        </VStack>
      </Box>
    </VStack>
  );
}

export async function getServerSideProps(context: GetServerSidePropsContext) {
  const csrfToken = await getCsrfToken(context);
  return {
    props: { csrfToken },
  };
}
