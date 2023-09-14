import React from 'react';
import { Field } from 'formik';
import { Container, Label, Input } from './styles';

export default function FormField({
  // nome do campo
  name,
  // etiqueta do campo
  label,
  // dica do campo
  placeholder,
  // tamanho do campo (valores entre 1 e 12)
  size = 12,
  // texto do input (entrada)
  value,
  // verifica campos requeridos
  isRequired = false,
  // tipo do input
  type = 'text',
  // está desabilitado
  disabled = false,
  // propaga o evento de alteração do texto
  onChange,
  // propaga o evento de onblur
  onBlur,
  // others props
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
    if (props.onChangeText) {
      const textValue = event.target.value;
      props.onChangeText(textValue);
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
        <Container className={`form-group col-${size}`}>
          {/* <h6 htmlFor={name}>
            {label}
            <small>{isRequired ? '*' : null}</small>
          </h6> */}
          <Label>
            {label}
          </Label>
          <Input
            type={type}
            inputmode="numeric"
            pattern={props.pattern}
            className="form-control"
            id={name}
            placeholder={placeholder}
            value={value}
            required={isRequired}
            onChange={handleChange}
            onBlur={handleBlur}
            disabled={disabled}
          />
          {meta.touched && meta.error ? (
            <small className="text-danger">{meta.error}</small>
          ) : null}
        </Container>
      )}
    </Field>
  );
}
