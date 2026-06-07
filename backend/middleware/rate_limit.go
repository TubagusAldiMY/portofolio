package middleware

import (
	"net/http"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/time/rate"
)

type ipEntry struct {
	limiter  *rate.Limiter
	lastSeen time.Time
}

// RateLimiter tracks per-IP token bucket limiters.
type RateLimiter struct {
	mu  sync.Mutex
	ips map[string]*ipEntry
	r   rate.Limit
	b   int
	ttl time.Duration
}

// NewRateLimiter creates a per-IP limiter with r tokens/sec and burst b.
// Old entries are evicted after ttl (default 10 min) to prevent memory growth.
func NewRateLimiter(r rate.Limit, b int) *RateLimiter {
	rl := &RateLimiter{
		ips: make(map[string]*ipEntry),
		r:   r,
		b:   b,
		ttl: 10 * time.Minute,
	}
	go rl.cleanup()
	return rl
}

func (rl *RateLimiter) cleanup() {
	ticker := time.NewTicker(5 * time.Minute)
	defer ticker.Stop()
	for range ticker.C {
		rl.mu.Lock()
		cutoff := time.Now().Add(-rl.ttl)
		for ip, e := range rl.ips {
			if e.lastSeen.Before(cutoff) {
				delete(rl.ips, ip)
			}
		}
		rl.mu.Unlock()
	}
}

func (rl *RateLimiter) get(ip string) *rate.Limiter {
	rl.mu.Lock()
	defer rl.mu.Unlock()

	e, ok := rl.ips[ip]
	if !ok {
		e = &ipEntry{limiter: rate.NewLimiter(rl.r, rl.b)}
		rl.ips[ip] = e
	}
	e.lastSeen = time.Now()
	return e.limiter
}

// Handler returns a Gin middleware that enforces the rate limit.
func (rl *RateLimiter) Handler() gin.HandlerFunc {
	return func(c *gin.Context) {
		if !rl.get(c.ClientIP()).Allow() {
			c.AbortWithStatusJSON(http.StatusTooManyRequests, gin.H{
				"error": "Too many requests. Please try again later.",
			})
			return
		}
		c.Next()
	}
}

func RateLimitPerMinute(requestsPerMinute, burst int) gin.HandlerFunc {
	if requestsPerMinute <= 0 {
		requestsPerMinute = 1
	}
	if burst <= 0 {
		burst = 1
	}

	interval := time.Minute / time.Duration(requestsPerMinute)
	return NewRateLimiter(rate.Every(interval), burst).Handler()
}
