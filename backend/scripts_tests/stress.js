import http from 'k6/http';
import { sleep, check } from 'k6';

export const options = {
    vus: 300,
    duration: '30s',
  }

export default function () {
  const response = http.post('');
  check(response, {
    "Lista os usuários com sucesso": (response) => response.status === 200,
  });
  sleep(1);
}
