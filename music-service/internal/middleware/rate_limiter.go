/*
so i need 2 structs one for storing how much max token i can have, how much at i have at present, fill rate
and other a hashmap for different ip to that struct
*/

package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"sync"
	"time"
)

type RateLimiter struct {
	currentTokens  float64
	maximumTokens  float64
	refillRate     float64
	lastRefillTime time.Time
	mutex		   sync.Mutex
}

type IPHash struct {
	limiters 	  map[string] *RateLimiter
	mutex    	  sync.Mutex
	maximumTokens float64
	refillRate    float64
}

func NewRateLimiter(maximumTokens float64, refillRate float64) *RateLimiter {
	return &RateLimiter{
		currentTokens:  maximumTokens,
		maximumTokens:  maximumTokens,
		refillRate:	    refillRate,
		lastRefillTime: time.Now(),
	}
}

func (r *RateLimiter) IsTooManyRequest() bool {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	now := time.Now()
	duration := now.Sub(r.lastRefillTime).Seconds()
	tokensToAdd := duration*r.refillRate
	r.currentTokens += tokensToAdd

	if r.currentTokens > r.maximumTokens {
		r.currentTokens = r.maximumTokens
	}

	r.lastRefillTime = now

	if r.currentTokens >= 1 {
		r.currentTokens--
		return false
	}

	return true
}

func NewIPHash(maxTokens float64, refillRate float64) *IPHash {
	return &IPHash{
		limiters: 		make(map[string]*RateLimiter),
		maximumTokens:  maxTokens,
		refillRate:		refillRate,
	}
}

func (i *IPHash) GetRateLimiter(ip string) *RateLimiter {
	i.mutex.Lock()
	defer i.mutex.Unlock()

	limiter, exists := i.limiters[ip]
	if !exists {
		limiter = NewRateLimiter(i.maximumTokens, i.refillRate)
		i.limiters[ip] = limiter
	}
	return limiter
}

func RateLimitingMiddleware(ipLimiter *IPHash) gin.HandlerFunc {
	return func(c *gin.Context) {
		ip := c.ClientIP()
		RLimiter := ipLimiter.GetRateLimiter(ip)
		
		if RLimiter.IsTooManyRequest() == false {
			c.Next()
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Too many Request"})
			c.Abort()
			return
		}
	}
}
