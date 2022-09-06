import { validationResult } from 'express-validator';
import responseHelper from '../../helpers/utils/response';
import indexDomain from '../../domains/index';
import validateHelper from '../../helpers/utils/validator_rebuild';
import helperCommon from '../../helpers/commons/index';

const register = async (req, res) => {
  try {
    const notValids = validationResult(req);
    if (!notValids.isEmpty()) {
      const valid = validateHelper(notValids);
      return responseHelper.errorValidate(res, valid);
    };
    req.body.password = helperCommon.password.generateChar(4);
    const resRegister = await indexDomain.authDomain.register(req.body);
    if (resRegister?.notValid) return responseHelper.errorValidate(res, resRegister.notValid);
    return responseHelper.success(res, 'Create User Success', resRegister);
  }
  catch (err) {
    return responseHelper.errorService(res, err.message);
  };
};

export default register;
