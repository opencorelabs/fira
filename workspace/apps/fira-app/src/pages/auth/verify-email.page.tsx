import { Box, Button, Container, Text } from '@chakra-ui/react';
import { V1AccountNamespace } from '@fira/api-sdk';
import { GetServerSidePropsContext } from 'next';
import { useCallback } from 'react';

import { getApi } from 'src/lib/fira-api';
import { withSessionSsr } from 'src/lib/session/session';

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
            Request a new verification link
          </Button>
        </Text>
      </Box>
    </Container>
  );
}

export const getServerSideProps = withSessionSsr(async function getServerSideProps(
  context: GetServerSidePropsContext
) {
  try {
    if (context.query?.verification_token) {
      // Send validation request to API and rediect to dashboard
      const response = await getApi().firaServiceVerifyAccount({
        // @ts-expect-error type is required
        type: 1,
        namespace: V1AccountNamespace.ACCOUNT_NAMESPACE_CONSUMER,
        token: context.query.verification_token as string,
      });
      if (!response.ok) {
        throw new Error('response.error');
      }
      return {
        redirect: {
          destination: '/dashboard',
          permanent: false,
        },
      };
    }
  } catch (error) {
    // TODO: Return error message to client
    return {
      props: {},
    };
  }

  return {
    props: {},
  };
});
