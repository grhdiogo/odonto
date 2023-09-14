export function arrayIsNullOrEmpty(arr) {
  if (!arr || !Array.isArray(arr)) return true;
  if (arr.length === 0) return true;
  return false;
}
