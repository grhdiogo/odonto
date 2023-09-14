import styled from 'styled-components';

type ButtonKind = 'primary' | 'outline' | 'danger';

const kindMap = new Map([
  ['primary', 'var(--primary)'],
  ['outline', 'var(--white)'],
  ['danger', 'var(--danger)'],
]);

const hoverMap = new Map([
  ['primary', 'var(--primary-light-color)'],
  ['outline', 'var(--light-gray)'],
  ['danger', 'var(--danger)'],
]);

const borderMap = new Map([
  ['primary', 'none'],
  ['outline', '1px solid var(--primary)'],
  ['danger', 'var(--white)'],
]);

const colorMap = new Map([
  ['primary', 'var(--white)'],
  ['outline', 'var(--primary)'],
  ['danger', 'var(--white)'],
]);

export const Container = styled.div`
`;

export const IconContainer = styled.div`
  display: flex;
  justify-content: center;
  align-items: center;
`;

export const Label = styled.button`
   width: 100%;
  height: 40px;
  display: flex;
  justify-content: center;
  align-items: center;
  padding: 10px;
  gap: 8px;
  border: ${(props: { kind: ButtonKind }) => borderMap.get(props.kind)};
  color: ${(props: { kind: ButtonKind }) => colorMap.get(props.kind)};
  border-radius: 40px;
  background: ${(props: { kind: ButtonKind }) => kindMap.get(props.kind)};
  transition: 0.2s ease-in-out;

  &:hover {
    transition: 0.2s ease-in-out;
    cursor: pointer;
    background: ${(props: { kind: ButtonKind }) => hoverMap.get(props.kind)};
  }
`;
