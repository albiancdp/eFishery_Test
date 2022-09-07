import jwt from 'jsonwebtoken';
import configs from '../configs/global_config';

const verifyToken = (token) => {
  try {
    const verified = jwt.verify(token, configs.jwt.key);
    if(!verified) return false;
    return verified;
  } catch (error) {
    return false;
  }
};

const generateToken = (data) => {
  try {
    const token = jwt.sign(data, configs.jwt.key, {
      algorithm: 'HS256',
      expiresIn: configs.jwt.expiredIn,
    });
    return token;
  } catch (error) {
    return false;
  }
};

export default {
  verifyToken,
  generateToken
};
