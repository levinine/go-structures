package multimap

import (
	"errors"
	"math/rand"
	"time"
)

const (
	levelOneDivisor   uint64 = 1e17
	levelTwoDivisor   uint64 = 1e15
	levelThreeDivisor uint64 = 1e13
	iinFiller         uint64 = 1e9
)

type CardInfo struct {
	value string
}

// multimap holds a map where key is first 2 digits of IIN
type multimap struct {
	maps map[uint64]*subMap1
}

// subMap1 holds a map where key is third and forth digits of IIN
type subMap1 struct {
	maps map[uint64]*subMap2
}

// subMap2 holds a map where key is fifth and sixth digits of IIN
type subMap2 struct {
	arrays map[uint64]*subArray
}

// subArray holds all ranges whos numbers start with first 6 digits needed to get this far.
// Idealy, these arrays should not have more than 100 elements but in some cases it is possible
// there will be arrays with up to 120 elements. Ranges are sorted in ascending order.
type subArray struct {
	ranges []*infoRange
}

type infoRange struct {
	start uint64
	end   uint64
	info  *CardInfo
}

// Search

func (mm *multimap) get(number uint64) (*CardInfo, bool) {
	// Calculate keys for top level map, sub map 1 and sub map 2
	num := number
	key := num / levelOneDivisor
	num %= levelOneDivisor
	subKey := num / levelTwoDivisor
	num %= levelTwoDivisor
	subSubKey := num / levelThreeDivisor

	if sm1, ok := mm.maps[key]; ok {
		ci := sm1.get(number, subKey, subSubKey)
		return ci, ci != nil
	} else {
		return nil, false
	}
}

func (sm1 *subMap1) get(number, key, subKey uint64) *CardInfo {
	if sm2, ok := sm1.maps[key]; ok {
		return sm2.get(number, subKey)
	} else {
		return nil
	}
}

func (sm2 *subMap2) get(number, key uint64) *CardInfo {
	if a, ok := sm2.arrays[key]; ok {
		ci, _ := a.find(number)
		return ci
	} else {
		return nil
	}
}

// Perform binary search algorithm, but without using recursion
// in order to make it faster
func (sa *subArray) find(number uint64) (*CardInfo, int) {
	// allocate all we need at the start
	start := 0
	end := len(sa.ranges) - 1

	var (
		index int
		ir    *infoRange
	)

	for {
		switch {
		case start == end:
			// make last check
			ir = sa.ranges[start]
			if ir.start <= number && number <= ir.end {
				return ir.info, start
			} else {
				return nil, -1
			}
		case start > end:
			// when this happens range is not found
			return nil, -1
		}

		index = (start + end) / 2 //nolint:gomnd
		ir = sa.ranges[index]

		switch {
		case ir.start > number:
			end = index - 1
			continue
		case ir.end < number:
			start = index + 1
			continue
		default:
			return ir.info, index
		}
	}
}

// Insert

func (mm *multimap) insert(low, high uint64, ci *CardInfo) error {
	key := low / levelOneDivisor
	if sm1, ok := mm.maps[key]; ok {
		return sm1.insert(low, high, ci)
	}

	mm.maps[key] = &subMap1{
		maps: map[uint64]*subMap2{},
	}

	return mm.maps[key].insert(low, high, ci)
}

func (sm1 *subMap1) insert(low, high uint64, ci *CardInfo) error {
	key := (low % levelOneDivisor) / levelTwoDivisor
	if sm2, ok := sm1.maps[key]; ok {
		return sm2.insert(low, high, ci)
	}

	sm1.maps[key] = &subMap2{
		arrays: map[uint64]*subArray{},
	}

	return sm1.maps[key].insert(low, high, ci)
}

func (sm2 *subMap2) insert(low, high uint64, ci *CardInfo) error {
	key := (low % levelTwoDivisor) / levelThreeDivisor
	if sa, ok := sm2.arrays[key]; ok {
		return sa.insert(low, high, ci)
	}

	sm2.arrays[key] = &subArray{
		ranges: []*infoRange{},
	}

	return sm2.arrays[key].insert(low, high, ci)
}

func (sa *subArray) insert(low, high uint64, ci *CardInfo) error {
	// check if range borders are part of some already inserted range
	if _, index := sa.find(low); index != -1 {
		return errors.New("cannot start inside another card range")
	}

	if _, index := sa.find(high); index != -1 {
		return errors.New("cannot end inside another card range")
	}

	// count all ranges ending before `low`
	leftCount := 0

	for i := 0; i < len(sa.ranges); i++ {
		if sa.ranges[i].start > high {
			break
		}
		leftCount++
	}

	// count all ranges ending after `high`
	rightCount := 0

	for i := len(sa.ranges) - 1; i >= 0; i-- {
		if sa.ranges[i].end < low {
			break
		}
		rightCount++
	}

	// if sum of leftCount and rightCount is less than full length of subArray
	// it means that new range envelops some of the already inserted ranges
	if leftCount+rightCount < len(sa.ranges) {
		return errors.New("canot overlap ranges")
	}
	// extend the array by one and shift all ranges that come after the new range
	// to the right by one position
	sa.ranges = append(sa.ranges, nil)
	for i := len(sa.ranges) - 2; i >= leftCount; i-- { //nolint:gomnd
		sa.ranges[i+1] = sa.ranges[i]
	}
	// save new range to its proper place
	sa.ranges[leftCount] = &infoRange{
		start: low,
		end:   high,
		info:  ci,
	}

	return nil
}

// Tear down the entire structure in order to help out GC
// by releasing all memory in one pass
func (mm *multimap) tearDown() {
	for mmK, mmV := range mm.maps {
		for sm1K, sm1V := range mmV.maps {
			for sm2K, sm2V := range sm1V.arrays {
				for i, sa := range sm2V.ranges {
					sa.info = nil
					sm2V.ranges[i] = nil
				}

				sm2V.ranges = nil
				sm1V.arrays[sm2K] = nil
			}

			sm1V.arrays = nil
			mmV.maps[sm1K] = nil
		}

		mmV.maps = nil
		mm.maps[mmK] = nil
	}

	mm.maps = nil
}

type input struct {
	from, to uint64
	value    string
}

// 1234567890123456789
// 1000000000000000000
// 100000
//      10000000000000
func generateSortedInputs(n int) []input {
	list := make([]input, n)
	for i := 0; i < n; i++ {
		list[i] = input{
			from:  (uint64(i) * 30000000000000) + 1000000000000000000,
			to:    (uint64(i) * 30000000000000) + 1000000000000000002,
			value: "QWERTY",
		}
	}
	return list
}

func generateMixedNumbers(n int) []input {
	list := generateSortedInputs(n)
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(n, func(i, j int) { list[i], list[j] = list[j], list[i] })
	return list
}
