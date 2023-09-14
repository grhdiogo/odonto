import { privClient } from './client';

export interface Doctor {
  pid ?: string
  name: string
  email: string
  cpf: string
  birthdate: string
}

interface ListDoctors {
  entities: Doctor[]
  total: number
}

//
export default class DoctorServices {
  token: string;

  constructor(token: string) {
    this.token = token;
  }

  list(text: string, page: number, limit: number): Promise<ListDoctors> {
    // create path url
    const path = `/doctors?text=${text}&page=${page}&limit=${limit}`;
    // create get request
    return new Promise((resolve, reject) => {
      // create a client for private routes
      const client = privClient(this.token);
      // request Data
      client.get(path).then((resp) => {
        const { data } = resp;
        resolve(data);
      }).catch((e) => {
        reject(e);
      });
    });
  }

  create(request: Doctor): Promise<any> {
    // create path url
    const path = '/doctor';
    // create get request
    return new Promise((resolve, reject) => {
      // create a client for private routes
      const client = privClient(this.token);
      // request Data
      client.post(path, request).then((resp) => {
        const { data } = resp;
        resolve({ data });
      }).catch((e) => {
        reject(e);
      });
    });
  }

  update(id: string, request: Doctor): Promise<any> {
    // create path url
    const path = `/doctor/${id}`;
    // create get request
    return new Promise((resolve, reject) => {
      // create a client for private routes
      const client = privClient(this.token);
      // request Data
      client.put(path, request).then((resp) => {
        const { data } = resp;
        resolve({ data });
      }).catch((e) => {
        reject(e);
      });
    });
  }

  delete(id: string): Promise<any> {
    // create path url
    const path = `/doctor/${id}`;
    // create get request
    return new Promise((resolve, reject) => {
      // create a client for private routes
      const client = privClient(this.token);
      // request Data
      client.delete(path).then((resp) => {
        const { data } = resp;
        resolve({ data });
      }).catch((e) => {
        reject(e);
      });
    });
  }
}
