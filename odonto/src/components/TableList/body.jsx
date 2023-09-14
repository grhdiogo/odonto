import React from 'react';
import Container from 'react-bootstrap/Container';
import { Pagination } from 'antd';
import Nav from 'react-bootstrap/Nav';
import Navbar from 'react-bootstrap/Navbar';
import NavDropdown from 'react-bootstrap/NavDropdown';
import { arrayIsNullOrEmpty } from '../../utils/array';
import { ButtonContainer } from './styles';

export default function Body({
  /**
   * actions
   * - [{action, name, variant}]
   * - action: is key
   * - name: used for display button
   * - variant: see in 'bootstrap button variant'
   */
  actions = [],
  /**
   * table rows data
   * - dataRow = [{columnKey, columnValue}]
   * - dataRows = [dataRow]
   */
  rows = [],
  /**
   * propagate action click event
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
  /**
   * handle action click
   */
  function handleActionClick(action, index) {
    if (onActionClick) onActionClick(action, index);
  }

  function handlePageChange(n) {
    if (onPageChange) onPageChange(n);
  }

  return (
    <tbody>
      {rows.map((cols, index) => (
        <tr key={`${index}`}>
          {cols.map(({ columnKey, columnValue }, colIndex) => (
            <td key={columnKey} className={`col-body col-body-${colIndex}`}>{columnValue}</td>
          ))}
          {!arrayIsNullOrEmpty(actions) ? (
            <td className={''}>
              <ButtonContainer>
                <Navbar>
                  <Container fluid>
                    <Navbar.Toggle aria-controls="navbar-dark-example" />
                    <Navbar.Collapse id="navbar-dark-example">
                      <Nav>
                        <NavDropdown
                          id="nav-dropdown"
                          title=""
                          alignRight
                        >
                          {actions.map((a, i) => (
                            <NavDropdown.Item
                              eventKey={i}
                              onClick={() => handleActionClick(a.action, index)}
                            >{a.name}
                            </NavDropdown.Item>
                          ))}
                        </NavDropdown>
                      </Nav>
                    </Navbar.Collapse>
                  </Container>
                </Navbar>
              </ButtonContainer>
            </td>
          ) : null}
        </tr>
      ))}
      <Pagination total={paginationTotal} defaultPageSize={paginationLimit} onChange={(v) => handlePageChange(v)} />
    </tbody>
  );
}
