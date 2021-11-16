package tinyLFU

import (
	"crypto/sha1"
	"encoding/hex"
	"hash/fnv"
)

func Hash32(val []byte) int {
	hashfunc := fnv.New32a()
	hashfunc.Write(val)

	n := int(hashfunc.Sum32())

	return (n % size)
}

func Hash64(val []byte) int {
	hashfunc := fnv.New64a()
	hashfunc.Write(val)

	n := int(hashfunc.Sum64())

	return (n % size)
}

func HashString(key string) string {
	h := sha1.New()
	h.Write([]byte(key))
	sha1_hash := hex.EncodeToString(h.Sum(nil))

	return sha1_hash
}
