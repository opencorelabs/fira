function makeRoutes(routes: Record<string, string>, basePath = '') {
  return Object.fromEntries(
    Object.keys(routes).map((key) => {
      return [key, `${basePath}${routes[key]}`];
    })
  );
}

export const PAGE_ROUTES = {
  DASHBOARD: '/dashboard',
  LOGIN: '/auth/login',
  REGISTER: '/auth/register',
  VERIFY_EMAIL: '/auth/verify-email',
};

export const API_ROUTES = {
  LOGIN: '/api/auth/login',
  LOGOUT: '/api/auth/logout',
  REGISTER: '/api/auth/register',
  VERIFY_EMAIL: '/api/auth/verify-email',
  HEALTH_CHECK: '/api/health-check',
};

export const API_WITH_BASEPATH = makeRoutes(
  API_ROUTES,
  process.env.NEXT_PUBLIC_BASE_PATH
);
