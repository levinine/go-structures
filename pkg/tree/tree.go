package tree

import (
	"math/rand"
	"time"
)

type nodeValue struct {
	value string
}

type node struct {
	from, to uint64
	v        *nodeValue
	left     *node
	right    *node
}

func (n *node) add(from, to uint64, value string) {
	if n.to < from {
		if n.right == nil {
			n.right = &node{
				from: from,
				to:   to,
				v: &nodeValue{
					value: value,
				},
			}
		} else {
			n.right.add(from, to, value)
		}
	} else if to < n.from {
		if n.left == nil {
			n.left = &node{
				from: from,
				to:   to,
				v: &nodeValue{
					value: value,
				},
			}
		} else {
			n.left.add(from, to, value)
		}
	} else {
		panic(1)
	}
}

func (n *node) get(number uint64) *nodeValue {
	switch {
	case n.from <= number && number <= n.to:
		return n.v
	case number < n.from:
		return n.left.get(number)
	case n.to < number:
		return n.right.get(number)
	default:
		return nil
	}
}

type input struct {
	from, to uint64
	value    string
}

func generateSortedInputs(n int) []input {
	list := make([]input, n)
	for i := 0; i < n; i++ {
		list[i] = input{
			from:  uint64(i) * 3,
			to:    (uint64(i) * 3) + 2,
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
