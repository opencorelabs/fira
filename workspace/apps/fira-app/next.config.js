/** @type {import('next').NextConfig} */
const config = {
  reactStrictMode: true,
  pageExtensions: ['page.tsx', 'api.ts'],
  basePath: '/app',
  images: {
    domains: ['fira.opencorelabs.com'],
    path: '/_next/image',
  },
};

module.exports = config;
