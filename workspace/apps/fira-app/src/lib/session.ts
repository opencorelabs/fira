import { withIronSessionApiRoute, withIronSessionSsr } from 'iron-session/next';
import type {
  GetServerSidePropsContext,
  GetServerSidePropsResult,
  NextApiHandler,
} from 'next';

declare module 'iron-session' {
  interface IronSessionData {
    user?: {
      id: number;
      verified: boolean;
      token?: string;
      // TODO: update with user info, name, role, etc.
      // admin?: boolean;
    };
  }
}

const sessionOptions = {
  password: process.env.SESSION_SECRET ?? 'complex_password_at_least_32_characters_long',
  cookieName: 'fira-session',
  cookieOptions: {
    secure: process.env.NODE_ENV === 'production',
  },
};

export function withSessionRoute(handler: NextApiHandler) {
  return withIronSessionApiRoute(handler, sessionOptions);
}

export function withSessionSsr<
  P extends { [key: string]: unknown } = { [key: string]: unknown }
>(
  handler: (
    context: GetServerSidePropsContext
  ) => GetServerSidePropsResult<P> | Promise<GetServerSidePropsResult<P>>
) {
  return withIronSessionSsr(handler, sessionOptions);
}
