import React from 'react';
import { Form as Frm } from 'react-bootstrap';

export default function Form({
  /**
   * children
   */
  children,
  /**
   * others props
   */
  ...props
}) {
  return (
    <Frm {...props}>
      {children}
    </Frm>
  );
}
