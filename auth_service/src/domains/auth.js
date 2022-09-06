import db from '../services/sqlite';
import bcrypt from '../services/bcrypt';

const register = (data) => {
  return new Promise((resolve, reject) => {
    try {
      let notValid = {};
      const password = bcrypt.getHash(data.password);
      const query = 'INSERT INTO user (name, phone, password, role) VALUES (?,?,?,?);';
      const value = [data.name, data.phone, password, data.role];
      db.run(query, value, (err, result) => {
        if (err) {
          //get error unique
          const errorMessage = err.message;
          const errorValid = errorMessage.split('failed: user.')[1];
          notValid[errorValid] = `${errorValid} is already exist`;
          resolve({ notValid: notValid });
        } else {
          resolve(data);
        }
      });
    } catch (error) {
      reject(false);
    }
  });
};

const login = (data) => {
  return new Promise((resolve, reject) => {
    try {
      const query = 'SELECT name, phone, role, password FROM user WHERE phone = ?;';
      const value = [data.phone];
      db.run(query, value, (err, result) => {
        if (err) {
          reject(err);
        } else {
          resolve(result);
        }
      });
    } catch (error) {
      reject(false);
    }
  });
};

export default {
  register,
  login
};
