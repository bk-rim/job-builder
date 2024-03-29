import { toast } from 'react-toastify';


export const notify = (msg: string) => {
    toast.success(msg, {
      position: "bottom-right",
      autoClose: 1000,
      hideProgressBar: true,
      closeOnClick: true,
      pauseOnHover: true,
      draggable: true,
      theme: "light",
    });
  }