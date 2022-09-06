import { query } from 'express-validator';

export function list_validator() {
  return [
    query('page')
      .notEmpty().withMessage('please input page'),

    query('limit')
      .notEmpty().withMessage('please input limit'),
  ];
};
