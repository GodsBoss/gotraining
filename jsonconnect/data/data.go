package data

import (
	"math/rand"
)

type Data struct {
	Number int64 `json:",int"`
}

func GetRandom() Data {
	return Data{rand.Int63()}
}
