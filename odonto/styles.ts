import styled from 'styled-components';

export const Container = styled.div`
  height: 100vh;
  display: flex;
  flex-direction: column; 
  
  @media (min-width: 920px) {
    flex-direction: row; 
  }
`;

export const Hero = styled.div`
  width: 60%;
  height: 100vh;
  padding: 30px;
  display: flex;
  flex-direction: column; 
  justify-content: center;
  align-items: center;
  gap: 8px;

  background: var(--primary);
  
  @media (max-width: 920px) {
    width: 100%;
    height: 20vh;
  }
`;

export const HeroText = styled.div`
  display: flex;
  flex-direction: column; 
  justify-content: center;
  align-items: center;

   @media (max-width: 920px) {
    display: none;
  }
`;

export const HeroTitle = styled.p`
  font-size: 24px;
  line-height: 48px;
  color: var(--white);

  @media (max-width: 920px) {
    font-size: 18px;
    line-height: 28px;
  }

  @media (max-width: 920px) {
    color: var(--primary);
  }
`;

export const HeroDescription = styled.p`
  font-size: 14px;
  line-height: 28px;
  color: var(--white);

  @media (max-width: 920px) {
    color: var(--primary);
  }
`;

export const FormContainer = styled.div`
  flex: 1;
  padding: 30px 80px;
  display: flex;
  flex-direction: column; 
  justify-content: start;
  align-items: center;
  gap: 32px;

  @media (max-width: 920px) {
    width: 100%;
  }
`;

export const TextContainer = styled.div`
   display: none;

   @media (max-width: 920px) {
    width: 100%;
    display: flex;
    flex-direction: column; 
    justify-content: center;
    align-items: center;

    white-space: nowrap;
  }
`;

export const Form = styled.form`
  width: 350px;
  display: flex;
  flex-direction: column; 
  justify-content: center;
  align-items: flex-end;
  gap: 16px;
`;

export const Link = styled.a`
  text-decoration: none;
  font-size: 14px;
  color: var(--link-text);
`;

export const ActionsContainer = styled.div`
  width: 350px;
  display: flex;
  flex-direction: column; 
  justify-content: center;
  align-items: center;
  gap: 16px;
`;

export const Divider = styled.hr`
  width: 30%;
  border: 0.3px solid var(--light-gray);
`;

export const Row = styled.div`
  width: 350px;
  display: flex;
  justify-content: center;
  align-items: center;
  gap: 8px;
`;

export const Span = styled.span`
  font-size: 14px;
  color: var(--primary-text-light);
`;