export function numberToMoney(value: number) {
  if (!value) return 'R$ 0,00';
  if (typeof (value) !== 'number') return 'R$ 0,00';
  // create our number formatter.
  const formatter = new Intl.NumberFormat('pt-BR', {
    style: 'currency',
    currency: 'BRL',
  });
  return formatter.format(value);
}
