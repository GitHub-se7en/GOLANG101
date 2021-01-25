package common

import (
	"github.com/prometheus/common/log"
	"reflect"
)

// Swap two array values given their indices.
func Swap(v interface{}, i, j int) {
	switch a := v.(type) {
	case []int:
		SwapInt(a, i, j)
	case []string:
		SwapString(a, i, j)
	}
}

// SwapInt two array values given their indices.
func SwapInt(a []int, i, j int) {
	tmp := a[i]
	a[i] = a[j]
	a[j] = tmp
}

// SwapString two array values given their indices.
func SwapString(a []string, i, j int) {
	tmp := a[i]
	a[i] = a[j]
	a[j] = tmp
}

// Equal checks if two input are deeply equal.
func Equal(expected, result interface{}) {
	if !reflect.DeepEqual(result, expected) {
		log.Error("should be %v instead of %v", expected, result)
	}
}

// Mimax returns min and max from a list of integers.
func Mimax(nums ...int) (int, int) {
	min, max := nums[0], nums[0]

	for _, num := range nums {
		if min > num {
			min = num
		}

		if max < num {
			max = num
		}
	}

	return min, max
}

// Min returns min from a list of integers.
func Min(nums ...int) int {
	min := nums[0]

	for _, num := range nums {
		if min > num {
			min = num
		}
	}

	return min
}

// Max returns max from a list of integers.
func Max(nums ...int) int {
	max := nums[0]

	for _, num := range nums {
		if max < num {
			max = num
		}
	}

	return max
}
