# rust-go-benchmark

# Benchmarking empty Http handler
`ab -n 100000 -c 100 http://localhost:8080/`


## GO
```bash

Document Path:          /
Document Length:        13 bytes

Concurrency Level:      100
Time taken for tests:   11.878 seconds
Complete requests:      100000
Failed requests:        0
Total transferred:      13000000 bytes
HTML transferred:       1300000 bytes
Requests per second:    8419.27 [#/sec] (mean)
Time per request:       11.878 [ms] (mean)
Time per request:       0.119 [ms] (mean, across all concurrent requests)
Transfer rate:          1068.85 [Kbytes/sec] received

Connection Times (ms)
              min  mean[+/-sd] median   max
Connect:        0    0   0.2      0       1
Processing:     2   12   1.1     11      26
Waiting:        0    8   2.8      8      18
Total:          2   12   1.1     12      26

Percentage of the requests served within a certain time (ms)
  50%     12
  66%     12
  75%     12
  80%     12
  90%     13
  95%     14
  98%     15
  99%     16
 100%     26 (longest request)
 ```


## Rust
```bash

Document Path:          /
Document Length:        12 bytes

Concurrency Level:      100
Time taken for tests:   13.271 seconds
Complete requests:      100000
Failed requests:        0
Total transferred:      8800000 bytes
HTML transferred:       1200000 bytes
Requests per second:    7535.28 [#/sec] (mean)
Time per request:       13.271 [ms] (mean)
Time per request:       0.133 [ms] (mean, across all concurrent requests)
Transfer rate:          647.56 [Kbytes/sec] received

Connection Times (ms)
              min  mean[+/-sd] median   max
Connect:        0    0   0.3      0       1
Processing:     2   11   0.9     11      17
Waiting:        0    8   1.8      8      17
Total:          2   11   0.9     11      17

Percentage of the requests served within a certain time (ms)
  50%     11
  66%     11
  75%     12
  80%     12
  90%     12
  95%     13
  98%     14
  99%     14
 100%     17 (longest request)
```


# Benchmarking BFS search with loading new graph on each request
`ab -n 100000 -c 100 http://localhost:8080/graph`


## GO BFS
```bash
Server Hostname:        localhost
Server Port:            8080

Document Path:          /graph
Document Length:        10 bytes

Concurrency Level:      100
Time taken for tests:   11.156 seconds
Complete requests:      100000
Failed requests:        0
Total transferred:      12700000 bytes
HTML transferred:       1000000 bytes
Requests per second:    8964.14 [#/sec] (mean)
Time per request:       11.156 [ms] (mean)
Time per request:       0.112 [ms] (mean, across all concurrent requests)
Transfer rate:          1111.76 [Kbytes/sec] received

Connection Times (ms)
              min  mean[+/-sd] median   max
Connect:        0    0   0.2      0       3
Processing:     2   11   0.8     11      17
Waiting:        1    8   2.3      8      17
Total:          2   11   0.8     11      17

Percentage of the requests served within a certain time (ms)
  50%     11
  66%     11
  75%     11
  80%     11
  90%     12
  95%     13
  98%     13
  99%     14
 100%     17 (longest request)
```

## RUST BFS
```bash
Server Hostname:        localhost
Server Port:            8080

Document Path:          /graph
Document Length:        14 bytes

Concurrency Level:      100
Time taken for tests:   12.814 seconds
Complete requests:      100000
Failed requests:        0
Total transferred:      9000000 bytes
HTML transferred:       1400000 bytes
Requests per second:    7803.89 [#/sec] (mean)
Time per request:       12.814 [ms] (mean)
Time per request:       0.128 [ms] (mean, across all concurrent requests)
Transfer rate:          685.89 [Kbytes/sec] received

Connection Times (ms)
              min  mean[+/-sd] median   max
Connect:        0    0   0.3      0       1
Processing:     2   11   0.7     11      17
Waiting:        0    8   1.6      8      16
Total:          2   11   0.7     11      17

Percentage of the requests served within a certain time (ms)
  50%     11
  66%     11
  75%     11
  80%     11
  90%     11
  95%     12
  98%     12
  99%     13
 100%     17 (longest request)
 ```