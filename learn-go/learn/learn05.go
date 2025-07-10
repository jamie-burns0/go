package learn

import (
	"fmt"
	"math/rand"
	"sort"
)

func Learn05_01() {

	fmt.Println("learn05 01")
	values := values(5)
	sum := sum(values)
	fmt.Printf("sum is %v\n", sum)
	sortedValues := reverseSort(values)
	fmt.Printf("reverse sorted values: %v", sortedValues)
}

func values(size int) []int {

	values := make([]int, size)

	for i := 0; i < size; i++ {
		values[i] = rand.Intn(100)
	}

	return values
}

// when passed a slice, sum2 will just
// use the slice. This is as efficient
// as sum(values []int)
func sum2(values ...int) int {
	sum := 0
	for _, v := range values {
		sum += v
	}
	return sum;
}

func sum(values []int) int {
	return sum2(values...)
}

func reverseSort(values []int) []int {

	sort.Slice(values, func(a, b int) bool {
		return values[a] > values[b]
	})
	return values
}