package lfu

import (
	"log"
	"math/rand"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
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

func TestLFU(t *testing.T) {
	rand.Seed(time.Now().UTC().UnixNano())

	lfus := NewLFU()
	// lfu := NewLFU()

	keys := []string{}

	for i := 0; i < 100; i++ {
		keys = append(keys, RandomString(8))
	}

	for i := 0; i < 10000; i++ {
		err := lfus.add(keys[i%100], keys[i%100])
		require.NoError(t, err)
	}

	for i := 0; i < 10000; i++ {
		val, err := lfus.get(keys[i%100])
		require.NoError(t, err)
		if val != nil {
			require.Equal(t, val.(string), keys[i%100])
		}
	}

}

func BenchmarkLFU_Int(b *testing.B) {
	rand.Seed(time.Now().UTC().UnixNano())

	lfus := NewLFU()
	keys := []string{}

	for i := 0; i < 100; i++ {
		keys = append(keys, RandomString(8))
	}

	for i := 0; i <= b.N; i++ {
		err := lfus.add(keys[i%100], "random_value_to_retrieve")
		if err != nil {
			log.Println(err) // return full?
		}

		_, _ = lfus.get(keys[(i+rand.Intn(100))%100])
	}
}

/*
goos: darwin
goarch: amd64
pkg: algo/lfu
cpu: Intel(R) Core(TM) i7-9750H CPU @ 2.60GHz
BenchmarkLFU_Int-12    	  341947	     79766 ns/op	  687923 B/op	       2 allocs/op
PASS
ok  	algo/lfu	28.159s
*/
