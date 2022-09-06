import dotenv from 'dotenv';
dotenv.config();

const NODE_ENV = (process.env.NODE_ENV).toLowerCase()  || 'development';
const config = {
  node_env: NODE_ENV,
  server: {
    port: process.env.PORT || 3000,
    https_enable: process.env.HTTPS_ENABLE || false,
    ssl_cert_dir: process.env.SSL_CERT_DIR || '.'
  },
  routeBase: process.env.ROUTE_BASE || '/api/v1',
};

export default config;
