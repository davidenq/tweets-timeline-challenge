/** @type {import('next').NextConfig} */
const nextConfig = {
  reactStrictMode: true,
  swcMinify: true,
  publicRuntimeConfig: {
    apiEndpoint: process.env.API_ENDPOINT,
    apiPort: process.env.API_PORT
  }
}

module.exports = nextConfig
