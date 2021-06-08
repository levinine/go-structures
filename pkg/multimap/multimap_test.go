package multimap

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func TestRun(t *testing.T) {
	var starts, end time.Time
	inputs := generateSortedInputs(100000)
	starts = time.Now()
	mm := multimap{
		maps: map[uint64]*subMap1{},
	}
	for _, input := range inputs {
		mm.insert(input.from, input.to, &CardInfo{value: input.value})
	}
	end = time.Now()
	fmt.Println(end.UnixNano() - starts.UnixNano())
}

func BenchmarkBuildMultimapFromSortedList(b *testing.B) {
	b.StopTimer()
	b.ReportAllocs()
	inputs := generateSortedInputs(10000)
	b.ResetTimer()
	b.StartTimer()
	mm := multimap{
		maps: map[uint64]*subMap1{},
	}
	for _, input := range inputs {
		mm.insert(input.from, input.to, &CardInfo{value: input.value})
	}
}

/*
10^1 -> 1000000000	         0.0000129 ns/op	       0 B/op	       0 allocs/op
10^2 -> 1000000000	         0.0000427 ns/op	       0 B/op	       0 allocs/op
10^3 -> 1000000000	         0.0002846 ns/op	       0 B/op	       0 allocs/op
10^4 -> 1000000000	         0.0027200 ns/op	       0 B/op	       0 allocs/op
10^5 -> 1000000000	         0.0307300 ns/op	       0 B/op	       0 allocs/op
*/

func BenchmarkBuildMultimapFromMixedList(b *testing.B) {
	b.StopTimer()
	b.ReportAllocs()
	inputs := generateMixedNumbers(10)
	b.ResetTimer()
	b.StartTimer()
	mm := multimap{
		maps: map[uint64]*subMap1{},
	}
	for _, input := range inputs {
		mm.insert(input.from, input.to, &CardInfo{value: input.value})
	}
}

/*
10^1 -> 1000000000	         0.0000095 ns/op	       0 B/op	       0 allocs/op
10^2 -> 1000000000	         0.0000409 ns/op	       0 B/op	       0 allocs/op
10^3 -> 1000000000	         0.0002934 ns/op	       0 B/op	       0 allocs/op
10^4 -> 1000000000	         0.0027720 ns/op	       0 B/op	       0 allocs/op
10^5 -> 1000000000	         0.0394000 ns/op	       0 B/op	       0 allocs/op
*/

func BenchmarkReadFromMultimap(b *testing.B) {
	b.StopTimer()
	b.ReportAllocs()
	n := 100000
	inputs := generateSortedInputs(n)
	mm := multimap{
		maps: map[uint64]*subMap1{},
	}
	for _, input := range inputs {
		mm.insert(input.from, input.to, &CardInfo{value: input.value})
	}
	rand.Seed(time.Now().UnixNano())
	number := uint64(rand.Intn(n * 3))
	b.ResetTimer()
	b.StartTimer()
	_, _ = mm.get(number)
}

/*
10^1 -> 1000000000	         0.0000002 ns/op	       0 B/op	       0 allocs/op
10^2 -> 1000000000	         0.0000002 ns/op	       0 B/op	       0 allocs/op
10^3 -> 1000000000	         0.0000002 ns/op	       0 B/op	       0 allocs/op
10^4 -> 1000000000	         0.0000002 ns/op	       0 B/op	       0 allocs/op
10^5 -> 1000000000	         0.0000003 ns/op	       0 B/op	       0 allocs/op
*/
