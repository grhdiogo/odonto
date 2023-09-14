import React from 'react';

export default function Head({
  /**
   * set table head with action column
   */
  withAction = false,
  /**
   * table row data
   * - [{columnKey, columnValue}]
   */
  row = [],
}) {
  return (
    <thead>
      <tr>
        {/* datas */}
        {row.map((col) => (
          <th key={`${col.columnKey}`}>{col.columnValue}</th>
        ))}
        {/* actions */}
        {withAction ? (
          <th>{'Ações'}</th>
        ) : null}
      </tr>
    </thead>
  );
}
