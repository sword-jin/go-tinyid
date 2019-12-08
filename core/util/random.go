package util

import (
	"math/rand"
	"time"
)

func RandomInt(max int) int {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return r.Intn(max + 1)
}
