import { validationResult } from 'express-validator';
import responseHelper from '../../helpers/utils/response';
import indexService from '../../services/index';
import validateHelper from '../../helpers/utils/validator_rebuild';

const register = async (req, res) => {
  try {
    const notValids = validationResult(req);
    if (!notValids.isEmpty()) {
      const valid = validateHelper(notValids);
      return responseHelper.errorValidate(res, valid);
    };
    const token = req.headers.authorization;
    if(!token) return responseHelper.errorAuthNotFound(res);
    const dataDecrypt = indexService.jwt.verifyToken(token);
    if (!dataDecrypt) return responseHelper.errorUnauthorized(res);
    delete dataDecrypt.iat;
    delete dataDecrypt.exp;
    return responseHelper.success(res, 'Create User Success', dataDecrypt);
  }
  catch (err) {
    return responseHelper.errorService(res, err.message);
  };
};

export default register;
