export const routes = {
  dashboard: `${process.env.NEXT_PUBLIC_BASE_PATH}/dashboard`,
  login: `${process.env.NEXT_PUBLIC_BASE_PATH}/auth/login`,
  register: `${process.env.NEXT_PUBLIC_BASE_PATH}/auth/register`,
  verifyEmail: `${process.env.NEXT_PUBLIC_BASE_PATH}/auth/verify-email`,
};

export const api = {
  login: `${process.env.NEXT_PUBLIC_BASE_PATH}/api/auth/login`,
  logout: `${process.env.NEXT_PUBLIC_BASE_PATH}/api/auth/logout`,
  register: `${process.env.NEXT_PUBLIC_BASE_PATH}/api/auth/register`,
  verifyEmail: `${process.env.NEXT_PUBLIC_BASE_PATH}/api/auth/verify-email`,
  healthCheck: `${process.env.NEXT_PUBLIC_BASE_PATH}/api/health-check`,
};
