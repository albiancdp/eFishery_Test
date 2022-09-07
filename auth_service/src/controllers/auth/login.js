import { validationResult } from 'express-validator';
import indexService from '../../services/index';
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
    const resLogin = await indexDomain.authDomain.login(req.body);
    if (!resLogin) return responseHelper.notFound(res, 'User Not Found');
    if (!indexService.bcrypt.checkHash(req.body.password, resLogin.password)) {
      return responseHelper.errorCustom(res, 'validate error', 'Wrong Password', { password: 'Wrong Password' });
    }
    const dataUser = {
      name: resLogin.name,
      phone: resLogin.phone,
      role: resLogin.role,
      created_at: resLogin.created_at
    };
    const dataToken = {
      token: indexService.jwt.generateToken(dataUser)
    };
    return responseHelper.success(res, 'Login Success', dataToken);
  }
  catch (err) {
    return responseHelper.errorService(res, err.message);
  };
};

export default register;
