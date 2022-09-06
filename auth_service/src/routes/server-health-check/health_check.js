import express from 'express';
const router = express.Router();
// import controller
import indexController from '../../controllers/index';

router.route('/check-health')
  .get(indexController.serverCheckHealthController.checkHealth);

export default router;
