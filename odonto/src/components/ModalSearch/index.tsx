import { Container } from './styles';

export default function ModalSearch() {
  function handleClick() {
    console.log('aaaaa');
  }
  return (
    <Container>
      <input disabled placeholder={'paciente'} />
      <button type={'button'} onClick={() => handleClick()}>Pesquisar</button>
    </Container>
  );
}
