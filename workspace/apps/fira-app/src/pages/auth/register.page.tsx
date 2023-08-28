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
import type { InferGetServerSidePropsType } from 'next';
import Link from 'next/link';
import { useRouter } from 'next/router';
import React, { useCallback, useState } from 'react';
import { useForm } from 'react-hook-form';

import { PAGE_ROUTES } from 'src/config/routes';
import { signup } from 'src/lib/auth';
import { withAuthSsr } from 'src/lib/session/authed';
import { withSessionSsr } from 'src/lib/session/session';

type FormValues = {
  full_name: string;
  email_address: string;
  password: string;
};

type RegisterResponse = {
  error?: string;
  ok: boolean;
  status: number;
  url: null | string;
};

export default function Register(
  _: InferGetServerSidePropsType<typeof getServerSideProps>
) {
  const [response, setResponse] = useState<null | RegisterResponse>(null);
  const router = useRouter();
  const {
    handleSubmit,
    register,
    formState: { errors, isSubmitting },
  } = useForm<FormValues>();

  const onSubmit = useCallback(
    async (values: FormValues) => {
      try {
        setResponse(null);
        const res = await signup(values);
        setResponse(res.data);
        // TODO [Jake]: Implement verification email flow
        router.push((router.query?.callbackUrl as string) ?? PAGE_ROUTES.DASHBOARD);
      } catch (error) {
        console.error(error);
      }
    },
    [router]
  );

  const onError = console.error;

  return (
    <VStack h="full" align="center" justify="center">
      <Heading color="gray.500">Register</Heading>
      <Box w="24rem">
        <VStack as="form" onSubmit={handleSubmit(onSubmit, onError)}>
          <FormControl isInvalid={Boolean(errors.full_name)}>
            <Input
              {...register('full_name', { required: 'Full name is required' })}
              placeholder="Name"
              type="text"
              bg="gray.700"
            />
            <FormErrorMessage>{errors.full_name?.message}</FormErrorMessage>
          </FormControl>
          <FormControl isInvalid={Boolean(errors.email_address)}>
            <Input
              {...register('email_address', { required: 'Email address is required' })}
              placeholder="Email"
              type="email"
              bg="gray.700"
              color="gray.100"
            />
            <FormErrorMessage>{errors.email_address?.message}</FormErrorMessage>
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
              bg="gray.700"
              color="gray.100"
            />
            <FormErrorMessage>{errors.password?.message}</FormErrorMessage>
          </FormControl>
          <Button
            colorScheme="primary"
            isLoading={isSubmitting}
            size="sm"
            type="submit"
            w="full"
          >
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
          {!!response?.error && (
            <Text color="red">
              {response?.error} - Status {response?.status}
            </Text>
          )}
        </VStack>
      </Box>
    </VStack>
  );
}

export const getServerSideProps = withSessionSsr(
  withAuthSsr(async function getServerSideProps() {
    return {
      props: {},
    };
  })
);
