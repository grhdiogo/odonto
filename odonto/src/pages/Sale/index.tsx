/* eslint-disable react/jsx-no-useless-fragment */
/* eslint-disable max-len */
/* eslint-disable react/jsx-closing-bracket-location */
import { useState, useEffect } from 'react';
import { Col } from 'react-bootstrap';
import ListEditor from '../../components/ListEditor';
import { Form, TextArea, SelectField } from '../../components';
import { WithOptions } from '../../components/Tooth';
import { numberToMoney } from '../../utils/money';
import {
  Container, ToothContainer, ToothsContainer,
  Label, ResumeContainer, Column, ResumeText, ResumeTitle, FormContainer, Row,
  ButtonContainer, StyledButton,
} from './style';
import { DOWN_TOOTHS, UP_TOOTHS } from './tooth';
import PatientService, { Patient } from '../../services/patient';
import { getToken } from '../../services/cookies';
import DoctorServices, { Doctor } from '../../services/doctor';
import handleNotify from '../../utils/notify';
import AppointmentServices, { Appointment } from '../../services/appointment';
import ProcedureServices from '../../services/procedure';
import AlertDialog from '../../hooks/alert';

const HEAD = [
  {
    columnKey: 'patientName',
    columnValue: 'Paciente',
  },
  {
    columnKey: 'doctorName',
    columnValue: 'Doutor',
  },
  {
    columnKey: 'date',
    columnValue: 'Data',
  },
  {
    columnKey: 'value',
    columnValue: 'Valor',
  },
  {
    columnKey: 'status',
    columnValue: 'Situação',
  },
];

interface OPT {
  id: string
  label: string
  value: number
}

interface MarkedOpts {
  toothID: number
  id: string
}

const initialValues = {
  observation: '',
  doctor: '',
  patient: ''
};

interface AppointmentList {
  aid ?: string
  status?: string
  observation: string
  doctorDid: string
  patientPid: string
  value: number
}

const StatusMap = new Map([
  ["pending", "Pendente"]
])

