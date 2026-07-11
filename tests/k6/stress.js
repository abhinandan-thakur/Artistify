import http from 'k6/http';
import { check, sleep } from 'k6';

const AUTH_URL = 'http://auth-service:8180';
const ALBUM_URL = 'http://music-service:8080';

export const options = {
  stages: [
    { duration: '30s', target: 50 },
    { duration: '30s', target: 100 },
    { duration: '30s', target: 200 },
    { duration: '30s', target: 400 },
    { duration: '30s', target: 600 },
    { duration: '30s', target: 800 },
    { duration: '30s', target: 1000 },
  ],
};

export function setup() {
  const payload = JSON.stringify({
    email: __ENV.TEST_EMAIL || 'root',
    password: __ENV.TEST_PASSWORD || 'targaryen',
  });

  const res = http.post(
    `${AUTH_URL}/auth/login`,
    payload,
    {
      headers: {
        'Content-Type': 'application/json',
      },
      timeout: '10s',
    }
  );

  check(res, {
    'login success':
      (r) =>
        r.status === 200 ||
        r.status === 201,
  });

  if (
    res.status !== 200 &&
    res.status !== 201
  ) {
    // console.log(res.body);
    // throw new Error('Login failed');
  }

  return {
    token: res.json('token'),
  };
}

export default function (data) {
  const params = {
    headers: {
      Authorization:
        `Bearer ${data.token}`,
      'Content-Type':
        'application/json',
    },
    timeout: '10s',
  };

  // console.log('GET /albums');

  const albumsRes = http.get(
    `${ALBUM_URL}/albums`,
    {
      ...params,
      tags: {
        name: 'get_albums',
      },
    }
  );

  // console.log(
  //   '/albums:',
  //   albumsRes.status
  // );

  check(albumsRes, {
    'albums status 200':
      (r) => r.status === 200,
  });

  sleep(Math.random() * 2);

  // const payload = JSON.stringify({
  //   album_name: 'Dummy',
  //   artist: 'dummy',
  //   sales: 603000,
  //   rating: 9.2,
  // });

  // console.log(
  //   'POST /admin/albums'
  // );

  // const createRes = http.post(
  //   `${ALBUM_URL}/admin/albums`,
  //   payload,
  //   {
  //     ...params,
  //     tags: {
  //       name: 'create_album',
  //     },
  //   }
  // );

  // console.log(
  //   'POST:',
  //   createRes.status
  // );

  // check(createRes, {
  //   'album created':
  //     (r) =>
  //       r.status === 200 ||
  //       r.status === 201,
  // });

  // if (
  //   createRes.status !== 200 &&
  //   createRes.status !== 201
  // ) {
  //   console.log(createRes.body);
  //   return;
  // }

  // sleep(Math.random() * 3);

  // const albumId =
  //   createRes.json('id');

  // console.log(
  //   'DELETE /admin/albums'
  // );

  // const delRes = http.del(
  //   `${ALBUM_URL}/admin/albums/${albumId}`,
  //   null,
  //   {
  //     ...params,
  //     tags: {
  //       name: 'delete_album',
  //     },
  //   }
  // );

  // console.log(
  //   'DELETE:',
  //   delRes.status
  // );

  // check(delRes, {
  //   'album deleted':
  //     (r) =>
  //       r.status === 200 ||
  //       r.status === 204,
  // });

  // sleep(Math.random() * 2);
}