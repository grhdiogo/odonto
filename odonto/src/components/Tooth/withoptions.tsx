import { useState } from 'react';
import Marked from '../illustations/Marked';
import {
  OptionsChildren, OptionsContainer, OptionsModal, Option,
  OptionContainer, MarkedContainer, OptionsTitle,
} from './styles';
import { colors } from '../../theme';

interface Opts {
  id: string
  label: string
  marked: boolean
}

interface Props {
  onOptionClick?: (id: string) => void
  icon: JSX.Element,
  options: Opts[],
  selected?: boolean,
}

export default function WithOptions({
  onOptionClick,
  options = [],
  icon,
  selected = false,
}: Props) {
  const [isVisible, setIsVIsible] = useState(false);
  // change modal visibility
  function showModal() {
    setIsVIsible(!isVisible);
  }
  // handle option click
  function handleOptionClick(id: string) {
    if (onOptionClick) onOptionClick(id);
  }

  const props = {
    selected,
  };
  //
  return (
    <OptionsContainer
      onMouseLeave={() => setIsVIsible(false)}
    >
      <OptionsChildren
        onClick={() => showModal()}
        {...props}
      >
        {icon}
      </OptionsChildren>
      <OptionsModal visible={isVisible}>
        <OptionsTitle>
          Opções
        </OptionsTitle>
        {options.map((v) => (
          <OptionContainer key={v.id} onClick={() => handleOptionClick(v.id)}>
            <MarkedContainer>
              <Marked color={v.marked ? colors.bgBlack : 'none'} size={10} />
            </MarkedContainer>
            <Option>{v.label}</Option>
          </OptionContainer>
        ))}
      </OptionsModal>
    </OptionsContainer>
  );
}
