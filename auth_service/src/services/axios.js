import axios from 'axios';

const apiCall = async (method, link, data) => {
  return new Promise((resolve, reject) => {
    try {
      const configAxios = {
        url: link,
        method: method,
        headers: data.header,
        params: data.param,
        data: data.body
      };
      axios(configAxios)
        .then(response => {
          resolve(response.data);
        }).catch(error => {
          reject(error);
        });
    } catch (error) {
      reject(error);
    }
  });
};

export default apiCall;
