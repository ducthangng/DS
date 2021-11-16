package lfu

import (
	"math/rand"
	"testing"
	"time"
)

// RandomNumber returns a random number.
func RandomNumber(min int64, max int64) int {
	return int(min + rand.Int63n(max-min+1))
}

// RandomString returns a random string
func RandomString(len int) string {
	bytes := make([]byte, len)

	for i := 0; i < len; i++ {
		bytes[i] = byte(RandomNumber(97, 122))
	}

	return string(bytes)
}

func BenchmarkLFUWithHash(b *testing.B) {
	rand.Seed(time.Now().UTC().UnixNano())

	// lfu := NewLFU()
	lfus := NewLFUS()

	for i := 0; i <= b.N; i++ {
		key := RandomString(8)

		// if err := lfu.add(key, key); err != nil {
		// 	panic(err)
		// }

		if err := lfus.add(key, key); err != nil {
			panic(err)
		}
	}

}

/*
LFUS
goos: darwin
goarch: amd64
pkg: lfu
cpu: Intel(R) Core(TM) i7-9750H CPU @ 2.60GHz
BenchmarkLFUWithHash-12    	  254838	     70137 ns/op	  513891 B/op	       5 allocs/op
PASS
ok  	lfu	18.414s
*/

/*
LFU
goos: darwin
goarch: amd64
pkg: lfu
cpu: Intel(R) Core(TM) i7-9750H CPU @ 2.60GHz
BenchmarkLFUWithHash-12    	  281365	     78179 ns/op	  566937 B/op	       5 allocs/op
PASS
ok  	lfu	22.483s
*/
