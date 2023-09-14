import styled from 'styled-components';
import { colors } from '../../theme';

export const Container = styled.div`
  cursor: pointer;
`;

export const Path = styled.path`
`;

export const Svg = styled.svg`
  &:hover ${Path} {
    fill: ${colors.bgDarkGray};
  }
`;

export const OptionsContainer = styled.div`
  position: relative;
`;

export const OptionsChildren = styled.div`
   & ${Path} {
     fill: ${(props: any) => (props.selected ? colors.bgDarkGray : colors.bgLight)};
   }
`;

export const OptionsTitle = styled.div`
  display: flex;
  align-items: center;
  justify-content: center;
  border-bottom: 1px solid ${colors.textBlack};
  padding: 5px 0px 10px 0px;
  color: ${colors.textBlack};
`;

export const OptionsModal = styled.div`
  background: ${colors.bgGray};
  position: absolute;
  border: 1px solid ${colors.bgBlack};
  box-shadow: 1px 1px;
  border-radius: 4px;

  top: 0;
  left: 50%;
  display: ${(props: { visible: boolean }) => (props.visible ? 'block' : 'none')};
  padding: 5px 15px 10px 15px;
  z-index: 99;
  justify-content: center;
`;

export const OptionContainer = styled.div`
  display: flex;

  color: ${colors.textBlack};
  font-size: 12px;
  align-items: center;
  margin-top: 5px;
  cursor: pointer;
`;

export const MarkedContainer = styled.div`
  display: flex;
  align-items: center;
  justify-content: center;
  border: 1px solid ${colors.textBlack};
  height: 15px;
  width: 15px;
`;

export const Option = styled.span`
  margin-left: 5px;
`;
