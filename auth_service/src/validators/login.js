import { body } from 'express-validator';

export function login() {
  return [
    body('password')
      .notEmpty().withMessage('please input password'),

    body('phone')
      .notEmpty().withMessage('please input phone'),
  ];
};
