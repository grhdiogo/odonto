import { pubClient, privClient } from './client';

interface CheckTokenResponse {
  isValid: boolean
}

// Client for authentication service
export default class AuthServices {
  login(email: string, password: string): Promise<any> {
    // create path url
    const path = '/auth';
    // create get request
    return new Promise((resolve, reject) => {
      // create a client for private routes
      const client = pubClient();
      // request Data
      const reqData = {
        identifier: email,
        secret: password,
      };
      client.post(path, reqData).then((resp) => {
        const { data } = resp;
        resolve(data);
      }).catch((e) => {
        reject(e);
      });
    });
  }

  checkToken(token: string): Promise<CheckTokenResponse> {
    // create path url
    const path = '/check-token';
    // create get request
    return new Promise((resolve) => {
      // create a client for private routes
      const client = privClient(token);
      client.get(path).then((resp) => {
        const { data } = resp;
        resolve({ isValid: data.isValid });
      }).catch(() => {
        resolve({ isValid: false });
      });
    });
  }

  create(name: string, age: string, email: string, password: string): Promise<any> {
    // create path url
    const path = '/';
    // create get request
    return new Promise((resolve, reject) => {
      // create a client for private routes
      const client = pubClient();
      // request Data
      const reqData = {
        name,
        age,
        email,
        password,
      };
      client.post(path, { ...reqData }).then((resp) => {
        const { data } = resp;
        resolve({ data });
      }).catch((e) => {
        reject(e);
      });
    });
  }
}
