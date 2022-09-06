import http from 'http';
import https from 'https';
import fs from 'fs';
import app from './app/app';
import configs from './configs/global_config';

const PORT = configs.server.port;
const HTTPS_ENABLE = configs.server.https_enable;
const SSL_CERT_DIR = configs.server.ssl_cert_dir;

app.set('port', PORT);

if (HTTPS_ENABLE === 'true') {
  const httpsOptions = {
    key: fs.readFileSync(`${SSL_CERT_DIR}/key.pem`),
    cert: fs.readFileSync(`${SSL_CERT_DIR}/cert.pem`),
    ca: [
      fs.readFileSync(`${SSL_CERT_DIR}/chain.pem`),
      fs.readFileSync(`${SSL_CERT_DIR}/fullchain.pem`)
    ]
  };
  https.createServer(httpsOptions, app).listen(PORT, (err) => {
    if (err) {
      process.exit(1);
    } else {
      // console.log(`Server running on port ${PORT}`);
    }
  });
} else {
  http.createServer(app).listen(PORT, (err) => {
    if (err) {
      process.exit(1);
    } else {
      //console.log(`Server running on port ${PORT}`);
    }
  });
};
