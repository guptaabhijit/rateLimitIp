package main

import (
	"time"

	"log"
	"net/http"
	"fmt"
)

var limiter = NewIPRateLimiter(1, 5)

func (i *IPRateLimiter) cleanupVisitors() {

	fmt.Printf(YELLOWSTART)
	log.Println("cleanupVisitors called")
	fmt.Printf(YELLOWEND)


	for {
		time.Sleep(time.Minute)
		i.mu.Lock()

		for ip, v := range i.ips {
			
			if time.Since(v.lastSeen) > 1*time.Minute {
				delete(i.ips, ip)

				log.Printf(REDSTART)
				log.Printf("IP deleted from map list: %s\n",ip)
				log.Println(len(i.ips))
				fmt.Printf(REDEND)
			}
		}
		i.mu.Unlock()
	}
}


func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", okHandler)
	go limiter.cleanupVisitors()
	if err := http.ListenAndServe(":8888", limitMiddleware(mux)); err != nil {
		log.Fatalf("unable to start server: %s", err.Error())
	}
}

func okHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Response from server"))
}


func limitMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		limiter := limiter.GetLimiter(r.RemoteAddr)
		if !limiter.Allow() {
			http.Error(w, http.StatusText(http.StatusTooManyRequests), http.StatusTooManyRequests)
			return
		}

		next.ServeHTTP(w, r)
	})
}
