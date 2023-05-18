import { V1AccountCredentialType, V1AccountNamespace } from '@fira/api-sdk';
import NextAuth, { NextAuthOptions } from 'next-auth';
import CredentialsProvider from 'next-auth/providers/credentials';

import { api } from 'src/lib/fira-api';

export const authOptions: NextAuthOptions = {
  pages: {
    signIn: '/auth/login',
    newUser: '/auth/register',
  },
  session: {
    strategy: 'jwt',
    maxAge: 24 * 60 * 60, // 24 hours (same as server-side)
  },
  events: {
    async signIn(message) {
      console.info('message', message);
    },
  },
  callbacks: {
    session: async ({ session, token }) => {
      if (token.accessToken) {
        // @ts-expect-error accessToken not defined on session
        session.accessToken = token.accessToken;
      }
      return session;
    },
    async jwt({ token, user }) {
      if (user) {
        // @ts-expect-error token not defined on user
        token.accessToken = user.token;
      }
      return token;
    },
  },
  providers: [
    CredentialsProvider({
      name: 'Login',
      type: 'credentials',
      credentials: {
        email: { label: 'email', type: 'email' },
        password: { label: 'password', type: 'password' },
      },
      async authorize(credentials) {
        try {
          const response = await api.firaServiceLoginAccount({
            namespace: V1AccountNamespace.ACCOUNT_NAMESPACE_CONSUMER,
            credential: {
              credentialType: V1AccountCredentialType.ACCOUNT_CREDENTIAL_TYPE_EMAIL,
              emailCredential: {
                email: credentials?.email,
                password: credentials?.password,
              },
            },
          });
          if (!response.ok) return null;
          // Not implemented yet
          // Needed for user id
          // const me = await api.firaServiceGetAccount({
          //   headers: {
          //     authorization: `Bearer ${response.data.jwt}`,
          //   },
          // });
          // console.log('me', me);
          return { id: '1', email: credentials?.email, token: response.data.jwt };
        } catch (error) {
          console.info('error', error);
        }

        return null;
      },
    }),
  ],
};

export default NextAuth(authOptions);
