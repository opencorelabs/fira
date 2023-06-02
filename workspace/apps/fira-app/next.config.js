/** @type {import('next').NextConfig} */
const config = {
  reactStrictMode: true,
  pageExtensions: ['page.tsx', 'api.ts'],
  basePath: process.env.NEXT_PUBLIC_BASE_PATH,
  images: {
    domains: [process.env.NEXT_PUBLIC_BASE_URL],
    path: '/_next/image',
  },
  redirects: async () => [
    {
      source: '/',
      destination: process.env.NEXT_PUBLIC_BASE_PATH,
      permanent: false,
      basePath: false,
    },
  ],
};

module.exports = config;
