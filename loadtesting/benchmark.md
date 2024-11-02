# Benchmark of ZexD

## Environment

- Machine: Lenovo Ideapad Flex 5
- CPU: 11th Gen Intel(R) Core(TM) i7-1165G7 @ 2.80GHz
- RAM: 16GB
- OS: Windows 11 23H2
- Docker: Docker Desktop Windows

## Setup

This benchmark was performed using [k6](https://k6.io/) load testing tool.
It works with the concept of virtual users (VUs) and iterations.
The test script is written in JavaScript. And it is run by the VUs as fast as possible.

## The Results

If we do not want to read the whole document, here are the results of the tests:

- [Tests for Url Creation](#tests-for-url-creation)
  - [Test 1](#test-1)
    - 1 VU for 30 seconds
    - 14845 requests
    - Avg. response time: 1.92ms
  - [Test 2](#test-2)
    - 100 VU for 2 minutes
    - 183464 requests
    - Avg. response time: 65.25ms
  - [Test 3](#test-3)
    - 200 VU for 2 minutes
    - 356080 requests
    - Avg. response time: 67.23ms

- [Tests for Url Redirection](#tests-for-url-redirection)
  - [Test 1](#test-1)
    - 1 VU for 30 seconds
    - 10713 requests
    - Avg. response time: 2.7ms
  - [Test 2](#test-2)
    - 100 VU for 2 minutes
    - 165830 requests
    - Avg. response time: 72.03ms
  - [Test 3](#test-3)
    - The server was not able to handle the load of 200 VUs for 2 mins.

## Tests for Url Creation

### Test 1

The test script is [`test_url_create.js`](test_url_create.js). It creates a new URL and checks if the response status is 200.
It was run by 1 VU for 30 seconds.

The results are as follows:

```bash
     execution: local
        script: ./test_url_create.js

     scenarios: (100.00%) 1 scenario, 1 max VUs, 1m0s max duration (incl. graceful stop):
              * default: 1 looping VUs for 30s (gracefulStop: 30s)

     checks.........................: 100.00% 14845 out of 14845
     http_req_duration..............: avg=1.92ms  min=503.4µs med=1.67ms max=16.95ms p(90)=2.62ms   p(95)=3.16ms
     http_req_failed................: 0.00%   0 out of 14845
     http_req_receiving.............: avg=53.03µs min=0s      med=0s     max=1.57ms  p(90)=156.56µs p(95)=521.5µs
     http_req_sending...............: avg=13.68µs min=0s      med=0s     max=1.5ms   p(90)=0s       p(95)=0s
     http_req_waiting...............: avg=1.85ms  min=503.4µs med=1.61ms max=16.95ms p(90)=2.58ms   p(95)=3.01ms
     http_reqs......................: 14845   494.813312/s
     iteration_duration.............: avg=2.01ms  min=503.4µs med=1.89ms max=16.95ms p(90)=2.66ms   p(95)=3.2ms
     iterations.....................: 14845   494.813312/s
     vus............................: 1       min=1              max=1

running (0m30.0s), 0/1 VUs, 14845 complete and 0 interrupted iterations
default ✓ [======================================] 1 VUs  30s
```

### Test 2

The test script is [`test_url_create.js`](test_url_create.js). It creates a new URL and checks if the response status is 200.
It was run by 100 VU concurrently for 2 minutes.

The results are as follows:

```bash
execution: local
        script: ./test_url_create.js

     scenarios: (100.00%) 1 scenario, 100 max VUs, 2m30s max duration (incl. graceful stop):
              * default: 100 looping VUs for 2m0s (gracefulStop: 30s)

     checks.........................: 100.00% 183464 out of 183464
     http_req_duration..............: avg=65.25ms min=2.09ms med=36.52ms max=1.04s    p(90)=170.27ms p(95)=236.89ms
     http_req_receiving.............: avg=80.94µs min=0s     med=0s      max=103.82ms p(90)=234.6µs  p(95)=532.59µs
     http_req_sending...............: avg=21.07µs min=0s     med=0s      max=71.79ms  p(90)=0s       p(95)=0s
     http_req_waiting...............: avg=65.15ms min=2.09ms med=36.43ms max=1.04s    p(90)=170.15ms p(95)=236.59ms
     http_reqs......................: 183464  1527.771976/s
     iteration_duration.............: avg=65.41ms min=2.28ms med=36.68ms max=1.04s    p(90)=170.42ms p(95)=237.02ms
     iterations.....................: 183464  1527.771976/s
     vus............................: 100     min=100              max=100


running (2m00.1s), 000/100 VUs, 183464 complete and 0 interrupted iterations
default ✓ [======================================] 100 VUs  2m0s
```

### Test 3

The test script is [`test_url_create.js`](test_url_create.js). It creates a new URL and checks if the response status is 200.
It was run by 200 VU concurrently for 2 minutes.

The results are as follows:

```bash
     execution: local
        script: ./test_url_create.js

     scenarios: (100.00%) 1 scenario, 200 max VUs, 2m30s max duration (incl. graceful stop):
              * default: 200 looping VUs for 2m0s (gracefulStop: 30s)

     checks.........................: 100.00% 356080 out of 356080
     http_req_duration..............: avg=67.23ms min=1.52ms med=58.81ms max=727.96ms p(90)=109.99ms p(95)=134.54ms
     http_req_receiving.............: avg=68.81µs min=0s     med=0s      max=93.42ms  p(90)=80.71µs  p(95)=528.1µs
     http_req_sending...............: avg=20.73µs min=0s     med=0s      max=69.92ms  p(90)=0s       p(95)=0s
     http_req_waiting...............: avg=67.14ms min=1.23ms med=58.74ms max=727.96ms p(90)=109.86ms p(95)=134.43ms
     http_reqs......................: 356080  2965.844153/s
     iteration_duration.............: avg=67.38ms min=1.6ms  med=58.94ms max=727.96ms p(90)=110.18ms p(95)=134.74ms
     iterations.....................: 356080  2965.844153/s
     vus............................: 200     min=200              max=200

running (2m00.1s), 000/200 VUs, 356080 complete and 0 interrupted iterations
default ✓ [======================================] 200 VUs  2m0s
```

## Tests for Url Redirection

For this test to run the code had to be modified to return the original URL in the response. As for some reason redirection counting was not working.

### Test 1

The test script is [`test_url_redirect.js`](test_url_delete.js). It redirects to the original URL and checks if the response status is 302.
It was run by 1 VU for 30 seconds.

The results are as follows:

```bash
     execution: local
        script: ./test_url_redirect.js

     scenarios: (100.00%) 1 scenario, 1 max VUs, 1m0s max duration (incl. graceful stop):
              * default: 1 looping VUs for 30s (gracefulStop: 30s)

     checks.........................: 100.00% 10713 out of 10713
     http_req_duration..............: avg=2.7ms   min=503.7µs med=2.43ms max=23.21ms  p(90)=3.86ms   p(95)=4.48ms
     http_req_failed................: 0.00%   0 out of 10713
     http_req_receiving.............: avg=58.65µs min=0s      med=0s     max=1.68ms   p(90)=363.12µs p(95)=524.9µs
     http_req_sending...............: avg=12.73µs min=0s      med=0s     max=1.51ms   p(90)=0s       p(95)=0s
     http_req_waiting...............: avg=2.63ms  min=503.7µs med=2.26ms max=22.16ms  p(90)=3.79ms   p(95)=4.37ms
     http_reqs......................: 10713   357.082608/s
     iteration_duration.............: avg=2.78ms  min=504.9µs med=2.6ms  max=23.21ms  p(90)=4.05ms   p(95)=4.7ms
     iterations.....................: 10713   357.082608/s
     vus............................: 1       min=1              max=1

running (0m30.0s), 0/1 VUs, 10713 complete and 0 interrupted iterations
default ✓ [======================================] 1 VUs  30s
```

### Test 2

The test script is [`test_url_redirect.js`](test_url_delete.js). It redirects to the original URL and checks if the response status is 302.
It was run by 100 VU concurrently for 2 minutes.

The results are as follows:

```bash
     execution: local
        script: ./test_url_create.js

     scenarios: (100.00%) 1 scenario, 100 max VUs, 2m30s max duration (incl. graceful stop):
              * default: 100 looping VUs for 2m0s (gracefulStop: 30s)

     checks.........................: 100.00% 165830 out of 165830
     http_req_duration..............: avg=72.03ms  min=2.08ms med=38.42ms max=4.98s    p(90)=183.32ms p(95)=248.31ms
     http_req_failed................: 0.00%   0 out of 165830
     http_req_receiving.............: avg=457.95µs min=0s     med=0s      max=3.4s     p(90)=241.32µs p(95)=533.5µs
     http_req_sending...............: avg=29.44µs  min=0s     med=0s      max=1.45s    p(90)=0s       p(95)=0s
     http_req_waiting...............: avg=71.54ms  min=2.08ms med=38.33ms max=2.03s    p(90)=183.23ms p(95)=248.17ms
     http_reqs......................: 165830  1379.642998/s
     iteration_duration.............: avg=72.39ms  min=2.08ms med=38.55ms max=4.99s    p(90)=183.65ms p(95)=248.76ms
     iterations.....................: 165830  1379.642998/s
     vus............................: 100     min=100              max=100

running (2m00.2s), 000/100 VUs, 165830 complete and 0 interrupted iterations
default ✓ [======================================] 100 VUs  2m0s
```

### Test 3

The test script is [`test_url_redirect.js`](test_url_delete.js). It redirects to the original URL and checks if the response status is 302.
It was run by 200 VU concurrently for 2 minutes.

Well this test was not run as the server was not able to handle the load of 200 VUs. The server was crashing and the response time was very high.
