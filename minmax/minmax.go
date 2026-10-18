package minmax

import (
	"slices"
	"sort"
)

func MinMaxOld(num []int) (int, int) {
	min := num[0]
	max := num[0]

	for i := 0; i < len(num); i++ {
		if num[i] > max {
			max = num[i]
		}

		if min > num[i] {
			min = num[i]
		}
	}

	return min, max
}

func MinMax(num []int) (int, int) {
	min := num[0]
	max := num[0]

	for _, v := range num {
		if v > max {
			max = v
		}

		if min > v {
			min = v
		}
	}

	return min, max
}

func MinMaxSorted(num []int) (int, int) {
	sort.Ints(num)

	return num[0], num[len(num)-1]
}

func MinMaxSlices(num []int) (int, int) {
	return slices.Min(num), slices.Max(num)
}
