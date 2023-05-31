/** @type {import('next').NextConfig} */
const config = {
  reactStrictMode: true,
  pageExtensions: ['page.tsx', 'api.ts'],
  images: {
    path: '/_next/image',
  },
};

module.exports = config;
