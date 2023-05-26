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
import type { GetServerSidePropsContext, InferGetServerSidePropsType } from 'next';
import Link from 'next/link';
import { useRouter } from 'next/router';
import { getServerSession } from 'next-auth';
import { getCsrfToken, signIn } from 'next-auth/react';
import React, { useCallback, useState } from 'react';
import { useForm } from 'react-hook-form';

import { authOptions } from 'src/pages/api/auth/[...nextauth].api';

type FormValues = {
  email: string;
  password: string;
};

type SignInResponse = {
  error?: string;
  ok: boolean;
  status: number;
  url: null | string;
};

export default function Login({
  csrfToken,
}: InferGetServerSidePropsType<typeof getServerSideProps>) {
  const [response, setResponse] = useState<SignInResponse | null>(null);
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
        const response = await signIn('credentials', {
          ...values,
          redirect: false,
        });
        if (response?.ok && !response?.error) {
          // TODO: validate callback url is on the same domain
          router.push((router.query?.callbackUrl as string) ?? '/dashboard');
        }

        response && setResponse(response);
      } catch (error) {
        console.error(error);
      }
    },
    [router]
  );

  const onError = console.error;

  return (
    <VStack h="full" align="center" justify="center">
      <Heading>Login</Heading>
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
          <Button
            as={Link}
            variant="ghost"
            size="sm"
            w="full"
            type="submit"
            colorScheme="blue"
            href="/auth/register"
          >
            Don't have an account? Sign up
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
