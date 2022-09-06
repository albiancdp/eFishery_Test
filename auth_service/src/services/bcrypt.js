import bcrypt from 'bcryptjs';

const getHash = (password) => {
  try {
    const salt = bcrypt.genSaltSync(10);
    return bcrypt.hashSync(password, salt);
  } catch (error) {
    return error;
  }
};

const checkHash = (password, hash) => {
  try {
    return bcrypt.compareSync(password, hash);
  } catch (error) {
    return error;
  }
};

export default {
  getHash,
  checkHash
};
