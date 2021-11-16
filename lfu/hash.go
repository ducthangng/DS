package lfu

import (
	"crypto/sha1"
	"encoding/hex"
	"hash/fnv"
)

// Hash and limit to size
func hash(s string) uint32 {
	h := fnv.New32a()
	h.Write([]byte(s))
	return (h.Sum32())
}

func hash_string(key string) string {
	h := sha1.New()
	h.Write([]byte(key))
	sha1_hash := hex.EncodeToString(h.Sum(nil))

	return sha1_hash
}
