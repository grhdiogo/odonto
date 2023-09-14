import { toast } from 'react-toastify';
import 'react-toastify/dist/ReactToastify.css';

const notify: any = {
  error(message: string) {
    toast.error(message, {
      position: toast.POSITION.BOTTOM_RIGHT,
    });
  },
  success(message: string) {
    toast.success(message, {
      position: toast.POSITION.BOTTOM_RIGHT,
    });
  },
  warning(message: string) {
    toast.warning(message, {
      position: toast.POSITION.BOTTOM_RIGHT,
    });
  },
  info(message: string) {
    toast.info(message, {
      position: toast.POSITION.BOTTOM_RIGHT,
    });
  },
  default(message: string) {
    toast(message, {
      position: toast.POSITION.BOTTOM_RIGHT,
    });
  },
};

export default function handleNotify(messageType: string, message: string) {
  // get function according to specified notification type
  const showNotify = notify[messageType];
  // execute the function
  showNotify(message);
}
