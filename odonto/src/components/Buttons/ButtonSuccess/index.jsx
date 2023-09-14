/* eslint-disable react/prop-types */
import React from 'react';
import { StyledButton } from './styles';

export default function ButtonSuccess({
  /**
   * labelbutton
   */
  label,
  /**
   * is loading
   */
  loading = false,
  /**
   * propagate on click event
   */
  onClick,
  /**
   * other props
   */
  ...props
}) {
  function handleClick(event) {
    if (onClick) onClick(event);
  }

  return (
    <StyledButton
      {...props}
      variant="success"
      onClick={(e) => handleClick(e)}
      disabled={loading}
    >
      {loading ? 'Aguarde...' : label}
    </StyledButton>
  );
}
