import http from 'k6/http';
import {check, sleep} from 'k6';

const BASE_URL = 'http://localhost:8080';

export const options = {
    vus:30,          //virtual users
    duration: '30s',

    thresholds:{
        http_req_duration: ['p(95)<500'], // 95% under 500ms
        http_req_failed: ['rate<0.01'],   // <1% failures
    },
};

export default function () {
  const res = http.get(`${BASE_URL}/albums/`);

  check(res, {
    'status is 200': (r) => r.status === 200,
    'response under 500ms': (r) => r.timings.duration < 500,
  });

  sleep(1);
}