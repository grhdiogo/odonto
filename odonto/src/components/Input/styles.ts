import styled from 'styled-components';

export const InputContainer = styled.div`
  width: 100%;
  max-width: 380px;
  height: 40px;
  padding: 8px;
  border: 1px solid var(--control-border);
  border-radius: var(--radius-default);

  display: flex;
  align-items: center;

  &:focus-within {
    border: 2px solid var(--primary);
  }
`;

export const IconContainer = styled.div`
  width: 32px;
  height: 32px;
  
  display: flex;
  align-items: center;
  justify-content: center;

  :hover {
    cursor: pointer;
  }
`;

export const InputStyled = styled.input`
  width: 100%;
  padding: 6px;
  
  color: var(--primary-text);
  font-weight: 400;
  font-size: 14px;
  border: none;

  ::placeholder {
    color: var(--primary-text-light);
  }
`;
