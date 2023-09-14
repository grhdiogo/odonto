import React, { useState, useEffect, useCallback } from 'react';
import Search from '../Search';
import { TableList } from '../TableList';

export default function ListPanel({
  /**
   * in loading
   */
  loading = false,
  /**
   * table head
   */
  head = [],
  /**
   * table raw data
   */
  data = [],
  /**
   * actions to rows
   */
  actions = [],
  /**
   * propagate search click
   */
  onSearchClick,
  /**
   * propagate action click
   */
  onActionClick,
  //
  onPageChange = null,
  paginationLimit = 0,
  paginationTotal = 0,
}) {
  // rows state
  const [rows, setRows] = useState([]);

  // callback: refresh news list
  const adapter = useCallback((internalData) => {
    // head keys selected
    const headKeys = head.map((h) => h.columnKey);
    // convert to columns (key, value)
    const parsedRows = internalData.map((d) => headKeys.map((k) => ({
      columnKey: k,
      columnValue: d[k],
    })));
    // set values
    setRows(parsedRows);
  }, [head]);

  // hook: on create component
  useEffect(() => {
    adapter(data);
  }, [adapter, data]);

  // handle search click event
  function handleSearchClick(text) {
    if (onSearchClick) onSearchClick(text);
  }

  return (
    <>
      {/* input search  */}
      <Search
        loading={loading}
        onSearchClick={(text) => handleSearchClick(text)}
      />
      {/* table list */}
      <TableList
        head={head}
        data={rows}
        loading={loading}
        actions={actions}
        onActionClick={onActionClick}
        onPageChange={onPageChange}
        paginationLimit={paginationLimit}
        paginationTotal={paginationTotal}
      />
    </>
  );
}
