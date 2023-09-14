import { privClient } from './client';

export interface Patient {
  pid ?: string
  name: string
  email: string
  cpf: string
  birthdate: string
}

interface ListPatients {
  entities: Patient[]
  total: number
}

//
export default class PatientServices {
  token: string;

  constructor(token: string) {
    this.token = token;
  }

  list(text: string, page: number, limit: number): Promise<ListPatients> {
    // create path url
    const path = `/patients?text=${text}&page=${page}&limit=${limit}`;
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

  create(request: Patient): Promise<any> {
    // create path url
    const path = '/patient';
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

  update(id: string, request: Patient): Promise<any> {
    // create path url
    const path = `/patient/${id}`;
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
    const path = `/patient/${id}`;
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
