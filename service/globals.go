package service

import (
	"crypto/rand"
	"math/big"
)

// Generate Random number
func getRand(max int64) (int, error) {
    rand, err := rand.Int(rand.Reader, big.NewInt(max))
    return int(rand.Int64()), err
}
