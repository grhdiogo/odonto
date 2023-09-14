import React, { useRef } from 'react';
import { Card, Tabs, Tab } from 'react-bootstrap';
import FormPanel from './form';
import ListPanel from './table';
// import styles from './styles.module.css';

interface Props {
  disableList?: boolean,
  disableEditor?: boolean,
  pageTitle: string,
  loading: boolean,
  head: any[],
  data: any[],
  actions?: {
    action: string,
    name: string,
    variant: string
  }[],
  selectedTab?: string,
  formInitialValues?: any,
  className?: string,
  formValidationSchema?: any,
  onSaveClick?: any,
  onCancelClick?: any,
  onSearchClick?: any,
  onActionClick?: (action: string, data: any) => void,
  onSelectTab?: (tab: string) => void,
  children?: any,
  ref?: any,
  onPageChange ?:any
  paginationLimit ?: number
  paginationTotal ?: number
}

function ListEditor({
  /**
   * disable list
   */
  disableList = false,
  /**
   * disable editor
   */
  disableEditor = false,
  /**
   * page title
   */
  pageTitle = 'Page title',
  // setSelectedTabKey
  /**
   * in loading
   */
  loading = false,
  /**
   * table head
   * - [{ columnKey, columnValue }, ...]
   */
  head = [],
  /**
   * generic data (convertible for rows)
   * - row = [{columnKey, columnValue}]
   * - rows = [row1, row2, ...]
   */
  data = [],
  /**
   * actions
   * - [{action, name, variant}]
   * - action: is key
   * - name: used for display button
   * - variant: see in 'bootstrap button variant'
   */
  actions = [],
  /**
   * selected tab
   */
  selectedTab = 'list',
  /**
   * style class name
   */
  className,
  /**
   * initial data
   */
  formInitialValues,
  /**
   * esquema de validação yup para o formik
   */
  formValidationSchema,
  /**
   * propagate save button click in form tab
   * - onSaveClick()
   */
  onSaveClick,
  /**
   * propagate cancel button click in form tab
   * - onCancelClick()
   */
  onCancelClick,
  /**
   * propagate search button click in table tab
   * - onSearchClick(filter-text)
   */
  onSearchClick,
  /**
   * propagate action button click in each row of table tab
   * - onActionClick(action, data[rowIndex])
   */
  onActionClick,
  /**
   * propagate selection tab event
   */
  onSelectTab,
  /**
   * children
   */
  children,
  //
  onPageChange,
  paginationLimit = 0,
  paginationTotal = 0,
  ref,
}: Props) {
  // ref
  const formRef = useRef();

  /**
   * handle ref actions
   */
  // useImperativeHandle(ref, () => ({
  //   resetForm: () => {
  //     if (formRef && formRef.current) {
  //       formRef.current.resetForm();
  //     }
  //   },
  // }), [formRef]);

  /**
   * handle action click in table row
   */
  function handleActionClick(action: string, index: number) {
    if (onActionClick) onActionClick(action, data[index]);
  }

  function handleSelectTab(k: string) {
    if (onSelectTab) onSelectTab(k);
  }

  return (
    <Card className={className}>
      <Card.Header as="h5">{pageTitle}</Card.Header>
      <Card.Body>
        <Tabs
          className={'cardTabs'}
          defaultActiveKey={disableList ? 'form' : 'list'}
          activeKey={selectedTab}
          onSelect={(k) => handleSelectTab(k || 'list')}
          transition={false}
        >
          {!disableList ? (
            <Tab eventKey="list" title="Listagem" className={'cardTabsContent'}>
              <ListPanel
                loading={loading}
                head={head}
                data={data}
                actions={actions}
                onSearchClick={onSearchClick}
                onActionClick={(action: string, index: number) => handleActionClick(action, index)}
                onPageChange={onPageChange}
                paginationLimit={paginationLimit}
                paginationTotal={paginationTotal}
              />
            </Tab>
          ) : null}
          {!disableEditor ? (
            <Tab eventKey="form" title="Cadastro" className={'cardTabsContent'}>
              <FormPanel
                ref={formRef}
                loading={loading}
                initialValues={formInitialValues}
                validationSchema={formValidationSchema}
                onSaveClick={onSaveClick}
                onCancelClick={onCancelClick}
              >
                {children}
              </FormPanel>
            </Tab>
          ) : null}
        </Tabs>
      </Card.Body>
    </Card>
  );
}

export default React.forwardRef(ListEditor);
