import { Container, Svg, Path } from './styles';

export default function Tooth1(props: any) {
  return (
    <Container {...props}>
      <Svg width="42" height="117" viewBox="0 0 42 117" fill="none" xmlns="http://www.w3.org/2000/svg">
        <Path d="M6 75.5C20.4 64.7 33 71 37.5 75.5C46 114 35.5 113 33.5 114.5C31.5 116 21 114.5 9.5 113C-2 111.5 1.5 87.5 6 75.5Z" fill="white" />
        <Path d="M8 69C15.2 61.8 29.6667 64.6667 36 67C34.6667 61.6667 31.3 48.6 28.5 39C25 27 30 -1.5 23 2C16 5.5 3 70.5 8 69Z" fill="white" />
        <Path d="M6 75.5C20.4 64.7 33 71 37.5 75.5C46 114 35.5 113 33.5 114.5C31.5 116 21 114.5 9.5 113C-2 111.5 1.5 87.5 6 75.5Z" stroke="black" stroke-width="2" />
        <Path d="M8 69C15.2 61.8 29.6667 64.6667 36 67C34.6667 61.6667 31.3 48.6 28.5 39C25 27 30 -1.5 23 2C16 5.5 3 70.5 8 69Z" stroke="black" stroke-width="2" />
      </Svg>
    </Container>
  );
}
