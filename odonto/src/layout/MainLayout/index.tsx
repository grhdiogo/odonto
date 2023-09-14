import { useLocation, useNavigate } from 'react-router-dom';
import {
  Aside, AsideItem, Body, BodyContainer, Container, TopBar,
} from './styles';

interface Props {
  children: any,
  asideData: {
    label: string,
    redirectTo: string,
  }[],
}

export default function MainLayout({
  children,
  asideData = [],
}: Props) {
  const location = useLocation();
  const history = useNavigate();
  //

  function redirectTo(path: string) {
    history(path);
  }
  return (
    <Container>
      <TopBar />
      <Body>
        <Aside>
          {asideData.map((v) => (
            <AsideItem
              key={v.redirectTo}
              selected={location.pathname.toLowerCase() === v.redirectTo.toLowerCase()}
              onClick={() => redirectTo(v.redirectTo)}
            >
              {v.label}
            </AsideItem>
          ))}
        </Aside>
        <BodyContainer>
          {children}
        </BodyContainer>
      </Body>
    </Container>
  );
}