export default function SalePage() {
  const token = getToken();
  // selected tab list or form
  const [selectedTab, setSelectedTab] = useState('list');
  const [markedOptions, setMarkedOptions] = useState<MarkedOpts[]>([]);
  const [internalInitialFormValues, setInternalInitialFormValues] = useState(initialValues);
  const [patients, setPatients] = useState<Patient[]>([]);
  const [doctors, setDoctors] = useState<Doctor[]>([]);
  const [appointments, setAppointments] = useState<AppointmentList[]>([]);
  const [opts, setOpts] = useState<OPT[]>([]);
  const [showAlert, setShowAlert] = useState(false);
  const [deleteData, setDeleteData] = useState<any>();

  function listPatients(t: string, p: number) {
    const service = new PatientService(t);
    //
    service.list('', p, -1).then((v) => {
      setPatients(v.entities);
    });
  }

  function listDoctors(t: string, p: number) {
    const service = new DoctorServices(t);
    //
    service.list('', p, -1).then((v) => {
      setDoctors(v.entities);
    });
  }

  function listAppointments(t: string, p: number) {
    const service = new AppointmentServices(t);
    //
    service.list('', p, -1).then((v) => {
      setAppointments(v.entities.map((v) => ({
        ...v,
        status: StatusMap.get(v.status),
        value: v.items.reduce(((tot, curr) => curr.Value + tot), 0)
      })));
    });
  }

  function listProcedures(t: string) {
    const service = new ProcedureServices(t);
    //
    service.list('', 1, -1).then((v) => {
      setOpts(v.entities.map((v) => ({
        id: v.pid as string,
        label: v.name,
        value: v.value
      })));
    });
  }

  function list(t: string) {
    listDoctors(t, 1);
    listPatients(t, 1);
    listProcedures(t);
    listAppointments(t, 1);
  }

  useEffect(() => {
    list(token);
  }, [token]);

  function handleActionClick(act: string, data: Appointment) {
    if (act === 'upt') {
      // alter values to form and change tab
      setInternalInitialFormValues({
        doctor: data.doctorDid,
        patient: data.patientPid,
        observation: data.observation
      });
      setMarkedOptions(data.items.map((v) => ({
        id: v.ID,
        toothID: v.Tooth
      })))
      setSelectedTab('form');
    }
    //
    if (act === 'del') {
      setShowAlert(true);
      setDeleteData(data);
    }
  }
  //
  // handle mark option from tooth
  function handleMarkOpt(toothID: number, id: string) {
    const updated = [...markedOptions];
    let i = -1;
    // iterate throu marked verify if exists
    for (let index = 0; index < markedOptions.length; index += 1) {
      const el = markedOptions[index];
      // case found
      if (el.id === id && el.toothID === toothID) {
        i = index;
        break;
      }
    }
    // if exists on array, remove
    if (i !== -1) {
      updated.splice(i, 1);
    } else {
      // if not exists, add
      updated.push({
        id,
        toothID,
      });
    }
    setMarkedOptions(updated);
  }

  function reset() {
    // change the tab to list
    setSelectedTab('list');
    setInternalInitialFormValues(initialValues);
    list(token);
  }

  function handleDelete(data: any) {
    setShowAlert(false);
    const service = new AppointmentServices(token);

    service.delete(data.aid).then(() => {
      handleNotify('success', 'Sucesso');
      reset();
    }).catch((e) => {
      handleNotify('error', `Falha ${e.response.data.message}`);
    });
  }

  function handleSelectTab(active: string) {
    setSelectedTab(active);
  }

  function create(data: any) {
    const service = new AppointmentServices(token);
    service.create({
      doctorDid: data.doctor,
      observation: data.observation,
      items: markedOptions.map((v) => ({
        ID: v.id,
        Name: opts.find((opt) => opt.id === v.id)?.label || '',
        Tooth: v.toothID,
        Value: opts.find((opt) => opt.id === v.id)?.value || 0
      })),
      patientPid: data.patient,
    }).then(() => {
      handleNotify('success', 'Sucesso');
      reset();
    }).catch((e) => {
      handleNotify('error', `Falha ${e.response.data.message}`);
    });
  }

  function update(id: any, data: any) {
    const service = new AppointmentServices(token);
    service.update(id, {
      doctorDid: data.doctor,
      observation: data.observation,
      items: markedOptions.map((v) => ({
        ID: v.id,
        Name: opts.find((opt) => opt.id === v.id)?.label || '',
        Tooth: v.toothID,
        Value: opts.find((opt) => opt.id === v.id)?.value || 0
      })),
      patientPid: data.patient,
    }).then(() => {
      handleNotify('success', 'Sucesso');
      reset();
    }).catch((e) => {
      handleNotify('error', `Falha ${e.response.data.message}`);
    });
  }

  function handleFormSubmit(data: any) {
    const id = data.aid;
    if (id) {
      // update
      update(id, data);
    } else {
      // create
      create(data);
    }
  }
  //
  return (
    <div>
    <Container>
      <ListEditor
        pageTitle={'Venda'}
        data={appointments}
        head={HEAD}
        selectedTab={selectedTab}
        actions={[
          { action: 'upt', name: 'Atualizar', variant: '' },
          { action: 'del', name: 'Remover', variant: 'danger' },
        ]}
        // onSearchClick={(text: string) => handleSearchClick(text)}
        onActionClick={(action, dataItem) => handleActionClick(action, dataItem)}
        onSelectTab={(active: string) => handleSelectTab(active)}
        loading={false}
        formInitialValues={internalInitialFormValues}
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
            <FormContainer>
              <Column>
                <Row>
                  <SelectField
                    onChangeText={() => null}
                    name={'patient'}
                    data={[
                      {
                        value: '',
                        text: 'Selecione um paciente',
                      },
                      ...patients.map((v) => ({
                        text: v.name,
                        value: v.pid,
                      })),
                    ]}
                    initialValue={values.patient}
                    placeholder={'Paciente'}
                    label={'Paciente'}
                    value={values.patient}
                    size={6}
                    onChange={handleChange}
                    onBlur={handleBlur}
                    errorMsg={errors.patient && touched.patient ? errors.patient : null}
                    disabled={loading}
                  />
                  <SelectField
                    onChangeText={() => null}
                    name={'doctor'}
                    data={[
                      {
                        value: '',
                        text: 'Selecione um médico',
                      },
                      ...doctors.map((v) => ({
                        text: v.name,
                        value: v.pid,
                      })),
                    ]}
                    initialValue={values.doctor}
                    placeholder={'Médico'}
                    label={'Médico'}
                    value={values.doctor}
                    size={6}
                    onChange={handleChange}
                    onBlur={handleBlur}
                    errorMsg={errors.doctor && touched.doctor ? errors.doctor : null}
                    disabled={loading}
                  />
                </Row>
                <Row>
                  <Column>
                    <ToothsContainer>
                      {UP_TOOTHS.map((v) => (
                        <ToothContainer key={v.id}>
                          <Label>{v.id}</Label>
                          <WithOptions
                            icon={v.icon}
                            options={opts.map((o) => ({
                              ...o,
                              marked: markedOptions.some((m) => m.toothID === v.id && o.id === m.id),
                            }))}
                            onOptionClick={(id) => handleMarkOpt(v.id, id)}
                            selected={markedOptions.some((m) => m.toothID === v.id)}
                          />
                        </ToothContainer>
                      ))}
                    </ToothsContainer>
                    <ToothsContainer>
                      {DOWN_TOOTHS.map((v) => (
                        <ToothContainer key={v.id}>
                          <WithOptions
                            icon={v.icon}
                            options={opts.map((o) => ({
                              ...o,
                              marked: markedOptions.some((m) => m.toothID === v.id && o.id === m.id),
                            }))}
                            onOptionClick={(id) => handleMarkOpt(v.id, id)}
                            selected={markedOptions.some((m) => m.toothID === v.id)}
                          />
                          <Label>{v.id}</Label>
                        </ToothContainer>
                      ))}
                    </ToothsContainer>
                  </Column>
                  <ResumeContainer>
                    <h3>
                      Total: {
                        numberToMoney(markedOptions.reduce((acc, curr) => acc + (opts.find((v) => v.id === curr.id)?.value || 0), 0))
                      }
                    </h3>
                    <br />
                    {opts.filter((v) => markedOptions.some((v1) => v1.id === v.id)).map((v) => (
                      <>
                        <ResumeTitle key={v.id}>
                          {v.label} - {numberToMoney(v.value)}
                        </ResumeTitle>
                        {markedOptions.filter((v1) => v1.id === v.id).map((v2) => (
                          <ResumeText key={`${v2.id}-${v2.toothID}`}>
                            Dente: {v2.toothID}
                          </ResumeText>
                        ))}
                      </>
                    ))}
                    <br />
                  </ResumeContainer>
                </Row>
                <Row>
                  <TextArea
                    onChangeText={() => null}
                    name={'observation'}
                    placeholder={'Observações'}
                    label={'Observações'}
                    value={values.observation}
                    size={12}
                    onChange={handleChange}
                    onBlur={handleBlur}
                    errorMsg={errors.observation && touched.observation ? errors.observation : null}
                    disabled={loading}
                  />
                </Row>
                <Row className="justify-content-md-center mt-3">
                  <Col xl={8} />
                  <Row className="d-flex justify-content-end">
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
              </Column>
            </FormContainer>
          </Form>
        )}
      </ListEditor>
    </Container>
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
