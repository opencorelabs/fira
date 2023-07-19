import {
  Box,
  Button,
  FormControl,
  FormErrorMessage,
  Heading,
  Input,
  VStack,
} from '@chakra-ui/react';
import type { GetServerSidePropsContext, InferGetServerSidePropsType } from 'next';
import Link from 'next/link';
import { useRouter } from 'next/router';
import React, { useCallback, useState } from 'react';
import { useForm } from 'react-hook-form';

import { PAGE_ROUTES } from 'src/config/routes';
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
  const [, setResponse] = useState<null | Record<string, unknown>>(null);
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
        router.push(PAGE_ROUTES.VERIFY_EMAIL);
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
          <FormControl isInvalid={Boolean(errors.name)}>
            <Input
              {...register('name', { required: 'Name is required' })}
              placeholder="Name"
              type="text"
              bg="gray.700"
            />
            <FormErrorMessage>{errors.name?.message}</FormErrorMessage>
          </FormControl>
          <FormControl isInvalid={Boolean(errors.email)}>
            <Input
              {...register('email', { required: 'Email is required' })}
              placeholder="Email"
              type="email"
              bg="gray.700"
              color="gray.100"
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
              bg="gray.700"
              color="gray.100"
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
          {/* 
          // TODO: Implement error handling
          {!!response?.errorMessage && (
            <Text color="red">
              {response?.errorMessage} - Status {response?.status}
            </Text>
          )} */}
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
