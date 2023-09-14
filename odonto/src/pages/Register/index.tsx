/* eslint-disable jsx-a11y/anchor-is-valid */
/* eslint-disable consistent-return */
import { useState } from 'react';
import {
  FiMail,
  FiLock,
  FiUser,
  FiSmile,
} from 'react-icons/fi';
import { useNavigate } from 'react-router-dom';
import { ButtonBase, Input } from '../../components';
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
  HeroTitle,
  HeroDescription,
  TextContainer,
  HeroText,
} from './styles';

export default function Register() {
  // states
  const [name, setName] = useState<string>('');
  const [age, setAge] = useState<string>('');
  const [email, setEmail] = useState<string>('');
  const [password, setPassword] = useState<string>('');
  const [confirmation, setConfirmation] = useState<string>('');
  const [loading, setLoading] = useState<boolean>(false);
  // navigation
  const history = useNavigate();

  function handleRegisterPress() {
    setLoading(true);

    const service = new AuthServices();

    service.create(name, age, email, password).then(() => {
      handleNotify('success', 'Cadatro realizado com sucesso');
      history('/');
    }).catch((e) => {
      handleNotify('error', `Erro: ${e.response.data.error}`);
      setLoading(false);
    });

    setTimeout(() => {
      setLoading(false);
    }, 2000);
  }

  function checkPassword() {
    if (password !== confirmation) {
      return handleNotify('error', 'Senhas não conferem');
    }
    handleRegisterPress();
  }

  function checkAge() {
    if (Number(age) < 16) {
      return handleNotify('error', 'Idade mínima para cadatro é 16 anos');
    }
    checkPassword();
  }

  function checkData() {
    if (!name || !age || !email || !password) {
      return handleNotify('error', 'Preencha todos os campos para realizar o cadastro!');
    }
    checkAge();
  }

  function handleGoBack() {
    history('/');
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
            placeholder={'Nome'}
            type={'text'}
            required
            icon={<FiUser size={24} color={'#828282'} />}
            onChangeText={(value) => setName(value)}
          />
          <Input
            placeholder={'Informe sua idade'}
            type={'number'}
            required
            icon={<FiSmile size={24} color={'#828282'} />}
            onChangeText={(value) => setAge(value)}
          />
          <Input
            placeholder={'Informe o seu melhor email'}
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
          <Input
            placeholder={'Confirmação de senha'}
            type={'password'}
            required
            icon={<FiLock size={24} color={'#828282'} />}
            onChangeText={(value) => setConfirmation(value)}
          />
          <ButtonBase
            kind={'primary'}
            label={'Registrar'}
            type={'submit'}
            isLoading={loading}
            onClick={() => checkData()}
          />
        </Form>
        <ActionsContainer>
          <Link onClick={() => handleGoBack()}>
            Cancelar
          </Link>
        </ActionsContainer>
      </FormContainer>
    </Container>
  );
}
