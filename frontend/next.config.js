/** @type {import('next').NextConfig} */

const nextConfig = {
  async redirects() {
    return [
      {
        source: "/",
        destination: "/dashboard",
        permanent: true,
      }
    ]
  },
  reactStrictMode: true,
};

module.exports = nextConfig;