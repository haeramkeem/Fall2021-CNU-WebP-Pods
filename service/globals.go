package service

import (
	"crypto/rand"
	"math/big"
)

// Generate Random number
func getRand(max int) (int, error) {
    rand, err := rand.Int(rand.Reader, big.NewInt(int64(max)))
    return int(rand.Int64()), err
}
