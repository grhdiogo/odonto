import { environment as environmentProd } from './production';
import { environment as environmentDev } from './development';

// select environment
let selectedEnvironment = environmentDev;
if (process.env.NODE_ENV === 'production') {
  selectedEnvironment = environmentProd;
}
const environment = selectedEnvironment;

export {
  environment,
};
