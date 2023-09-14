import { Container, Svg, Path } from './styles';

export default function Tooth1(props: any) {
  return (
    <Container {...props}>
      <Svg width="32" height="109" viewBox="0 0 32 109" fill="none" xmlns="http://www.w3.org/2000/svg">
        <Path d="M8 85C5.6 74.6 3.33333 50.3333 2.5 39.5C14.1 53.1 26.3333 44.1667 31 38C28.8333 50 24.1 75.1 22.5 79.5C20.5 85 22 103.5 17.5 107C13 110.5 11 98 8 85Z" fill="white" />
        <Path d="M1 7.5C1 2.3 6 1.33333 8.5 1.5C24.1 0.699998 29.3333 5.16667 30 7.5C30.3333 12.8333 30.8 24.5 30 28.5C29 33.5 26 39.5 16.5 42C8.9 44 3 37.1667 1 33.5V7.5Z" fill="white" />
        <Path d="M8 85C5.6 74.6 3.33333 50.3333 2.5 39.5C14.1 53.1 26.3333 44.1667 31 38C28.8333 50 24.1 75.1 22.5 79.5C20.5 85 22 103.5 17.5 107C13 110.5 11 98 8 85Z" stroke="black" stroke-width="2" />
        <Path d="M1 7.5C1 2.3 6 1.33333 8.5 1.5C24.1 0.699998 29.3333 5.16667 30 7.5C30.3333 12.8333 30.8 24.5 30 28.5C29 33.5 26 39.5 16.5 42C8.9 44 3 37.1667 1 33.5V7.5Z" stroke="black" stroke-width="2" />
      </Svg>

    </Container>
  );
}
