/* eslint-disable react/prop-types */
import React, { useRef, useImperativeHandle } from 'react';
import { Formik } from 'formik';

function FormPanel({
  /**
   * in loading
   */
  loading = false,
  /**
   * initial data
   */
  initialValues,
  /**
   * esquema de validação yup para o formik
   */
  validationSchema,
  /**
   * propaga o evento do resultado de salvar (apenas sucesso)
   * - parametro: lista de usuários
   */
  onSaveClick,
  /**
   * propaga o evento do resultado de salvar (apenas sucesso)
   * - parametro: nenhum
   */
  onCancelClick,
  /**
   * children
   */
  children,
}, ref) {
  // formik reference
  const formikRef = useRef();

  /**
   * handle ref actions
   */
  useImperativeHandle(ref, () => ({
    resetForm: () => {
      if (formikRef && formikRef.current) {
        formikRef.current.resetForm();
      }
    },
  }), [formikRef]);

  /**
   * handle save click
   */
  function handleSaveSubmit(data) {
    if (onSaveClick) {
      onSaveClick(data);
    }
  }

  /**
   * handle reset click
   */
  function handleReset() {
    if (onCancelClick) {
      onCancelClick();
    }
  }

  return (
    <Formik
      enableReinitialize
      innerRef={formikRef}
      initialValues={initialValues}
      validationSchema={validationSchema}
      onSubmit={(data) => handleSaveSubmit(data)}
      onReset={() => handleReset()}
    >
      {(formikProps) => children({ ...formikProps, loading })}
    </Formik>
  );
}

export default React.forwardRef(FormPanel);
