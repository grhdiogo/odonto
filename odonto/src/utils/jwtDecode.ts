import jwtDecode from 'jwt-decode';

export function jwtTokenToObject(token: string): object {
  return jwtDecode(token);
}
