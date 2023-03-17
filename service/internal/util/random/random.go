package random

import (
	"math/rand"
	"time"
)

func init() {
	InitRandom()
}

func InitRandom() {
	nowTime := time.Now().UnixNano()
	rand.Seed(nowTime)
}
