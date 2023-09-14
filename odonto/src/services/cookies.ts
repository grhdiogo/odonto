import Cookies from 'universal-cookie';

// get cookie by name
export function getToken() {
  const cookies = new Cookies();
  return cookies.get('stoken');
}

export function removeCookie() {
  const cookies = new Cookies();
  return cookies.remove('stoken');
}

export function setCookie(tkn: string) {
  const cookies = new Cookies();
  return cookies.set('stoken', tkn);
}
