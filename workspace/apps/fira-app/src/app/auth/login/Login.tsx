'use client';

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
import { useRouter, useSearchParams } from 'next/navigation';
import { signIn } from 'next-auth/react';
import React, { useCallback, useState } from 'react';
import { useForm } from 'react-hook-form';

type FormValues = {
  email: string;
  password: string;
};

type SignInResponse = {
  error?: string | null;
  ok: boolean;
  status: number;
  url: null | string;
};

type LoginProps = {
  csrfToken?: string;
};

export function Login({ csrfToken }: LoginProps) {
  const [response, setResponse] = useState<SignInResponse | null>(null);
  const router = useRouter();
  const params = useSearchParams();
  const {
    handleSubmit,
    register,
    formState: { errors },
  } = useForm<FormValues>();

  const onSubmit = useCallback(
    async (values: FormValues) => {
      try {
        setResponse(null);
        const result = await signIn('credentials', {
          ...values,
          redirect: false,
        });
        if (result?.ok && !result?.error) {
          router.push(params?.get('callbackUrl') ?? '/');
        }
        result && setResponse(result);
      } catch (error) {
        console.error(error);
      }
    },
    [params, router]
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

// export async function getServerSideProps(context: GetServerSidePropsContext) {
//   const session = await getServerSession(context.req, context.res, authOptions);
//   if (session) {
//     return {
//       redirect: {
//         destination: '/',
//         permanent: false,
//       },
//     };
//   }
//   const csrfToken = await getCsrfToken(context);
//   return {
//     props: { csrfToken },
//   };
// }
