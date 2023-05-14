import { V1AccountCredentialType, V1AccountNamespace } from '@fira/api-sdk';
import NextAuth, { NextAuthOptions } from 'next-auth';
import CredentialsProvider from 'next-auth/providers/credentials';

import { api } from 'src/lib/fira-api';

export const options: NextAuthOptions = {
  pages: {
    signIn: '/auth/login',
  },
  session: {
    strategy: 'jwt',
    maxAge: 24 * 60 * 60, // 1 days
  },
  events: {
    async signIn(message) {
      console.info('message', message);
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
            namespace: V1AccountNamespace.ACCOUNT_NAMESPACE_UNSPECIFIED,
            credential: {
              credentialType: V1AccountCredentialType.ACCOUNT_CREDENTIAL_TYPE_UNSPECIFIED,
              emailCredential: {
                email: credentials.email,
                password: credentials.password,
              },
            },
          });
          // const res = await fetch('/api/v1/accounts/login', {
          //   method: 'POST',
          //   body: JSON.stringify(credentials),
          //   headers: { 'Content-Type': 'application/json' },
          // });
          // const user = await res.json();
          console.info('response', response);
        } catch (error) {
          console.info('error', error);
        }

        // const user = users[credentials.email];
        // if (user && user.password === credentials.password) {
        //   // estlint disable for `password`
        //   // eslint-disable-next-line unused-imports/no-unused-vars
        //   const { password, ...rest } = user;
        //   return rest;
        // }
        // user not found
        return null;
      },
    }),
  ],
};

export default NextAuth(options);
