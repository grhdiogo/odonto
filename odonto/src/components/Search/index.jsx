import React, { useState, useRef } from 'react';
import { FormControl, InputGroup } from 'react-bootstrap';

export default function Search({
  /**
   * in loading
   */
  loading,
  /**
   * propagate search click event
   * - parameter: filterText
   */
  onSearchClick,
}) {
  // state
  const [filterText, setFilterText] = useState('');
  const timerRef = useRef(null);

  /**
   * handle text filter change in input
   */
  function handleFilterTextChange(e) {
    const { value } = e.target;
    setFilterText(value);
    clearTimeout(timerRef.current);
    timerRef.current = setTimeout(() => {
      if (onSearchClick) onSearchClick(filterText);
    }, 500);
  }

  return (
    <InputGroup className="mb-3">
      <FormControl
        placeholder="Busque na tabela"
        aria-label="Busque na tabela"
        aria-describedby="basic-addon2"
        value={filterText}
        onChange={(e) => handleFilterTextChange(e)}
      />
    </InputGroup>
  );
}
