import { Box, Button, Container, Input, Text } from '@chakra-ui/react';
import { V1AccountNamespace } from '@fira/api-sdk';
import { GetServerSidePropsContext } from 'next';
import { useRouter } from 'next/router';
import { getServerSession } from 'next-auth/next';
import { useCallback } from 'react';
import { useForm } from 'react-hook-form';

import { getApi } from 'src/lib/fira-api';

import { authOptions } from '../api/auth/[...nextauth].api';

type FormValues = {
  token: string;
};

export default function VerifyEmail() {
  const router = useRouter();
  const { register, handleSubmit } = useForm<FormValues>();

  const handleVerify = useCallback(
    async (data: FormValues) => {
      const response = await getApi().firaServiceVerifyAccount({
        // @ts-expect-error type is required
        type: 1,
        namespace: V1AccountNamespace.ACCOUNT_NAMESPACE_CONSUMER,
        token: data.token,
      });
      console.info('response', response);
      router.push('/dashboard');
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
        <Button type="submit" mt={2} colorScheme="blue">
          Verify Email
        </Button>
      </Box>
    </Container>
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
  return {
    props: {},
  };
}
