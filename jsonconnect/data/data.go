package data

import (
	"math/rand"
)

// Data wraps a number.
type Data struct {
	Number int64 `json:",int"`
}

// GetRandom creates random data.
func GetRandom() Data {
	return Data{rand.Int63()}
}
