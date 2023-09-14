/* eslint-disable react/no-array-index-key */
import React from 'react';

export default function EmptyBody({
  /**
   * Number of empty columns
   */
  columnsCount = 0,
}) {
  return (
    <tbody>
      <tr>{Array(columnsCount + 1).fill().map((_, ix) => (<td key={`col-${ix}`}><br /></td>))}</tr>
    </tbody>
  );
}
