package middlewares

import (
	"math"
	"net/http"
	"strconv"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
)

const (
	bucketCapacity = 5.0
	refillRate     = 1.0 // tokens per second
)

const (
	reateLimitHeader    = "X-RateLimit-Limit"
	rateRemainingHeader = "X-RateLimit-Remaining"
	retryAfterHeader    = "Retry-After"
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
		allowed, remaining, retryAfter := rl.allow(merchantID)
		c.Header(reateLimitHeader, strconv.Itoa(int(bucketCapacity)))
		c.Header(
			rateRemainingHeader,
			strconv.Itoa(remaining),
		)

		if !allowed {
			c.Header(
				retryAfterHeader,
				strconv.Itoa(retryAfter),
			)

			c.JSON(http.StatusTooManyRequests, gin.H{
				"error": "rate limit exceeded",
			})
			c.Abort()
			return
		}

		c.Next()
	}
}

func (rl *RateLimiter) allow(key string) (allowed bool, remaining int, retryAfter int) {
	rl.mu.Lock() // as map is not thread safe , to lock the map
	defer rl.mu.Unlock()

	now := time.Now()

	bucket, exists := rl.buckets[key]

	if !exists {
		bucket = &tokenBucket{
			tokens:     bucketCapacity - 1,
			lastRefill: now,
		}
		rl.buckets[key] = bucket
	}

	elapsed := now.Sub(bucket.lastRefill).Seconds()

	bucket.tokens += elapsed * refillRate

	if bucket.tokens > bucketCapacity {
		bucket.tokens = bucketCapacity
	}

	bucket.lastRefill = now

	if bucket.tokens < 1 {
		waitSeconds := (1 - bucket.tokens) / refillRate

		return false, int(bucket.tokens), int(math.Ceil(waitSeconds))
	}

	bucket.tokens--

	return true, int(bucket.tokens), 0
}
