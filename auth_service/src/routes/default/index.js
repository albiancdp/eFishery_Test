import express from 'express';
import def from './default';

const router = express.Router();

router.use(def);

export default router;
