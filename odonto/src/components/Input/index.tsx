/* eslint-disable react/jsx-no-bind */
import { useRef, useState } from 'react';
import { FiEye, FiEyeOff } from 'react-icons/fi';
import {
  InputContainer,
  IconContainer,
  InputStyled,
} from './styles';

interface Props extends React.InputHTMLAttributes<HTMLInputElement> {
  icon?: React.ReactNode
  type?: string
  onChangeText: (text: string) => void
}

export default function Input({
  icon,
  type = 'text',
  onChangeText,
  ...props
}: Props) {
  const [viewPassword, setViewPassword] = useState<boolean>(false);
  const inputRef = useRef<HTMLInputElement>(null);
  const isPassowrdType: boolean = type === 'password';

  function handleChange() {
    const { current: input } = inputRef;

    if (!input || input.value.trim().length === 0) {
      return;
    }

    if (onChangeText) onChangeText(input.value);
  }

  return (
    <InputContainer>
      {
        icon ? (
          <IconContainer>
            {icon}
          </IconContainer>
        ) : null
      }
      <InputStyled
        onChange={handleChange}
        type={isPassowrdType && viewPassword ? 'text' : type}
        ref={inputRef}
        {...props}
      />
      {
        isPassowrdType ? (
          <IconContainer
            onClick={() => setViewPassword(!viewPassword)}
          >
            {
              viewPassword ? (
                <FiEye size={24} color={'#828282'} />
              ) : (
                <FiEyeOff size={24} color={'#828282'} />
              )
            }
          </IconContainer>
        ) : null
      }
    </InputContainer>
  );
}
