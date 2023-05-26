/** @type {import('next').NextConfig} */
const config = {
  reactStrictMode: true,
  pageExtensions: ['page.tsx', 'api.ts'],
  images: {
    domains: [process.env.NEXT_PUBLIC_BASE_URL],
    path: '/_next/image',
  },
};

module.exports = config;
