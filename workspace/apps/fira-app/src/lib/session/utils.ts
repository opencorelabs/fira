declare module 'iron-session' {
  interface IronSessionData {
    user?: {
      id?: string;
      verified: boolean;
      token?: string;
      refresh?: string;
      name?: string;
      email?: string;
    };
  }
}

export const options = {
  password: process.env.SESSION_SECRET ?? 'complex_password_at_least_32_characters_long',
  cookieName: 'fira:session',
  cookieOptions: {
    secure: process.env.NODE_ENV === 'production',
  },
};
