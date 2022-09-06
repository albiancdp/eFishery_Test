import { body } from 'express-validator';

const role = ['admin', 'user'];

export function register() {
  return [
    body('name')
      .notEmpty().withMessage('please input name'),

    body('phone')
      .notEmpty().withMessage('please input phone'),

    body('role')
      .notEmpty().withMessage('please input role')
      .isIn(role).withMessage('please input valid role'),
  ];
};
