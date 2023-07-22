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
  LOGIN: '/private-api/auth/login',
  LOGOUT: '/private-api/auth/logout',
  REGISTER: '/private-api/auth/register',
  VERIFY_EMAIL: '/private-api/auth/verify-email',
  HEALTH_CHECK: '/private-api/health-check',
};

export const API_WITH_BASEPATH = makeRoutes(API_ROUTES);
