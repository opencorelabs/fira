import { Box, Button, Container, Text } from '@chakra-ui/react';
import { V1AccountNamespace } from '@fira/api-sdk';
import { GetServerSidePropsContext } from 'next';
import { getServerSession } from 'next-auth/next';
import { useCallback } from 'react';

import { getApi } from 'src/lib/fira-api';

import { authOptions } from '../api/auth/[...nextauth].api';

type FormValues = {
  token: string;
};

export default function VerifyEmail() {
  const handleRequestVerifyLink = useCallback(async () => {
    //
  }, []);

  return (
    <Container maxW="container.xl">
      <Box>
        <Text>Verify Email</Text>
        <Text>
          Didn't get an email?{' '}
          <Button variant="link" onClick={handleRequestVerifyLink}>
            Request a verification link
          </Button>
        </Text>
      </Box>
    </Container>
  );
}

export async function getServerSideProps(context: GetServerSidePropsContext) {
  if (context.query?.token) {
    // Send validation request to API and rediect to dashboard
    const response = await getApi().firaServiceVerifyAccount({
      // @ts-expect-error type is required
      type: 1,
      namespace: V1AccountNamespace.ACCOUNT_NAMESPACE_CONSUMER,
      token: context.query.verification_token as string,
    });
    console.info('response', response);
    return {
      redirect: {
        destination: '/dashboard',
        permanent: false,
      },
    };
  }

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
