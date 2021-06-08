package tree

import (
	"math/rand"
	"testing"
	"time"
)

func BenchmarkUnsortedListToTree(b *testing.B) {
	b.StopTimer()
	b.ReportAllocs()
	n := 100000
	list := generateMixedNumbers(n)
	b.ResetTimer()
	b.StartTimer()
	root := node{
		from: list[0].from,
		to:   list[0].to,
		v:    &nodeValue{value: list[0].value},
	}
	for i := 1; i < n; i++ {
		root.add(list[i].from, list[i].to, list[i].value)
	}
}

/*
10^1 -> 1000000000	         0.0000016 ns/op	       0 B/op	       0 allocs/op
10^2 -> 1000000000	         0.0000096 ns/op	       0 B/op	       0 allocs/op
10^3 -> 1000000000	         0.0001157 ns/op	       0 B/op	       0 allocs/op
10^4 -> 1000000000	         0.0018010 ns/op	       0 B/op	       0 allocs/op
10^5 -> 1000000000	         0.0387200 ns/op	       0 B/op	       0 allocs/op
*/

func BenchmarkSortedListToTree(b *testing.B) {
	b.StopTimer()
	b.ReportAllocs()
	n := 100000
	list := generateSortedInputs(n)
	b.ResetTimer()
	b.StartTimer()
	root := node{
		from: list[0].from,
		to:   list[0].to,
		v:    &nodeValue{value: list[0].value},
	}
	for i := 1; i < n; i++ {
		root.add(list[i].from, list[i].to, list[i].value)
	}
}

/*
10^1 -> 1000000000	         0.0000016 ns/op	       0 B/op	       0 allocs/op
10^2 -> 1000000000	         0.0000457 ns/op	       0 B/op	       0 allocs/op
10^3 -> 1000000000	         0.0036550 ns/op	       0 B/op	       0 allocs/op
10^4 -> 1000000000	         0.3809000 ns/op	       0 B/op	       0 allocs/op
10^5 ->          1	 (50s) 50712742316 ns/op	 6410688 B/op	  200016 allocs/op
*/

func BenchmarkReadingFromUnsortedListTree(b *testing.B) {
	b.StopTimer()
	b.ReportAllocs()
	n := 1000
	list := generateSortedInputs(n)
	root := node{
		from: list[0].from,
		to:   list[0].to,
		v:    &nodeValue{value: list[0].value},
	}
	for i := 1; i < n; i++ {
		root.add(list[i].from, list[i].to, list[i].value)
	}
	rand.Seed(time.Now().UnixNano())
	number := uint64(rand.Intn(n * 3))
	b.ResetTimer()
	b.StartTimer()
	root.get(number)
}

/*
10^1 -> 1000000000	         0.0000002 ns/op	       0 B/op	       0 allocs/op
10^2 -> 1000000000	         0.0000003 ns/op	       0 B/op	       0 allocs/op
10^3 -> 1000000000	         0.0000065 ns/op	       0 B/op	       0 allocs/op
10^4 -> 1000000000	         0.0000511 ns/op	       0 B/op	       0 allocs/op
10^5 -> 1000000000	         0.0006820 ns/op	       0 B/op	       0 allocs/op
*/
/*
func TestMyTest(t *testing.T) {
	n := 100000
	list := generateSortedNumbers(n)
	var a, b, c time.Time
	a = time.Now()
	root := node{v: &nodeValue{value: list[0]}}
	for i := 1; i < n; i++ {
		root.add(list[i])
	}
	b = time.Now()
	root.contains(n - 1)
	c = time.Now()
	fmt.Println(b.UnixNano() - a.UnixNano())
	fmt.Println(c.UnixNano() - b.UnixNano())
}
*/
