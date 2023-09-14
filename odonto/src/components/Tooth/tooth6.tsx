import { Container, Svg, Path } from './styles';

export default function Tooth1(props: any) {
  return (
    <Container {...props}>
      <Svg width="44" height="130" viewBox="0 0 44 130" fill="none" xmlns="http://www.w3.org/2000/svg">
        <Path d="M10 50C12.4 24 18.6667 6.83333 21.5 1.5C30.5 3.5 31 24.5 31.5 34.5C31.9 42.5 34 65.5 35 76C34.3333 75.5 31.1 73.9 23.5 71.5C16.166 69.184 9.48378 75.0629 6.77847 78.6224C6.63076 79.2502 6.54156 79.2909 6.50001 79C6.58749 78.8775 6.68037 78.7515 6.77847 78.6224C7.22321 76.732 8.19835 69.5179 10 50Z" fill="white" />
        <Path d="M1.50001 100.5C3.10001 81.3 16.1667 76.8333 22.5 77C41.5 79.5 42.5 100.5 42 109.5C41.5 118.5 22.5 129 20 129C17.5 129 -0.499992 124.5 1.50001 100.5Z" fill="white" />
        <Path d="M21.5 1.5C18.6667 6.83333 12.4 24 10 50C7.60001 76 6.66667 80.1667 6.50001 79C9.00001 75.5 15.9 69.1 23.5 71.5C31.1 73.9 34.3333 75.5 35 76C34 65.5 31.9 42.5 31.5 34.5C31 24.5 30.5 3.5 21.5 1.5ZM22.5 77C16.1667 76.8333 3.10001 81.3 1.50001 100.5C-0.499992 124.5 17.5 129 20 129C22.5 129 41.5 118.5 42 109.5C42.5 100.5 41.5 79.5 22.5 77Z" stroke="black" stroke-width="2" />
      </Svg>
    </Container>
  );
}
