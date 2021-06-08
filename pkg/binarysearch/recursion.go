package binarysearch

func (sn sortedNumbers) containsRec(value int) bool {
	return sn.findRec(value, 0, len(sn)-1)
}

func (sn sortedNumbers) findRec(value, start, end int) bool {
	switch {
	case start > end:
		return false
	case start == end:
		return sn[start] == value
	case value < sn[start] || value > sn[end]:
		return false
	case value == sn[start] || value == sn[end]:
		return true
	}
	middle := (start + end) / 2
	switch {
	case sn[middle] == value:
		return true
	case value < sn[middle]:
		return sn.findRec(value, start, middle-1)
	case sn[middle] < value:
		return sn.findRec(value, middle+1, end)
	}
	panic(1)
}
