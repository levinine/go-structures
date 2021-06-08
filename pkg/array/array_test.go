package array

import (
	"fmt"
	"testing"
)

func BenchmarkBasicSorting(b *testing.B) {
	b.StopTimer()
	b.ReportAllocs()
	n := 1000000
	mn := generate(n)
	b.StartTimer()
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if mn[i] > mn[j] {
				mn[i], mn[j] = mn[j], mn[i]
			}
		}
	}
}

/*
10^1 -> 1000000000	         0.0000006 ns/op	       0 B/op	       0 allocs/op
10^2 -> 1000000000	         0.0000100 ns/op	       0 B/op	       0 allocs/op
10^3 -> 1000000000	         0.0007214 ns/op	       0 B/op	       0 allocs/op
10^4 -> 1000000000	         0.0828300 ns/op	       0 B/op	       0 allocs/op
10^5 ->          1	       11991271995 ns/op	       0 B/op	       0 allocs/op
*/

func TestCopySorted(t *testing.T) {
	mn := generate(30)
	//mn := mixedNumbers{5, 0, 7, 8, 1, 2, 6, 3, 4, 9}
	fmt.Printf("%v\n", mn)
	sorted := copySorted(mn)
	fmt.Printf("%v\n", sorted)
}

func BenchmarkCopySorted(b *testing.B) {
	b.StopTimer()
	b.ReportAllocs()
	mn := generate(10000)
	b.ResetTimer()
	b.StartTimer()
	_ = copySorted(mn)
}

/*
10^1 -> 1000000000	         0.0000018 ns/op	       0 B/op	       0 allocs/op
10^2 -> 1000000000	         0.0000113 ns/op	       0 B/op	       0 allocs/op
10^3 -> 1000000000	         0.0001708 ns/op	       0 B/op	       0 allocs/op
10^4 -> 1000000000	         0.0152500 ns/op	       0 B/op	       0 allocs/op
10^5 ->          1	        1494798326 ns/op	 5537392 B/op	      59 allocs/op
*/
