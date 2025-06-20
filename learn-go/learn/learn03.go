package learn

import (
	"fmt"
	"sort"
)

func Learn03_01() {

	m1 := make(map[string]int)

	m1["a"] = 1
	m1["b"] = 2

	var m2 map[string]int

	m2 = m1

	fmt.Println("m1: ", m1)
	fmt.Println("m2: ", m2)
	
	m2["c"] = 3
	fmt.Println("m1: ", m1)
	fmt.Println("m2: ", m2)
}

func Learn03_02() {
	
	i := 100

	var p1 *int = &i

	fmt.Println("  i: ", i)     // i returns i's value
	fmt.Println("*p1: ", *p1) // *p1 returns i's value

	fmt.Println(" &i: ", &i)   // &i returns the memory location holding the value of i
	fmt.Println(" p1: ", p1)   // p1 returns the memory location holding the value of i
	
	fmt.Println("&p1: ", &p1) // &pi returns the memory location holding the value of pointer 

	f1 := 234.5
	pointer := &f1
	*pointer = *pointer / 10
	fmt.Println("pointer: ", *pointer)
	fmt.Println("f1: ", f1)
}

func Learn03_03() {

	var ca [3]string // arrays are a fixed size - you cannot grow them

	ca[0] = "red"
	ca[1] = "green"
	ca[2] = "blue"

	fmt.Println("ca: ", ca)
	fmt.Println("ca[1]: ", ca[1])

	var na = [5]int {5,3,4,1,2}
	fmt.Println("na: ", na)

	fmt.Println("ca.len:", len(ca))
	fmt.Println("na.len:", len(na))
}

func Learn03_04() {

	var ar1 = [3]string{"red","green","blue"}
	var sl1 = []string{"red","green","blue"}

	fmt.Println("ar1: ", ar1)
	fmt.Printf("ar1.type %T\n", ar1)
	fmt.Println("sl1: ", sl1)
	fmt.Printf("sl1.type %T\n", sl1)

	var sl2 = make([]string, 0, 3)
	sl2 = append(sl2, "red", "yellow", "pink", "blue")

	fmt.Println("sl2: ", sl2)
	fmt.Printf("sl2.type %T\n", sl2)

	sl2 = append(sl2, "orange", "purple", "green")
	fmt.Println("sl2: ", sl2)

	// sort sl2 by string length
	sort.Slice(sl2, func(i, j int) bool {
		return len(sl2[i]) < len(sl2[j])
	})
	
	fmt.Println("sl2: ", sl2)

	n := 2
	// There is no standard library function to remove an item from a slice.
	// The idiomatic way is to use slicing and append, as shown below:
	sl2 = append(sl2[:n], sl2[n+1:]...)
	fmt.Println("sl2: ", sl2)
	
	// sort sl2 by string
	sort.Strings(sl2)
	fmt.Println("sl2: ", sl2)
}

func Learn03_05() {

	m1 := make(map[string]string)
	m1["a"] = "date"
	m1["d"] = "apple"
	m1["b"] = "cherry"
	m1["c"] = "banana"
	
	fmt.Println(m1)

	for k,v := range m1 {
		fmt.Printf("[%s,%s]\n", k, v)
	}

	f := m1["b"]
	fmt.Println(f)

	delete(m1, "b")
	fmt.Println(m1)
	
	m1["b"] = "cherry2"
	fmt.Println(m1)
	
	for k,v := range m1 {
		fmt.Printf("[%s,%s]\n", k, v)
	}

	fmt.Println("----")

	// make a copy of m1
	m2 := make(map[string]string, len(m1))
	for k, v := range m1 {
		m2[k] = v
	}

	// sort m2 by key
	keys := make([]string, 0, len(m2))
	for k := range m2 {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for _, k := range keys {
		fmt.Printf("[%s,%s]\n", k, m2[k])
	}
	
	fmt.Println("----")

	// sort m2 by value
	type kv struct {
		key   string
		value string
	}

	var kvs []kv
	for k, v := range m2 {
		kvs = append(kvs, kv{k, v})
	}
	sort.Slice(kvs, func(i, j int) bool {
		return kvs[i].value < kvs[j].value
	})
	for _, kv := range kvs {
		fmt.Printf("[%s,%s]\n", kv.key, kv.value)
	}

}