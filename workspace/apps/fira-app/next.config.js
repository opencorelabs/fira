/** @type {import('next').NextConfig} */

const config = {
  reactStrictMode: true,
  pageExtensions: ['page.tsx', 'api.ts'],
  basePath: process.env.NEXT_PUBLIC_BASE_PATH,
  images: {
    domains: [process.env.NEXT_PUBLIC_BASE_URL],
    path: '/_next/image',
  },
  async rewrites() {
    return [
      {
        source: '/',
        destination: process.env.NEXT_PUBLIC_BASE_PATH,
      },
    ];
  },
};

module.exports = config;
