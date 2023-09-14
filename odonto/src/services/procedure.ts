import { privClient } from './client';

export interface Procedure {
  pid ?: string
  name: string
  description: string
  value: number
}

interface ListProcedures {
  entities: Procedure[]
  total: number
}

//
export default class ProcedureServices {
  token: string;

  constructor(token: string) {
    this.token = token;
  }

  list(text: string, page: number, limit: number): Promise<ListProcedures> {
    // create path url
    const path = `/procedures?text=${text}&page=${page}&limit=${limit}`;
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

  create(request: Procedure): Promise<any> {
    // create path url
    const path = '/procedure';
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

  update(id: string, request: Procedure): Promise<any> {
    // create path url
    const path = `/procedure/${id}`;
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
    const path = `/procedure/${id}`;
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
