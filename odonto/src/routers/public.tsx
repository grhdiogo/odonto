/* eslint-disable react/no-unstable-nested-components */
import {
  Routes,
  Route,
  useNavigate,
} from 'react-router-dom';
import Login from '../pages/Login';
import Register from '../pages/Register';

interface Props {
  onLogin: (token: string) => void
}

export default function PublicRouter({
  onLogin,
}: Props) {
  const history = useNavigate();

  function handleRegisterPress() {
    history('/register');
  }

  function LoginPage() {
    return (
      <Login
        onRegisterPress={() => handleRegisterPress()}
        onRedirectToDashBoardPage={(token) => onLogin(token)}
      />
    );
  }

  return (
    <Routes>
      <Route index element={LoginPage()} />
      <Route path="/register" element={<Register />} />
    </Routes>
  );
}
