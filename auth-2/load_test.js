import http from 'k6/http';
import { check, sleep } from 'k6';

// Configure load testing.
export const options = {
  stages: [
    { duration: '30s', target: 10 },
    { duration: '30s', target: 50 },
    { duration: '30s', target: 100 },
    { duration: '30s', target: 200 },
    { duration: '30s', target: 0 },
  ],

  thresholds: {
    http_req_failed: ['rate<0.01'],
    http_req_duration: ['p(95)<500'],
  },
};

// Execute API requests.
export default function () {
  const response = http.get('http://localhost:8000/users');

  check(response, {
    'Status is 200': (r) => r.status === 200,
  });

  sleep(1);
}
