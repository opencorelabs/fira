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
import {
  V1AccountCredentialType,
  V1AccountNamespace,
  V1CreateAccountResponse,
} from '@fira/api-sdk';
import type { GetServerSidePropsContext, InferGetServerSidePropsType } from 'next';
import { useRouter } from 'next/router';
import { getServerSession } from 'next-auth/next';
import { getCsrfToken } from 'next-auth/react';
import React, { useCallback, useState } from 'react';
import { useForm } from 'react-hook-form';

import { getApi } from 'src/lib/fira-api';

import { authOptions } from '../api/auth/[...nextauth].api';

type FormValues = {
  name: string;
  email: string;
  password: string;
};

export default function Register({
  csrfToken,
}: InferGetServerSidePropsType<typeof getServerSideProps>) {
  const [response, setResponse] = useState<V1CreateAccountResponse | null>(null);
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
        const response = await getApi().firaServiceCreateAccount({
          namespace: V1AccountNamespace.ACCOUNT_NAMESPACE_CONSUMER,
          credential: {
            credentialType: V1AccountCredentialType.ACCOUNT_CREDENTIAL_TYPE_EMAIL,
            emailCredential: {
              email: values.email,
              password: values.password,
              verificationBaseUrl: process.env.NEXT_PUBLIC_VERIFICATION_BASE_URL,
              name: '',
            },
          },
        });
        if (!response.ok && response.error) {
          throw response.error;
        }

        setResponse(response.data);
        console.info('response', response);

        router.push('/auth/verify-email');
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
          <FormControl isInvalid={Boolean(errors.name)}>
            <Input
              {...register('name', { required: 'Name is required' })}
              placeholder="Name"
              type="text"
            />
            <FormErrorMessage>{errors.name?.message}</FormErrorMessage>
          </FormControl>
          <FormControl isInvalid={Boolean(errors.email)}>
            <Input
              {...register('email', { required: 'Email is required' })}
              placeholder="Email"
              type="email"
            />
            <FormErrorMessage>{errors.email?.message}</FormErrorMessage>
          </FormControl>
          <FormControl isInvalid={Boolean(errors.password)}>
            <Input
              {...register('password', {
                required: 'Password is required',
                validate: {
                  minLength: (value) =>
                    value.length >= 10 || 'Password must be at least 10 characters',
                },
              })}
              placeholder="Password"
              type="password"
            />
            <FormErrorMessage>{errors.password?.message}</FormErrorMessage>
          </FormControl>
          <Button size="sm" w="full" type="submit" colorScheme="blue">
            Sign up
          </Button>
          {!!response?.errorMessage && (
            <Text color="red">
              {response.errorMessage} - Status {response.status}
            </Text>
          )}
        </VStack>
      </Box>
    </VStack>
  );
}

export async function getServerSideProps(context: GetServerSidePropsContext) {
  const session = await getServerSession(context.req, context.res, authOptions);
  if (session) {
    return {
      redirect: {
        destination: '/dashboard',
        permanent: false,
      },
    };
  }
  const csrfToken = await getCsrfToken(context);
  return {
    props: { csrfToken },
  };
}
