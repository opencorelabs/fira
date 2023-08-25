import type { GetServerSidePropsContext, GetServerSidePropsResult } from 'next';

import { PAGE_ROUTES } from 'src/config/routes';

export function withAuthSsr<
  P extends { [key: string]: unknown } = { [key: string]: unknown }
>(
  getServerProps: (
    context: GetServerSidePropsContext
  ) => GetServerSidePropsResult<P> | Promise<GetServerSidePropsResult<P>>
) {
  return async (
    context: GetServerSidePropsContext
  ): Promise<GetServerSidePropsResult<P>> => {
    if (context.req.session?.user?.verified) {
      return {
        redirect: {
          destination: PAGE_ROUTES.DASHBOARD,
          permanent: false,
        },
      };
    }
    return getServerProps(context);
  };
}
