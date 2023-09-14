import { privClient } from './client';

export interface Items {
  Name: string
  Value: number
  Tooth: number
  ID: string
}

export interface Appointment {
  aid ?: string
  status?: string
  observation: string
  doctorDid: string
  patientPid: string
  items: Items[]
}

export interface AppointmentProxy {
  aid ?: string
  status: string
  observation: string
  doctorDid: string
  patientPid: string
  patientName: string
  doctorName: string
  items: Items[]
}

interface ListProcedures {
  entities: AppointmentProxy[]
  total: number
}

export interface CreateOrUpdateAppointment {
  aid ?: string
  observation: string
  doctorDid: string
  patientPid: string
  items: Items[]
}

//
export default class AppointmentServices {
  token: string;

  constructor(token: string) {
    this.token = token;
  }

  list(text: string, page: number, limit: number): Promise<ListProcedures> {
    // create path url
    const path = `/appointments?text=${text}&page=${page}&limit=${limit}`;
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

  create(request: CreateOrUpdateAppointment): Promise<any> {
    // create path url
    const path = '/appointment';
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

  update(id: string, request: CreateOrUpdateAppointment): Promise<any> {
    // create path url
    const path = `/appointment/${id}`;
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
    const path = `/appointment/${id}`;
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
