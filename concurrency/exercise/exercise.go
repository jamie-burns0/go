package exercise

import "fmt"

func generate(genChan chan<- int) {
	for i := 2; i < 100; i++ {
		genChan<- i
	}
}


func filter(prime int, upstream <-chan int, downstream chan<- int) {
	for {
		n := <-upstream
		if n % prime == 0 {
			continue // drop this non-prime
		}
		fmt.Printf("filter[%v]: sending %v to downstream filter\n", prime, n)
		if downstream != nil {
			downstream<- n
			continue // get our next n
		}
		go filter(n, )
	}
}


func Launch() {
	genChan := make(chan int)
	done := make(chan bool)

	go generate(genChan)
	go filter(2, genChan, nil)

	<-done
}