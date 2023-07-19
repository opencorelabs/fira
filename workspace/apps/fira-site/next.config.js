/** @type {import('next').NextConfig} */
const config = {
  reactStrictMode: true,
  output: 'export',
  pageExtensions: ['page.tsx', 'api.ts'],
  images: {
    unoptimized: true,
  },
};

module.exports = config;
