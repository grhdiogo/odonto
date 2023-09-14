interface Props {
  size: number
  color?: string
}

export default function Marked({
  size,
  color = 'black',
  ...props
}: Props) {
  return (
    <svg {...props} width={size} height={size / 1.34} viewBox="0 0 35 26" fill="none" xmlns="http://www.w3.org/2000/svg">
      <path d="M31.6018 0.400772C31.0694 -0.133591 30.1964 -0.133591 29.664 0.400772L13.545 16.5158L5.24381 8.15785C4.7114 7.62153 3.84428 7.62153 3.30796 8.15785L0.399304 11.0645C-0.133101 11.595 -0.133101 12.468 0.399304 13.0023L12.5683 25.2516C13.1007 25.782 13.9678 25.782 14.5061 25.2516L34.5085 5.24723C35.0487 4.71287 35.0487 3.83792 34.5085 3.3016L31.6018 0.400772Z" fill={color} />
    </svg>

  );
}
