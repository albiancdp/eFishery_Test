import { body } from 'express-validator';

export function create_validator() {
  return [
    body('bookId')
      .notEmpty().withMessage('please input bookId'),

    body('title')
      .notEmpty().withMessage('please input title'),

    // body('thumbnail')
    //   .notEmpty().withMessage('please input thumbnail'),

    // body('author')
    //   .isArray().withMessage('please input valid author')
    //   .notEmpty().withMessage('please input author'),

    // body('rating')
    //   .notEmpty().withMessage('please input rating'),

    // body('description')
    //   .notEmpty().withMessage('please input description'),

    // body('link')
    //   .notEmpty().withMessage('please input link'),
  ];
};
