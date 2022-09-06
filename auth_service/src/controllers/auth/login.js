import { validationResult } from 'express-validator';
import responseHelper from '../../helpers/utils/response';
import indexDomain from '../../domains/index';
import validateHelper from '../../helpers/utils/validator_rebuild';

const register = async (req, res) => {
  try {
    const notValids = validationResult(req);
    if (!notValids.isEmpty()) {
      const valid = validateHelper(notValids);
      return responseHelper.errorValidate(res, valid);
    };
    const resLogin = await indexDomain.authDomain.logins(req.body);
    if (!resLogin) return responseHelper.notFound(res, 'User Not Found');
    return responseHelper.success(res, 'Create Wishlist', resLogin.data);
  }
  catch (err) {
    return responseHelper.errorService(res, err.message);
  };
};

export default register;
