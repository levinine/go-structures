package binarysearch

import (
	"math/rand"
	"testing"
	"time"
)

func BenchmarkBinarySearchRecursive(b *testing.B) {
	b.StopTimer()
	b.ReportAllocs()
	sn := generate()
	rand.Seed(time.Now().UnixNano())
	value := rand.Intn(200000)
	b.StartTimer()
	_ = sn.containsRec(value)
}
