/* eslint-disable react/no-unstable-nested-components */
import {
  Routes,
  Route,
  // useNavigate,
} from 'react-router-dom';
import SalePage from '../pages/Sale';
import DoctorPage from '../pages/Doctor';
import PatientPage from '../pages/Patient';
import ProcedurePage from '../pages/Procedure';
import { MainLayout } from '../layout';

export default function Main() {
  // const history = useNavigate();

  function RenderPage(page: JSX.Element): JSX.Element {
    return (
      <MainLayout
        asideData={[
          {
            label: 'Doutor',
            redirectTo: '/doctor',
          },
          {
            label: 'Paciente',
            redirectTo: '/patient',
          },
          {
            label: 'Procedimento',
            redirectTo: '/procedure',
          },
          {
            label: 'Consulta',
            redirectTo: '/sale',
          },
        ]}
      >
        {page}
      </MainLayout>
    );
  }

  return (
    <Routes>
      <Route index element={RenderPage(<DoctorPage />)} />
      <Route path="/doctor" element={RenderPage(<DoctorPage />)} />
      <Route path="/sale" element={RenderPage(<SalePage />)} />
      <Route path="/patient" element={RenderPage(<PatientPage />)} />
      <Route path="/procedure" element={RenderPage(<ProcedurePage />)} />
    </Routes>
  );
}
