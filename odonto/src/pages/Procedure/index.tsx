/* eslint-disable react/jsx-no-undef */
import moment from 'moment';
import { useEffect, useState } from 'react';
import { Col, Row } from 'react-bootstrap';
import { Form, InputField } from '../../components';
import ListEditor from '../../components/ListEditor';
import AlertDialog from '../../hooks/alert';
import handleNotify from '../../utils/notify';
import ProcedureService, { Procedure } from '../../services/procedure';
import { getToken } from '../../services/cookies';
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
    columnKey: 'value',
    columnValue: 'Valor',
  },
];

const initialValues = {
  name: '',
  value: '',
};

export default function ProcedurePage() {
  // selected tab list or form
  const [selectedTab, setSelectedTab] = useState('list');
  const token = getToken();
  //
  const [showAlert, setShowAlert] = useState(false);
  const [deleteData, setDeleteData] = useState<any>();
  const [internalInitialFormValues, setInternalInitialFormValues] = useState(initialValues);
  const [searchText, setSearchText] = useState('');
  const [body, setBody] = useState<Procedure[]>([]);

  function updateList(t: string) {
    const service = new ProcedureService(t);
    //
    service.list(searchText, 1, 10).then((v) => {
      setBody(v.entities);
    });
  }

  useEffect(() => {
    updateList(token);
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
    updateList(token);
  }

  function handleDelete(data: any) {
    setShowAlert(false);
    const service = new ProcedureService(token);

    service.delete(data.pid).then(() => {
      handleNotify('success', 'Sucesso');
      reset();
    }).catch((e) => {
      handleNotify('error', `Falha ${e.response.data.message}`);
    });
  }

  function create(data: any) {
    const service = new ProcedureService(token);
    service.create({
      name: data.name,
      value: data.value,
      description: data.description,
    }).then(() => {
      handleNotify('success', 'Sucesso');
      reset();
    }).catch((e) => {
      handleNotify('error', `Falha ${e.response.data.message}`);
    });
  }

  function update(id: any, data: any) {
    const service = new ProcedureService(token);
    service.update(id, {
      name: data.name,
      value: data.value,
      description: data.description,
    }).then(() => {
      handleNotify('success', 'Sucesso');
      reset();
    }).catch((e) => {
      handleNotify('error', `Falha ${e.response.data.message}`);
    });
  }

  function handleFormSubmit(data: any) {
    console.log('aqui');
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
        pageTitle={'Procedimento'}
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
                onChangeText={() => null}
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
                onChangeText={() => null}
                type={'number'}
                isRequired
                name={'value'}
                placeholder={'Valor'}
                label={'Valor'}
                value={values.value}
                size={4}
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
        text='Remover procedimento'
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
