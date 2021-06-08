package array

import (
	"math/rand"
	"time"
)

type mixedNumbers []int

func generate(n int) mixedNumbers {
	mn := make(mixedNumbers, n)
	for i := 0; i < n; i++ {
		mn[i] = i
	}
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(n, func(i, j int) { mn[i], mn[j] = mn[j], mn[i] })
	return mn
}

func copySorted(mn mixedNumbers) []int {
	var sorted []int = nil
	n, i := 0, 0
	for _, v := range mn {
		switch {
		case sorted == nil:
			sorted = []int{v}
			n = 1
		case sorted[n-1] < v:
			sorted = append(sorted, v)
			n++
		case v < sorted[0]:
			sorted = append([]int{v}, sorted...)
			n++
		default:
			sorted = append(sorted, v)
			for i = n - 1; i >= 0; i-- {
				if sorted[i] > v {
					sorted[i+1] = sorted[i]
				} else {
					break
				}
			}
			sorted[i+1] = v
			n++
		}
	}
	return sorted
}
