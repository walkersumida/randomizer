package seed

import (
	"crypto/rand"
	"math"
	"math/big"
)

func Int64() int64 {
	seed, _ := rand.Int(rand.Reader, big.NewInt(math.MaxInt64))

	return seed.Int64()
}
