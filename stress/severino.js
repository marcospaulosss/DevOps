import { check, sleep } from "k6";
import http from "k6/http";


export default function() {
  let res = http.get(`${__ENV.SEVERINO_URL}/api/albums`);
  check(res, {
    "is status 200": (r) => r.status === 200
  });

  sleep(3);
};
