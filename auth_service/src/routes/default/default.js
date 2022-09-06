import express from 'express';
const router = express.Router();
// import controller
import indexController from '../../controllers/index';

router.route('/')
  .get(indexController.defaultController.defaults);

export default router;
