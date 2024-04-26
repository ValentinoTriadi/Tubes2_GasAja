package main

import (
	"sync"
	"time"
)

// Buat struktur untuk token bucket
type TokenBucket struct {
	tokens int
	mutex  sync.Mutex
}

// Buat fungsi untuk mengkonsumsi token
func (tb *TokenBucket) Consume() bool {
	tb.mutex.Lock()
	defer tb.mutex.Unlock()

	if tb.tokens > 0 {
		tb.tokens--
		return true
	} else {
		return false
	}
}

// Buat fungsi untuk menambahkan token secara berkala
func (tb *TokenBucket) AddTokens(rate int, period time.Duration) {
	for range time.Tick(period) {
		tb.mutex.Lock()
		tb.tokens = rate
		tb.mutex.Unlock()
	}
}

var MaxToken = 30
// Inisialisasi token bucket
var tb = TokenBucket{tokens: MaxToken} // Atur jumlah token awal dan tingkat maksimum
