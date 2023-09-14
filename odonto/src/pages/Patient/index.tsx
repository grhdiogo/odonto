/* eslint-disable react/jsx-no-undef */
import moment from 'moment';
import { useEffect, useState } from 'react';
import { Col, Row } from 'react-bootstrap';
import { Form, InputField } from '../../components';
import ListEditor from '../../components/ListEditor';
import AlertDialog from '../../hooks/alert';
import PatientService, { Patient } from '../../services/patient';
import { getToken } from '../../services/cookies';
import handleNotify from '../../utils/notify';
import { ButtonContainer, StyledButton } from './styles';

const HEAD = [
  {
    columnKey: 'pid',
    columnValue: 'Identificador',
  },
  {
    columnKey: 'name',
    columnValue: 'Nome',
  },
  {
    columnKey: 'email',
    columnValue: 'E-mail',
  },
  {
    columnKey: 'cpf',
    columnValue: 'CPF',
  },
  {
    columnKey: 'birthdate',
    columnValue: 'Data de nascimento',
  },
];

const initialValues = {
  name: '',
  email: '',
  cpf: '',
  birthdate: '',
};

const LIMIT = 10;

export default function PatientPagePage() {
  // selected tab list or form
  const [selectedTab, setSelectedTab] = useState('list');
  const token = getToken();
  //
  const [showAlert, setShowAlert] = useState(false);
  const [deleteData, setDeleteData] = useState<any>();
  const [internalInitialFormValues, setInternalInitialFormValues] = useState(initialValues);
  const [searchText, setSearchText] = useState('');
  const [body, setBody] = useState<Patient[]>([]);
  const [page, setPage] = useState(1);
  const [totalEntities, setTotalEntities] = useState(0);

  function updateList(t: string, p: number) {
    const service = new PatientService(t);
    //
    service.list(searchText, p, LIMIT).then((v) => {
      setBody(v.entities);
      setTotalEntities(v.total);
    });
  }

  useEffect(() => {
    updateList(token, page);
  }, [token]);

  //
  function handleSelectTab(active: string) {
    setSelectedTab(active);
  }

  function handleSearchClick(text: string) {
    setSearchText(text);
  }

  function handleActionClick(act: string, data: any) {
    if (act === 'upt') {
      // alter values to form and change tab
      setInternalInitialFormValues({
        ...data,
        birthdate: moment(data.birthdate, 'DD/MM/YYYY').format('YYYY-MM-DD'),
      });
      setSelectedTab('form');
    }
    //
    if (act === 'del') {
      setShowAlert(true);
      setDeleteData(data);
    }
  }

  function reset() {
    // change the tab to list
    setSelectedTab('list');
    setInternalInitialFormValues(initialValues);
    setPage(1);
    updateList(token, 1);
  }

  function handleChangePage(p: number) {
    setPage(p);
    updateList(token, p);
  }

  function handleDelete(data: any) {
    setShowAlert(false);
    const service = new PatientService(token);

    service.delete(data.pid).then(() => {
      handleNotify('success', 'Sucesso');
      reset();
    }).catch((e) => {
      handleNotify('error', `Falha ${e.response.data.message}`);
    });
  }

  function create(data: any) {
    const service = new PatientService(token);
    service.create({
      name: data.name,
      email: data.email,
      cpf: String(data.cpf),
      birthdate: moment(data.birthdate).format('DD/MM/YYYY'),
    }).then(() => {
      handleNotify('success', 'Sucesso');
      reset();
    }).catch((e) => {
      handleNotify('error', `Falha ${e.response.data.message}`);
    });
  }

  function update(id: any, data: any) {
    const service = new PatientService(token);
    service.update(id, {
      name: data.name,
      email: data.email,
      cpf: String(data.cpf),
      birthdate: moment(data.birthdate).format('DD/MM/YYYY'),
    }).then(() => {
      handleNotify('success', 'Sucesso');
      reset();
    }).catch((e) => {
      handleNotify('error', `Falha ${e.response.data.message}`);
    });
  }

  function handleFormSubmit(data: any) {
    const id = data.pid;
    if (id) {
      // update
      update(id, data);
    } else {
      // create
      create(data);
    }
  }

  return (
    <div>
      <ListEditor
        pageTitle={'Paciente'}
        data={body}
        head={HEAD}
        selectedTab={selectedTab}
        actions={[
          { action: 'upt', name: 'Atualizar', variant: '' },
          { action: 'del', name: 'Remover', variant: 'danger' },
        ]}
        formInitialValues={internalInitialFormValues}
        onSearchClick={(text: string) => handleSearchClick(text)}
        onActionClick={(action, dataItem) => handleActionClick(action, dataItem)}
        onSelectTab={(active: string) => handleSelectTab(active)}
        loading={false}
        onSaveClick={(data: any) => handleFormSubmit(data)}
        onPageChange={(v: number) => handleChangePage(v)}
        paginationLimit={LIMIT}
        paginationTotal={totalEntities}
      >
        {({
          loading,
          values,
          errors,
          touched,
          handleChange,
          handleBlur,
          handleReset,
          handleSubmit,
        }: any) => (
          <Form
            className="mt-3"
            onReset={handleReset}
            onSubmit={handleSubmit}
          >
            <Row>
              <InputField
                isRequired
                name={'name'}
                placeholder={'Nome'}
                label={'Nome'}
                value={values.name}
                size={8}
                onChange={handleChange}
                onBlur={handleBlur}
                errorMsg={errors.name && touched.name ? errors.name : null}
                disabled={loading}
              />
              <InputField
                type={'date'}
                isRequired
                name={'birthdate'}
                placeholder={'Data de nascimento'}
                label={'Data de nascimento'}
                value={values.birthdate}
                size={4}
                onChange={handleChange}
                onBlur={handleBlur}
                errorMsg={errors.name && touched.name ? errors.name : null}
                disabled={loading}
              />
            </Row>
            <Row>
              <InputField
                isRequired
                name={'email'}
                placeholder={'Email'}
                label={'Email'}
                value={values.email}
                size={7}
                onChange={handleChange}
                onBlur={handleBlur}
                errorMsg={errors.name && touched.name ? errors.name : null}
                disabled={loading}
              />
              <InputField
                isRequired
                name={'cpf'}
                type={'numeric'}
                pattern="[0-9]*"
                placeholder={'Cpf'}
                label={'Cpf (apenas números)'}
                value={values.cpf}
                size={5}
                onChange={handleChange}
                onBlur={handleBlur}
                errorMsg={errors.name && touched.name ? errors.name : null}
                disabled={loading}
              />
            </Row>
            <Row className="justify-content-md-center mt-3">
              <Col xl={8} />
              <Row sm={6} className="d-flex justify-content-end">
                <ButtonContainer>
                  <StyledButton
                    label={'Cancelar'}
                    kind={'primary'}
                    onClick={() => reset()}
                  />
                </ButtonContainer>
                <ButtonContainer>
                  <StyledButton
                    label={'Salvar'}
                    kind={'primary'}
                    isLoading={loading}
                    type='submit'
                  />
                </ButtonContainer>
              </Row>
            </Row>
          </Form>
        )}
      </ListEditor>
      <AlertDialog
        title='Remover'
        text='Remover paciente'
        show={showAlert}
        opts={[
          {
            title: 'Sim',
            do: () => handleDelete(deleteData),
          },
        ]}
        onClose={() => setShowAlert(false)}
      />
    </div>
  );
}
