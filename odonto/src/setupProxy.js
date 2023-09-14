// eslint-disable-next-line import/no-extraneous-dependencies
const { createProxyMiddleware } = require('http-proxy-middleware');

const target = 'http://localhost:18080';

module.exports = (app) => {
  app.use('/pub', createProxyMiddleware({
    target,
    changeOrigin: true,
  }));
  app.use('/priv', createProxyMiddleware({
    target,
    changeOrigin: true,
  }));
};
