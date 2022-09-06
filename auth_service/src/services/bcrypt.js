import bcrypt from 'bcryptjs';

const getHash = (password) => {
  return new Promise((resolve, reject) => {
    try {
      const salt = bcrypt.genSaltSync(10);
      resolve(bcrypt.hashSync(password, salt));
    } catch (error) {
      reject(error);
    }
  });
};

const checkHash = (password, hash) => {
  return new Promise((resolve, reject) => {
    try {
      resolve(bcrypt.compareSync(password, hash));
    } catch (error) {
      reject(error);
    }
  });
};

export default {
  getHash,
  checkHash
};
