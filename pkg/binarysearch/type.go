package binarysearch

type sortedNumbers []int

func generate() sortedNumbers {
	sn := make(sortedNumbers, 100000)
	for i, j := 0, 0; i < 100000; i, j = i+1, j+2 {
		sn[i] = j
	}
	return sn
}
