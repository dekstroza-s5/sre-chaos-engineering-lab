import http from "k6/http";
import {check, sleep} from "k6";
export const options = {
  stages: [{duration:"30s",target:25},{duration:"2m",target:25},{duration:"30s",target:0}],
  thresholds: {http_req_failed:["rate<0.001"],http_req_duration:["p(95)<300"]}
};
export default function(){
  const response=http.get((__ENV.BASE_URL||"http://localhost:8080")+"/work");
  check(response,{"status is 200":r=>r.status===200});
  sleep(0.1);
}
