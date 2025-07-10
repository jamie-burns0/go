package learn

import (
	"fmt"
	"math/rand"
)

func Learn04_01() { 
	d1 := Dog{Name: "Fred", Age: 8}
	fmt.Println("structs", d1)
	fmt.Printf("Name: %s, Age: %d\n", d1.Name, d1.Age)
	fmt.Printf("%+v\n", d1)
	d1.Name = "Fred2"
	fmt.Printf("%+v\n", d1)
	fmt.Printf("%+v\n", d1)
}

func Learn04_02() {
	value := rand.Intn(100)

	if value < 50 {
		fmt.Println("less than 50")
	} else if value > 50 {
		fmt.Println("greater than 50")
	} else {
		fmt.Println("equal to 50")
	}

	mod := rand.Intn(9) + 1

	if remainder := value % mod; remainder < 5 {
		fmt.Println("remainder < 5, value=", value, " mod=", mod, " remainder=", remainder)
	} else {
		fmt.Println("remainder >= 5, value=", value, " mod=", mod, " remainder=", remainder)
	}

	remainder2 := value % mod

	fmt.Println("value=", value, " remainder2=", remainder2)
}

func Learn04_03() {
	value := rand.Intn(59) + 1

	switch {
		case value <= 20 :
			fmt.Println(value, "is 1-20")
		case value > 20 && value <= 40 :
			fmt.Println(value, "is 21-40")
		default:
			fmt.Println(value, "is 41-60")
	}

	mod := (rand.Intn(5) + 1)
	remainder := value % mod

	switch remainder == 0 {
		case true:
			fmt.Printf("with mod=%v, remainder is zero (%v)\n", mod, remainder)
		case false: 
			fmt.Printf("with mod=%v, remainder is non-zero (%v)\n", mod, remainder)
	}
}

func Learn04_04() {

	rainbow := []string{"red","yellow","pink","blue","orange","purple","green"}

	for i := 0; i < len(rainbow); i++ {
		fmt.Printf("%v ", rainbow[i])
	}

	fmt.Println()

	for i := range rainbow {
		fmt.Printf("%v ", rainbow[i])
	}
	fmt.Println()
	
	for _, c := range rainbow {
		fmt.Printf("%v ", c)
	}

label1:

	count := 50

	for count > 0 {
		if rand.Intn(99) < 50 {
			break
		}
		count--
	}

	fmt.Printf("looped with break %v times\n", 50 - count)
	
	count2 := 50
	
	for count2 > 0 {
		if rand.Intn(99) < 10 {
			goto label1
		}
		count2--
	}
	fmt.Printf("looped with goto %v times\n", 50 - count2)
}

type Dog struct {
	Name string
	Age int
}