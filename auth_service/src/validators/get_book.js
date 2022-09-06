import { query } from 'express-validator';

export function get_book_validator() {
  return [
    query('q')
      .notEmpty().withMessage('please input q'),
  ];
};
