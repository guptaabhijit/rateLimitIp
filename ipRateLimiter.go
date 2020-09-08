package main

import(
	"golang.org/x/time/rate"
	"sync"
	"log"
	"time"
    "fmt"
)

type visitor struct{
	limiter  *rate.Limiter
	lastSeen time.Time
}


// IPRateLimiter .
type IPRateLimiter struct {
	ips map[string]*visitor
	mu  *sync.RWMutex
	r   rate.Limit
	b   int
}


// NewIPRateLimiter .
func NewIPRateLimiter(r rate.Limit, b int) *IPRateLimiter {
	i := &IPRateLimiter{
		ips: make(map[string]*visitor),
		mu:  &sync.RWMutex{},
		r:   r,
		b:   b,
	}

	return i
}

func (i *IPRateLimiter) AddIP(ip string) *rate.Limiter {
	i.mu.Lock()
	defer i.mu.Unlock()

	l := rate.NewLimiter(i.r, i.b)

	i.ips[ip] = &visitor{l, time.Now()}

	return l
}

func (i *IPRateLimiter) GetLimiter(ip string) *rate.Limiter {
	i.mu.Lock()
	l, exists := i.ips[ip]
	fmt.Printf(BLUESTART)
	if !exists {
		i.mu.Unlock()
		log.Printf("IP doesnt exist: %s \n", ip)
		return i.AddIP(ip)
	}

	log.Printf("IP exist: %s\n",ip)
	fmt.Printf(BLUEEND)
	i.mu.Unlock()

	l.lastSeen = time.Now()

	return l.limiter
}

