#### Token Bucket algorithm for Rate limiting

- Simple implementation of Rate limiting HTTP requests 
- Uses GO package x/time/rate 

##### Test 
- Install `brew install vegeta`
- go run main.go
- `vegeta attack -duration=10s -rate=100 -targets=vegeta.conf | vegeta report` <br>Output for vegeta with rate 1 rps and burst is 5
```Requests      [total, rate, throughput]         1000, 100.11, 1.60
  Duration      [total, attack, wait]             9.99s, 9.989s, 411.503µs
  Latencies     [min, mean, 50, 90, 95, 99, max]  245.648µs, 596.287µs, 552.175µs, 870.758µs, 994.79µs, 1.556ms, 8.993ms
  Bytes In      [total, mean]                     18032, 18.03
  Bytes Out     [total, mean]                     0, 0.00
  Success       [ratio]                           1.60%
  Status Codes  [code:count]                      200:16  429:984  
  Error Set:
  429 Too Many Requests```