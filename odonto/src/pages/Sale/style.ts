import styled from 'styled-components';
import Select from 'react-select';
import { ButtonBase } from '../../components';

export const Container = styled.div`
`;

export const StyledSelect = styled(Select)`
  margin-bottom: 30px;
`;

export const FormContainer = styled.div`
  display: flex;
  flex-direction: row;
  height: 100%;
  width: 100%;
`;

export const Column = styled.div`
  height: 100%;
  width: 100%;
  display: flex;
  flex-direction: column;
`;

export const Row = styled.div`
  height: 100%;
  width: 100%;
  display: flex;
  flex-direction: row;
`;

export const ToothsContainer = styled.div`
  display: flex;
  flex-direction: row;
  justify-content: space-evenly;
  padding: 10px 10px;
`;

export const ToothContainer = styled.div`
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: space-between;
`;

export const Label = styled.p`
  font-size: 25px;
  padding: 0;
  margin: 0;
`;

export const ResumeContainer = styled.div`
  display: flex;
  width: 300px;
  flex-direction: column;
  border: 1px solid black;
  border-radius: 5px;
  padding: 10px;
`;

export const ResumeTitle = styled.p`
  font-size: 18px;
`;

export const ResumeText = styled.p`
  font-size: 10px;
`;

export const ButtonContainer = styled.div`
  display: flex;
  margin-left: 10px;
`;

export const StyledButton = styled(ButtonBase)`
`;
