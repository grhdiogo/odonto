import styled from 'styled-components';
// theme
import { defaultTheme } from '../../theme';

export const Label = styled.p`
  ${defaultTheme.txRegular16Black};
  margin-bottom: 8px;
`;

export const Container = styled.div`
  padding-top: 5px;
`;

export const Input = styled.input`
  border-radius: 8px;
  ${defaultTheme.txRegular16Black};
`;

export const TextArea = styled.textarea`
  border-radius: 8px;
  ${defaultTheme.txRegular16Black};
`;
