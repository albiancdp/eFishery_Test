import response_code from './response_code';
import response_message from './response_response';

const response = {
  success: (res, msg, data = {}) => {
    const result = {
      status: true,
      code: response_code.success,
      message: msg,
      data: data,
      error_code: null,
      errors: {}
    };
    res.header('Content-Type','application/json');
    return res.status(response_code.success).send(JSON.stringify(result, null, 2) + '\n');
  },
  notFound: (res, msg = 'Data Not Found') => {
    const result = {
      status: false,
      code: response_code.notFound,
      message: msg,
      data: {},
      error_code: response_message.errorNotFound,
      errors: {}
    };
    res.header('Content-Type','application/json');
    return res.status(response_code.notFound).send(JSON.stringify(result, null, 2) + '\n');
  },
  errorValidate: (res, notValid = {}, msg = 'Validate Error', data = {}) => {
    const result = {
      status: false,
      code: response_code.validationError,
      message: msg,
      data: data,
      error_code: response_message.errorValidate,
      errors: notValid,
    };
    res.header('Content-Type','application/json');
    return res.status(response_code.validationError).send(JSON.stringify(result, null, 2) + '\n');
  },
  errorCustom: (res, error = 'errorServer', msg = 'Internal Server Error', notValid = {}, data = {}) => {
    const result = {
      status: false,
      code: response_code.badRequest,
      message: msg,
      data: data,
      error_code: response_message[error],
      errors: notValid,
    };
    res.header('Content-Type','application/json');
    return res.status(response_code.badRequest).send(JSON.stringify(result, null, 2) + '\n');
  },
  errorAuthNotFound: (res, msg = 'No token provided !!!', data = {}) => {
    const result = {
      status: false,
      code: response_code.authNotFound,
      message: msg,
      data: data,
      error_code: response_message.errorAuthNotFound,
      errors: {},
    };
    res.header('Content-Type','application/json');
    return res.status(response_code.authNotFound).send(JSON.stringify(result, null, 2) + '\n');
  },
  errorUnauthorized: (res, msg = 'Invalid Token !!!', data = {}) => {
    const result = {
      status: false,
      code: response_code.unAuthorizedRequest,
      message: msg,
      data: data,
      error_code: response_message.errorUnauthorized,
      errors: {},
    };
    res.header('Content-Type','application/json');
    return res.status(response_code.unAuthorizedRequest).send(JSON.stringify(result, null, 2) + '\n');
  },
  errorService: (res, msg = 'Internal Server Error', data = {}) => {
    const result = {
      status: false,
      code: response_code.internalServerError,
      message: msg,
      data: data,
      error_code: response_message.errorServer,
      errors: {},
    };
    res.header('Content-Type','application/json');
    return res.status(response_code.internalServerError).send(JSON.stringify(result, null, 2) + '\n');
  },
};

export default response;
