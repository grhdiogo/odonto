import React from 'react';
import { BrowserRouter } from 'react-router-dom';
import { Flip, ToastContainer } from 'react-toastify';
import MainRouter from './routers/main';
import { environment } from './environments';

function App() {
  return (
    <BrowserRouter basename={environment.publicRoot}>
      <ToastContainer
        transition={Flip}
        theme={'colored'}
      />
      <MainRouter />
    </BrowserRouter>
  );
}
export default App;
