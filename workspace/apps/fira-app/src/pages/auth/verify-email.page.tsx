import { Box, Button, Container, Flex, Heading, Text, useToast } from '@chakra-ui/react';
import { GetServerSidePropsContext } from 'next';
import { useCallback } from 'react';

import { PAGE_ROUTES } from 'src/config/routes';
import { withSessionSsr } from 'src/lib/session/session';

export default function VerifyEmail() {
  const toast = useToast();
  const handleRequestVerifyLink = useCallback(async () => {
    toast({
      title: 'Not Implemented',
      status: 'error',
      duration: 4000,
      isClosable: true,
    });
  }, [toast]);

  return (
    <Container maxW="container.xl">
      <Flex justify="space-between">
        <Box>
          <Heading color="gray.500">Verify Email</Heading>
          <Text color="gray.500">
            Didn't recieve an email?{' '}
            <Button
              variant="link"
              onClick={handleRequestVerifyLink}
              colorScheme="primary"
            >
              Request a new verification link
            </Button>
          </Text>
        </Box>
        <Box></Box>
      </Flex>
    </Container>
  );
}

export const getServerSideProps = withSessionSsr(async function getServerSideProps(
  context: GetServerSidePropsContext
) {
  try {
    if (context.req.session?.user?.verified) {
      return {
        redirect: {
          destination: PAGE_ROUTES.DASHBOARD,
          permanent: false,
        },
      };
    }

    if (context.query?.verification_token) {
      // TODO: Implement verification
      // // Send validation request to API and rediect to dashboard
      // const response = await getApi().firaServiceVerifyAccount({
      //   // @ts-expect-error type is required
      //   type: 1,
      //   namespace: V1AccountNamespace.ACCOUNT_NAMESPACE_CONSUMER,
      //   token: context.query.verification_token as string,
      // });
      // context.req.session.user = {
      //   ...context.req.session.user,
      //   verified: true,
      //   token: response.data.jwt,
      // };
      // await context.req.session.save();
      return {
        redirect: {
          destination: PAGE_ROUTES.DASHBOARD,
          permanent: false,
        },
      };
    }
  } catch (error) {
    console.error('\n\nerror', error);
    // TODO: Return error message to client
    return {
      props: {},
    };
  }

  if (!context.req.session?.user) {
    return {
      redirect: {
        destination: PAGE_ROUTES.LOGIN,
        permanent: false,
      },
    };
  }

  return {
    props: {},
  };
});
