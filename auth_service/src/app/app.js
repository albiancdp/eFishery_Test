import express from 'express';
import cors from 'cors';
import helmet from 'helmet';
import morgan from 'morgan';
import rateLimit from 'express-rate-limit';
import compression from 'compression';
import rateLimitConfig from '../configs/rate_limit';
// import corsConfig from '../configs/cors';
import routers from '../routes/index';

const app = express();
const apiLimiter = rateLimit(rateLimitConfig);
const apiCors = cors();

app.use(apiCors);
app.use(apiLimiter);
app.use(express.json());
app.use(express.urlencoded({ extended: true }));
app.use(helmet());
app.use(compression());
app.use(morgan('dev'));
app.disable('x-powered-by');
// import routers
app.use(routers);

// handle route not found
app.use((req, res, next) => {
  res.status(404).send({
    'status': false,
    'code': 404,
    'message': 'Not Found',
    'data': {},
    'error_code': 'not_found',
    'errors': {}
  });
});

export default app;
