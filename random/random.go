package random

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

const RetryMax = 3

func Intn(max int) (int, error) {
	var errCounter = 0
	randomIndex, err := rand.Int(rand.Reader, big.NewInt(int64(max)))
	if err != nil {
		for retry := 0; retry < RetryMax; retry++ {
			randomIndex, err = rand.Int(rand.Reader, big.NewInt(int64(max)))
			if err != nil {
				errCounter++
				if errCounter == 3 {
					return 0, fmt.Errorf("Too many errors in random generation")
				}
				continue
			}
			return int(randomIndex.Int64()), nil
		}
	}
	return int(randomIndex.Int64()), nil
}
