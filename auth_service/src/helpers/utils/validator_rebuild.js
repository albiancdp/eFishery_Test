const validate = (errors) => {

  const extractedErrors = {};
  errors.array().map(err => [
    extractedErrors[err.param] = err.msg
  ]);

  return extractedErrors;

};

export default validate;
