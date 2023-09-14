/* eslint-disable jsx-a11y/anchor-is-valid */
import { useState } from 'react';
import { FiMail, FiLock } from 'react-icons/fi';
import {
  ButtonBase, Input,
} from '../../components';
import Astronaut from '../../components/illustations/Astronaut';
import AuthServices from '../../services/auth';
import handleNotify from '../../utils/notify';
import {
  Container,
  FormContainer,
  Hero,
  Form,
  Link,
  ActionsContainer,
  Span,
  HeroTitle,
  HeroDescription,
  TextContainer,
  HeroText,
} from './styles';

interface LoginProps {
  onRegisterPress: () => void
  onRedirectToDashBoardPage: (token: string) => void
}

export default function Login({
  onRegisterPress,
  onRedirectToDashBoardPage,
}: LoginProps) {
  // states
  const [email, setEmail] = useState<string>('digsobreira26@gmail.com');
  const [password, setPassword] = useState<string>('123456');
  const [loading, setLoading] = useState<boolean>(false);

  function handleEnterPress() {
    setLoading(true);

    const service = new AuthServices();

    service.login(email, password).then((v) => {
      handleNotify('success', 'Login realizado com sucesso');
      if (onRedirectToDashBoardPage) onRedirectToDashBoardPage(v.jwtToken);
    }).catch(() => {
      handleNotify('error', 'Credenciais inválidas');
      setLoading(false);
    });

    setTimeout(() => {
      setLoading(false);
    }, 2000);
  }

  function handleRegisterPress() {
    if (onRegisterPress) onRegisterPress();
  }

  function handleOthersActions() {
    handleNotify('info', 'Funcionalidade em desenvolvimento');
  }

  return (
    <Container>
      <Hero>
        <Astronaut />
        <HeroText>
          <HeroTitle>
            Bem-vindo a bordo meu amigo!
          </HeroTitle>
          <HeroDescription>
            apenas alguns cliques e começamos.
          </HeroDescription>
        </HeroText>
      </Hero>
      <FormContainer>
        <TextContainer>
          <HeroTitle>
            Bem-vindo a bordo meu amigo!
          </HeroTitle>
          <HeroDescription>
            apenas alguns cliques e começamos.
          </HeroDescription>
        </TextContainer>
        <Form>
          <Input
            placeholder={'Email'}
            type={'email'}
            required
            icon={<FiMail size={24} color={'#828282'} />}
            onChangeText={(value) => setEmail(value)}
          />
          <Input
            placeholder={'Senha'}
            type={'password'}
            required
            icon={<FiLock size={24} color={'#828282'} />}
            onChangeText={(value) => setPassword(value)}
          />
          <Link onClick={() => handleOthersActions()}>
            Esqueceu a senha?
          </Link>
          <ButtonBase
            kind={'primary'}
            label={'Entrar'}
            type={'submit'}
            isLoading={loading}
            onClick={() => handleEnterPress()}
          />
        </Form>
        <ActionsContainer>
          {/* <Row>
            <Divider />
            <Span>OU</Span>
            <Divider />
          </Row>
          <Row>
            <ButtonBase
              kind={'outline'}
              label={'Google'}
              icon={<Google />}
              onClick={() => handleOthersActions()}
            />
            <ButtonBase
              kind={'outline'}
              label={'Gitlab'}
              icon={<Gitlab />}
              onClick={() => handleOthersActions()}
            />
          </Row> */}
          <Span>Ainda não tem uma conta?</Span>
          <ButtonBase
            kind={'outline'}
            label={'Registre-se'}
            onClick={() => handleRegisterPress()}
          />
        </ActionsContainer>
      </FormContainer>
    </Container>
  );
}
