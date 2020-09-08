#### Token Bucket algorithm for Rate limiting

- Simple implementation of Rate limiting HTTP requests 
- Uses GO package x/time/rate 

##### Test 
- Install `brew install vegeta`
- go run main.go
- `vegeta attack -duration=10s -rate=100 -targets=vegeta.conf | vegeta report`