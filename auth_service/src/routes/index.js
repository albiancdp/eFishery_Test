import express from 'express';
// import router
import defaultRouter from './default/index';
import healthCheckRouter from './server-health-check/index';

const router = express.Router();

router.use(defaultRouter);
router.use(healthCheckRouter);

export default router;
