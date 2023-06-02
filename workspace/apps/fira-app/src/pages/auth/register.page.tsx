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
import { V1CreateAccountResponse } from '@fira/api-sdk';
import type { GetServerSidePropsContext, InferGetServerSidePropsType } from 'next';
import Link from 'next/link';
import { useRouter } from 'next/router';
import React, { useCallback, useState } from 'react';
import { useForm } from 'react-hook-form';

import { routes } from 'src/config/routes';
import { signup } from 'src/lib/auth';
import { withSessionSsr } from 'src/lib/session/session';

type FormValues = {
  name: string;
  email: string;
  password: string;
};

export default function Register(
  _: InferGetServerSidePropsType<typeof getServerSideProps>
) {
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
        const response = await signup(values);
        setResponse(response.data);
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
          {/* <input name="csrfToken" type="hidden" defaultValue={csrfToken} /> */}
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
          <Button size="sm" w="full" type="submit" colorScheme="primary">
            Sign up
          </Button>
          <Button
            as={Link}
            variant="ghost"
            size="sm"
            w="full"
            type="submit"
            colorScheme="primary"
            href="/auth/login"
          >
            Have an account? Login
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

export const getServerSideProps = withSessionSsr(async function getServerSideProps(
  context: GetServerSidePropsContext
) {
  if (context.req.session?.user?.verified) {
    return {
      redirect: {
        destination: routes.dashboard,
        permanent: false,
      },
    };
  }
  return {
    props: {},
  };
});
