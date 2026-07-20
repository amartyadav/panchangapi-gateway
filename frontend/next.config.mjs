// next.config.js
/** @type {import('next').NextConfig} */
const nextConfig = {
    async rewrites() {
      return [
        {
          source: '/', // The route to redirect from
          destination: '/landingpage', // The route to redirect to
        },
      ];
    },
  };
  
  export default nextConfig;
  