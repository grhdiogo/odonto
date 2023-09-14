import styled from 'styled-components';
import { colors } from '../../theme';

export const Container = styled.div`
  height: 100vh;
  width: 100vw;
  display: flex;
  flex-direction: column;
`;

export const TopBar = styled.div`
  width: 100%;
  height: 70px;
  border: 1px solid ${colors.bgBlack};
`;

export const Aside = styled.div`
  height: 100%;
  width: 150px;
  background: ${colors.bgGray};
  padding: 10px;
  display: flex;
  flex-direction: column;
  gap: 15px;
`;

export const Body = styled.div`
  height: 100%;
  width: 100%;
  display: flex;
  flex-direction: row;
`;

export const BodyContainer = styled.div`
  width: 100%;
  height: 100%;
`;

export const AsideItem = styled.div`
  color: ${colors.textBlack};
  text-decoration: ${(props: { selected: boolean }) => (props.selected ? 'underline' : 'none')};
  cursor: pointer;
`;
