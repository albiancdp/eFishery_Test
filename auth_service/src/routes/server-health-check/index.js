import express from 'express';
import healthCheck from './health_check';

const router = express.Router();

router.use(healthCheck);

export default router;
