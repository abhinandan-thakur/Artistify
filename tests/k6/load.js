import http from 'k6/http';
import { check } from 'k6';

const HOST = __ENV.URL_HOST || 'localhost';
const PORT = __ENV.URL_PORT || '8080';

const BASE_URL = `http://${HOST}:${PORT}`;

export const options = {
  stages: [
    { duration: '10s', target: 2 },
    // { duration: '10s', target: 10 },
    { duration: '10s', target: 0 },
  ],
};

export function setup() {
  const loginPayload = JSON.stringify({
    email: __ENV.TEST_EMAIL || 'root',
    password: __ENV.TEST_PASSWORD || 'targaryen',
  });

  const loginHeaders = {
    headers: {
      'Content-Type': 'application/json',
    },
  };

  const loginResponse = http.post(
    `${BASE_URL}/auth/login`,
    loginPayload,
    loginHeaders
  );

  console.log(loginResponse.status);
  console.log(loginResponse.body);

  const token = loginResponse.json('token');

  return { token };
}

export default function (data) {
  const headers = {
    headers: {
      Authorization: `Bearer ${data.token}`,
      'Content-Type': 'application/json',
    },
  };

  const res = http.get(`${BASE_URL}/albums`);

  check(res, {
    'albums status is 200': (r) => r.status === 200,
  });

  const payload = JSON.stringify({
    album_name: 'Dummy',
    artist: 'dummy',
    sales: 603000,
    rating: 9.2,
  });

  const postRes = http.post(
    `${BASE_URL}/admin/albums`,
    payload,
    headers
  );

  check(postRes, {
    'album created': (r) =>
      r.status === 200 || r.status === 201,
  });

  if (postRes.status !== 200 &&
      postRes.status !== 201) {
    console.log(postRes.body);
    return;
  }

  const albumId = postRes.json().id;

  const delRes = http.del(
    `${BASE_URL}/admin/albums/${albumId}`,
    null,
    headers
  );

  check(delRes, {
    'album deleted': (r) =>
      r.status === 200 || r.status === 204,
  });

  console.log(delRes.status);
console.log(delRes.body);
}