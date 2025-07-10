package learn

import (
	"fmt"
	"math/rand"
	"sort"
)

func A() {

	var s1 = make([]int, 0, 3)

	s1 = append(s1, 1, 4, 3, 2, 5)
    s1 = append(s1, 44, 32, 19)

	for n := 0; n < 100; n++ {
		s1 = append(s1, rand.Intn(100))
	}

	sort.Slice(s1, func(i, j int) bool {
		return s1[i] > s1[j]
	})

	i, found := sort.Find(len(s1), func (i int) int {
		if( s1[i] < 50 ) {
			return -1
		} else if( s1[i] > 50 ) {
			return 1
		} else {
			return 0
		}
	})

	if( found ) {
		fmt.Println( "50 fount at position ", i)
	} else {
		fmt.Println( "50 not found")
	}

	fmt.Println(s1)
}

func compare(candidate int, target int) int {
	if( candidate < target ) {
			return -1
		} else if( candidate > target ) {
			return 1
		} else {
			return 0
		}
}