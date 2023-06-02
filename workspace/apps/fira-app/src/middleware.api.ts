// /middleware.ts
import { getIronSession } from 'iron-session/edge';
import type { NextRequest } from 'next/server';
import { NextResponse } from 'next/server';

import { PAGE_ROUTES } from 'src/config/routes';
import { options } from 'src/lib/session/utils';

const strings = [
  '/api',
  '/auth',
  '/_next',
  '/static',
  '/public',
  '/favicon.ico',
  '/robots.txt',
  '/images',
];

export const middleware = async (req: NextRequest) => {
  const res = NextResponse.next();
  const session = await getIronSession(req, res, options);
  const hasMatch = strings.some(
    (str) => req.nextUrl.pathname.startsWith(str) || req.nextUrl.pathname === '/'
  );
  if (hasMatch) {
    return NextResponse.next();
  }
  const { user } = session;

  // if user is not logged in, redirect to login page
  if (!user) {
    return NextResponse.redirect(new URL(PAGE_ROUTES.LOGIN, req.url));
  }

  // if user is not verified, redirect to verify-email page
  if (!user?.verified) {
    return NextResponse.redirect(new URL(PAGE_ROUTES.VERIFY_EMAIL, req.url));
  }

  return res;
};
