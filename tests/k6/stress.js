import http from 'k6/http';
import { check } from 'k6';

const BASE_URL = 'http://localhost:8080';
const TOKEN = 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3NzkwODcyMzYsImlkIjoyLCJyb2xlIjoiYXJ0aXN0In0.SRLkNvbkoel-tZoskJL2ecUquPkxHES6Z6KtdUca51M';

export const options = {
  stages: [
    {duration: '10s', target: 2},
    { duration: '15s', target: 5}
    // { duration: '30s', target: 10 },
    // { duration: '30s', target: 100 },
    // { duration: '30s', target: 250 },
    // { duration: '30s', target: 500 },
    // { duration: '30s', target: 1000 },
    // { duration: '30s', target: 0 },
  ],

  thresholds: {
    http_req_failed: ['rate<0.05'],
    http_req_duration: ['p(95)<1000'],
  },
};

export default function () {
  const res = http.get(`${BASE_URL}/albums/`);

  check(res, {
    'status is 200': (r) => r.status === 200,
  });

  const siaRes = http.get (`http://localhost:8080/albums/artist/Sia`);
    check(res, {
    'status is 200': (r) => r.status === 200,
  });


    const payload = JSON.stringify({
        "album_name": "Dummy",
        "artist": "dummy",
        "sales": 603000,
        "rating": 9.2
    });

    const headers = {
        headers: {
            'Content-Type': 'application/json',
            Authorization: `Bearer ${TOKEN}`
        },
    };

    const postRes = http.post(
        `${BASE_URL}/artist/albums/`,
        payload,
        headers
    );

    check(postRes, {
    'Post status is 200': (r) => r.status === 200 || r.status === 201,
  });

//   console.log(postRes.status);
// console.log(postRes.body);

    // parse response JSON
  const createdAlbum = postRes.json();

  // extract id
  const albumId = createdAlbum.id;
  console.log(albumId)

      const delRes = http.del(
    `${BASE_URL}/artist/albums/${albumId}`,
    null,
    headers
  );

    check(delRes, {
    'Delete status is 200': (r) => r.status === 200 || r.status === 204,
  });

  console.log(delRes.status);
console.log(delRes.body);


}