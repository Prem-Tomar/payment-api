package middlewares

import (
	"net/http"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
)

const (
	bucketCapacity = 5.0
	refillRate     = 1.0 // tokens per second
)

type tokenBucket struct {
	tokens     float64   // number of tokens
	lastRefill time.Time // last refil time
}

type RateLimiter struct {
	mu      sync.Mutex
	buckets map[string]*tokenBucket
}

func NewRateLimiter() *RateLimiter { // creating new map for each merchant
	return &RateLimiter{
		buckets: make(map[string]*tokenBucket),
	}
}

func (rl *RateLimiter) Middleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		merchantIDValue, exists := c.Get(MerchantIDKey)

		if !exists {
			c.Next()
			return
		}

		merchantID, ok := merchantIDValue.(string)
		if !ok || merchantID == "" {
			c.Next()
			return
		}

		if !rl.allow(merchantID) {
			c.Header("Retry-After", "1")
			c.JSON(http.StatusTooManyRequests, gin.H{
				"error": "rate limit exceeded",
			})
			c.Abort()
			return
		}

		c.Next()
	}
}

func (rl *RateLimiter) allow(key string) bool {
	rl.mu.Lock() // as map is not thread safe , to lock the map
	defer rl.mu.Unlock()

	now := time.Now()

	bucket, exists := rl.buckets[key]

	if !exists {
		rl.buckets[key] = &tokenBucket{
			tokens:     bucketCapacity - 1,
			lastRefill: now,
		}
		return true
	}

	elapsed := now.Sub(bucket.lastRefill).Seconds()

	bucket.tokens += elapsed * refillRate

	if bucket.tokens > bucketCapacity {
		bucket.tokens = bucketCapacity
	}

	bucket.lastRefill = now

	if bucket.tokens < 1 {
		return false
	}

	bucket.tokens--

	return true
}
