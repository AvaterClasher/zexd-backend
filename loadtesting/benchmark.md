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
    - 16519 requests
    - Avg. response time: 1.71ms
  - [Test 2](#test-2)
    - 100 VU for 2 minutes
    - 253477 requests
    - Avg. response time: 47.19ms
  - [Test 3](#test-3)
    - 200 VU for 2 minutes
    - 330481 requests
    - Avg. response time: 72.48ms
  - [Test 4](#test-4)
    - 1000 VU for 2 minutes
    - 324233 requests
    - Avg. response time: 369.2ms

- [Tests for Url Redirection](#tests-for-url-redirection)
  - [Test 1](#test-1)
    - 1 VU for 30 seconds
    - 25053 requests
    - Avg. response time: 1.09ms
  - [Test 2](#test-2)
    - 100 VU for 2 minutes
    - 170563 requests
    - Avg. response time: 12.39ms
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

     checks.........................: 100.00% 16519 out of 16519
     http_req_blocked...............: avg=3.47µs  min=0s      med=0s     max=5.89ms  p(90)=0s     p(95)=0s
     http_req_connecting............: avg=60ns    min=0s      med=0s     max=1ms     p(90)=0s     p(95)=0s
     http_req_duration..............: avg=1.71ms  min=503.1µs med=1.57ms max=15.04ms p(90)=2.14ms p(95)=2.64ms
     http_req_failed................: 0.00%   0 out of 16519
     http_req_receiving.............: avg=48.45µs min=0s      med=0s     max=1.56ms  p(90)=0s     p(95)=520.1µs
     http_req_sending...............: avg=15.4µs  min=0s      med=0s     max=1.01ms  p(90)=0s     p(95)=0s
     http_req_waiting...............: avg=1.64ms  min=503.1µs med=1.56ms max=15.04ms p(90)=2.12ms p(95)=2.61ms
     http_reqs......................: 16519   550.614475/s
     iteration_duration.............: avg=1.8ms   min=503.1µs med=1.58ms max=15.04ms p(90)=2.18ms p(95)=2.67ms
     iterations.....................: 16519   550.614475/s
     vus............................: 1       min=1              max=1

running (0m30.0s), 0/1 VUs, 16519 complete and 0 interrupted iterations
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

     checks.........................: 100.00% 253477 out of 253477
     http_req_blocked...............: avg=5.01µs  min=0s      med=0s      max=7.67ms   p(90)=0s       p(95)=0s
     http_req_connecting............: avg=580ns   min=0s      med=0s      max=7.12ms   p(90)=0s       p(95)=0s
     http_req_duration..............: avg=47.19ms min=11.57ms med=37.58ms max=308.31ms p(90)=70.94ms  p(95)=99.63ms
     http_req_failed................: 0.00%   0 out of 253477
     http_req_receiving.............: avg=57.09µs min=0s      med=0s      max=8.23ms   p(90)=285.39µs p(95)=527.8µs
     http_req_sending...............: avg=18.67µs min=0s      med=0s      max=5.12ms   p(90)=0s       p(95)=0s
     http_req_waiting...............: avg=47.12ms min=10.26ms med=37.49ms max=307.15ms p(90)=70.87ms  p(95)=99.59ms
     http_reqs......................: 253477  2111.820179/s
     iteration_duration.............: avg=47.32ms min=13.58ms med=37.73ms max=308.31ms p(90)=71.1ms   p(95)=99.75ms
     iterations.....................: 253477  2111.820179/s
     vus............................: 100     min=100              max=100

running (2m00.0s), 000/100 VUs, 253477 complete and 0 interrupted iterations
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

     checks.........................: 100.00% 330481 out of 330481
     http_req_blocked...............: avg=5.66µs  min=0s     med=0s      max=11.66ms  p(90)=0s      p(95)=0s
     http_req_connecting............: avg=824ns   min=0s     med=0s      max=7.51ms   p(90)=0s      p(95)=0s
     http_req_duration..............: avg=72.48ms min=6.56ms med=69.85ms max=241.67ms p(90)=85.48ms p(95)=92.78ms
     http_req_failed................: 0.00%   0 out of 330481
     http_req_receiving.............: avg=60.15µs min=0s     med=0s      max=8.07ms   p(90)=432.7µs p(95)=530.5µs
     http_req_sending...............: avg=18.68µs min=0s     med=0s      max=8.05ms   p(90)=0s      p(95)=0s
     http_req_waiting...............: avg=72.4ms  min=5.74ms med=69.77ms max=241.67ms p(90)=85.41ms p(95)=92.71ms
     http_reqs......................: 330481  2752.537886/s
     iteration_duration.............: avg=72.61ms min=6.56ms med=69.98ms max=241.67ms p(90)=85.63ms p(95)=92.94ms
     iterations.....................: 330481  2752.537886/s
     vus............................: 200     min=200              max=200

running (2m00.1s), 000/200 VUs, 330481 complete and 0 interrupted iterations
default ✓ [======================================] 200 VUs  2m0s
```

### Test 4

The test script is [`test_url_create.js`](test_url_create.js). It creates a new URL and checks if the response status is 200.
It was run by 1000 VU concurrently for 2 minutes.

The results are as follows:

```bash
     execution: local
        script: ./test_url_create.js

     scenarios: (100.00%) 1 scenario, 200 max VUs, 2m30s max duration (incl. graceful stop):
              * default: 200 looping VUs for 2m0s (gracefulStop: 30s)

     checks.........................: 100.00%  324233 out of 324233
     http_req_blocked...............: avg=47.26µs  min=0s      med=0s       max=109.98ms p(90)=0s       p(95)=0s
     http_req_connecting............: avg=41.37µs  min=0s      med=0s       max=64.63ms  p(90)=0s       p(95)=0s
     http_req_duration..............: avg=369.2ms  min=0s      med=352.92ms max=754.41ms p(90)=464.7ms  p(95)=504.43ms
     http_req_failed................: 0.00%   0 out of 324233
     http_req_receiving.............: avg=61.31µs  min=0s      med=0s       max=11.68ms  p(90)=504.1µs  p(95)=532.4µs
     http_req_sending...............: avg=39.66µs  min=0s      med=0s       max=49.92ms  p(90)=0s       p(95)=0s
     http_req_waiting...............: avg=369.1ms  min=0s      med=352.84ms max=753.79ms p(90)=464.52ms p(95)=504.34ms
     http_reqs......................: 324233  2691.127018/s
     iteration_duration.............: avg=370.44ms min=31.01ms med=353.19ms max=802.57ms p(90)=465.99ms p(95)=504.92ms
     iterations.....................: 324233  2691.127018/s
     vus............................: 1000    min=1000             max=1000

running (2m00.5s), 0000/1000 VUs, 324233 complete and 0 interrupted iterations
default ✓ [======================================] 1000 VUs  2m0s
```

## Tests for Url Redirection

For this test to run the code had to be modified to return the original URL in the response. As for some reason redirection counting was not working.

### Test 1

The test script is [`test_url_redirect.js`](test_url_redirect.js). It redirects to the original URL and checks if the response status is 302.
It was run by 1 VU for 30 seconds.

The results are as follows:

```bash
     execution: local
        script: ./test_url_redirect.js

     scenarios: (100.00%) 1 scenario, 1 max VUs, 1m0s max duration (incl. graceful stop):
              * default: 1 looping VUs for 30s (gracefulStop: 30s)

     checks.........................: 100.00% 25053 out of 25053
     http_req_blocked...............: avg=3.48µs  min=0s      med=0s     max=9.52ms  p(90)=0s       p(95)=0s
     http_req_connecting............: avg=20ns    min=0s      med=0s     max=519.9µs p(90)=0s       p(95)=0s
     http_req_duration..............: avg=1.09ms  min=0s      med=1.04ms max=56ms    p(90)=1.59ms   p(95)=2.15ms
     http_req_failed................: 0.00%   0 out of 25053
     http_req_receiving.............: avg=83.7µs  min=0s      med=0s     max=5ms     p(90)=509.75µs p(95)=527.2µs
     http_req_sending...............: avg=12.64µs min=0s      med=0s     max=2.93ms  p(90)=0s       p(95)=0s
     http_req_waiting...............: avg=1ms     min=0s      med=1.03ms max=56ms    p(90)=1.55ms   p(95)=2.02ms
     http_reqs......................: 25053   835.096693/s
     iteration_duration.............: avg=1.18ms  min=107.6µs med=1.04ms max=56ms    p(90)=1.65ms   p(95)=2.27ms
     iterations.....................: 25053   835.096693/s
     vus............................: 1       min=1              max=1

running (0m30.0s), 0/1 VUs, 25053 complete and 0 interrupted iterations
default ✓ [======================================] 1 VUs  30s
```

### Test 2

The test script is [`test_url_redirect.js`](test_url_redirect.js). It redirects to the original URL and checks if the response status is 302.
It was run by 100 VU concurrently for 2 minutes.

The results are as follows:

```bash
     execution: local
        script: ./test_url_redirect.js

     scenarios: (100.00%) 1 scenario, 100 max VUs, 2m30s max duration (incl. graceful stop):
              * default: 100 looping VUs for 2m0s (gracefulStop: 30s)

     checks.........................: 100.00% 170563 out of 170563
     http_req_blocked...............: avg=3.88µs  min=0s      med=0s      max=21.76ms  p(90)=0s      p(95)=0s
     http_req_connecting............: avg=32ns    min=0s      med=0s      max=1.59ms   p(90)=0s      p(95)=0s
     http_req_duration..............: avg=12.39ms min=521.9µs med=10.56ms max=229.62ms p(90)=20.85ms p(95)=26.06ms
     http_req_failed................: 0.00%   0 out of 170563
     http_req_receiving.............: avg=59.1µs  min=0s      med=0s      max=41.03ms  p(90)=0s      p(95)=515.7µs
     http_req_sending...............: avg=13.49µs min=0s      med=0s      max=42.15ms  p(90)=0s      p(95)=0s
     http_req_waiting...............: avg=12.32ms min=507.6µs med=10.51ms max=229.62ms p(90)=20.75ms p(95)=25.92ms
     http_reqs......................: 170563  1421.358333/s
     iteration_duration.............: avg=12.49ms min=528.5µs med=10.65ms max=231.33ms p(90)=20.97ms p(95)=26.18ms
     iterations.....................: 170563  1421.358333/s
     vus............................: 100     min=100              max=100

running (2m00.0s), 000/100 VUs, 170563 complete and 0 interrupted iterations
default ✓ [======================================] 100 VUs  2m0s
```

### Test 3

The test script is [`test_url_redirect.js`](test_url_redirect.js). It redirects to the original URL and checks if the response status is 302.
It was run by 200 VU concurrently for 2 minutes.

Well this test was not run as the server was not able to handle the load of 200 VUs. The server was crashing and the response time was very high.
