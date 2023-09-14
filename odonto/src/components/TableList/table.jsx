/* eslint-disable react/prop-types */
import React from 'react';
import { Table } from 'react-bootstrap';
import Head from './head';
import Body from './body';
import EmptyBody from './empty';
import { arrayIsNullOrEmpty } from '../../utils/array';
// import Button from '../Button';

export default function TableList({
  /**
   * table head
   * - [{columnKey, columnValue}]
   */
  head = [],
  /**
   * table data
   * - dataRow = [{columnKey, columnValue}]
   * - dataRows = [dataRow]
   */
  data = [],
  /**
   * table actions
   * - [{action, name, variant}]
   */
  actions = [],
  /**
   * propagate action click
   */
  onActionClick,
  /**
   * in loading
   */
  loading = false,
  //
  onPageChange = null,
  paginationLimit = 0,
  paginationTotal = 0,
}) {
  return (
    <Table striped bordered hover variant="light">
      {/* table head */}
      <Head row={head} withAction={!arrayIsNullOrEmpty(actions)} />
      {/* table body */}
      {Array.isArray(data) && data.length > 0 ? (
        <Body
          actions={actions}
          loading={loading}
          rows={data}
          emptyRows={[head.map((h) => ({ columnKey: h.columnKey }))]}
          onActionClick={onActionClick}
          onPageChange={onPageChange}
          paginationLimit={paginationLimit}
          paginationTotal={paginationTotal}
        />
      ) : (
        <EmptyBody columnsCount={head.length} />
      )}

    </Table>
  );
}
