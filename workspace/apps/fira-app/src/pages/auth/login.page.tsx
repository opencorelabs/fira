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
import React, { useCallback, useState } from 'react';
import { useForm } from 'react-hook-form';

import { PAGE_ROUTES } from 'src/config/routes';
import { login } from 'src/lib/auth';
import { withSessionSsr } from 'src/lib/session/session';

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

export default function Login(_: InferGetServerSidePropsType<typeof getServerSideProps>) {
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
        const response = await login(values);
        router.push((router.query?.callbackUrl as string) ?? PAGE_ROUTES.DASHBOARD);
        response && setResponse(response);
      } catch (error) {
        console.error(error.message);
        console.error(error);
      }
    },
    [router]
  );

  const onError = console.error;

  return (
    <VStack h="full" align="center" justify="center">
      <Heading color="gray.500">Login</Heading>
      <Box w="24rem">
        <VStack as="form" onSubmit={handleSubmit(onSubmit, onError)}>
          <FormControl isInvalid={Boolean(errors.email)}>
            <Input
              {...register('email', { required: 'required' })}
              placeholder="Email"
              type="email"
              bg="gray.700"
            />
            <FormErrorMessage>{errors.email?.message}</FormErrorMessage>
          </FormControl>
          <FormControl isInvalid={Boolean(errors.password)}>
            <Input
              {...register('password', { required: 'required' })}
              placeholder="Password"
              type="password"
              bg="gray.700"
            />
            <FormErrorMessage>{errors.password?.message}</FormErrorMessage>
          </FormControl>
          <Button size="sm" w="full" type="submit" colorScheme="primary">
            Login
          </Button>
          <Button
            as={Link}
            variant="ghost"
            size="sm"
            w="full"
            type="submit"
            colorScheme="primary"
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

export const getServerSideProps = withSessionSsr(async function getServerSideProps(
  context: GetServerSidePropsContext
) {
  if (context.req.session?.user?.verified) {
    return {
      redirect: {
        destination: PAGE_ROUTES.DASHBOARD,
        permanent: false,
      },
    };
  }

  return {
    props: {},
  };
});
