import {
  Container,
  IconContainer,
  Label,
} from './styles';

interface ButtonBaseProps extends React.InputHTMLAttributes<HTMLButtonElement> {
  label: string
  icon?: React.ReactNode
  isLoading?: boolean
  onClick?: () => void
  type ?: 'button' | 'submit'
  kind: ButtonKind
}

export type ButtonKind = 'primary' | 'outline' | 'danger';

export default function ButtonBase({
  label,
  icon,
  isLoading = false,
  onClick,
  kind = 'primary',
  type = 'button',
}: ButtonBaseProps) {
  function handleClick() {
    if (onClick) onClick();
  }

  return (
    <Container>
      {
        icon ? (
          <IconContainer>
            {icon}
          </IconContainer>
        ) : null
      }
      <Label kind={kind} color={kind} type={type} onClick={() => handleClick()}>
        {!isLoading ? label : 'Aguarde...'}
      </Label>
    </Container>
  );
}
