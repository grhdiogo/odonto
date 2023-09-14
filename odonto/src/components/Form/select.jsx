/* eslint-disable react/prop-types */
import React from 'react';
import { Field } from 'formik';
import styles from './styles.module.css';

export default function SelectField({
  // nome do campo
  name,
  // etiqueta do campo
  label,
  // dica do campo
  placeholder,
  // tamanho do campo (valores entre 1 e 12)
  size = 12,
  // valor inicial
  initialValue,
  // options = [{value, text}]
  data = [],
  // verifica campos requeridos
  isRequired = false,
  // tipo do input
  type = 'text',
  // propaga o evento de alteração do texto
  onChange,
  // propaga o evento de alteração do texto
  onChangeText,
  // propaga o evento de onblur
  onBlur,
  // outras propriedades
  ...props
}) {
  /**
   * Manipula o evento de alteração do texto
   * @param {*} event DOM Event
   */
  function handleChange(event) {
    if (onChange) {
      onChange(event);
    }
    if (onChangeText) {
      const textValue = event.target.value;
      onChangeText(textValue);
    }
  }

  function handleBlur(event) {
    if (onBlur) {
      onBlur(event);
    }
  }
  return (
    <Field {...props} name={name}>
      {({ meta }) => (
        <div className={`form-group col-${size} ${styles.formGroup}`}>
          {/* label */}
          <h6 htmlFor={name}>
            {label}
            <small>{isRequired ? '*' : null}</small>
          </h6>
          {/* input (select) */}
          <select
            type={type}
            className="form-control"
            id={name}
            placeholder={placeholder}
            required={isRequired}
            onChange={handleChange}
            onBlur={handleBlur}
          >
            {data && data.map((dat) => (
              <option value={dat.value} selected={dat.value === initialValue}>{dat.text}</option>
            ))}
          </select>
          {/* error */}
          {meta.touched && meta.error ? (
            <small className="text-danger">{meta.error}</small>
          ) : null}
        </div>
      )}
    </Field>
  );
}
