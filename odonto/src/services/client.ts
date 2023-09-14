import axios from 'axios';

// create axios api
const pubClient = () => axios.create({
  baseURL: '/pub/v1',
});

const privClient = (token: string) => axios.create({
  baseURL: '/priv/v1',
  headers: {
    authorization: `Bearer ${token}`,
  },
});

export { pubClient, privClient };
