import express from 'express';
const router = express.Router();
// import controller
import configs from '../../configs/global_config';
import indexController from '../../controllers/index';
import indexValidator from '../../validators/index';

router.route(configs.routeBase + '/auth/register')
  .post(indexValidator.register(), indexController.authController.register);

router.route(configs.routeBase + '/auth/login')
  .post(indexValidator.login(), indexController.authController.login);

export default router;
