package utils

import (
  "math/rand"
  "time"
)

var rng = rand.New(rand.NewSource(time.Now().UnixNano()))

func RandomNumbers() []int {
  return []int{
    rng.Intn(50),
    rng.Intn(50),
    rng.Intn(50),
  }
}
