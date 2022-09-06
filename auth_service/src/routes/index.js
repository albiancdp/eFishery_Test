import express from 'express';
// import router
import defaultRouter from './default/index';
import healthCheckRouter from './server-health-check/index';
import authRouter from './auth/index';

const router = express.Router();

router.use(defaultRouter);
router.use(healthCheckRouter);
router.use(authRouter);

export default router;
