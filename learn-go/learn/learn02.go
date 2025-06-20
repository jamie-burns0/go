package learn

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
	"time"
)

func Learn02() {

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("enter text: ")

	str, _ := reader.ReadString('\n')
	fmt.Println(str)

	fmt.Print("enter a number: ")
	str, _ = reader.ReadString('\n')

	fpn, err := strconv.ParseFloat(strings.TrimSpace(str), 64)

	if( err != nil ) {
		fmt.Println(err)
	} else {
		fmt.Println(fpn)
	}
}

func Learn02_02() {

	i1, i2, i3 := 1, 2, 3
	total := i1 + i2 + i3
	fmt.Println("Total: ", total)

	var fSum float64
	f1, f2, f3 := 1.1, 2.2, 3.3
	fSum = f1 + f2 + f3 + float64(i1)
	fmt.Println("Total: ", fSum)

	nSum := i1 + i2 + i3 + int(f1)

	fmt.Println("Total: ", nSum)
	fmt.Println("Total: ", int(fSum))
	
	nProduct := f1 * float64(i2)
	fmt.Println("Product: ", nProduct)
	fmt.Println("Product: ", int(nProduct))
}

func Learn02_03() {

	total := 1.23 + 2.34 * 3.65 / 4.33

	fmt.Println(total)
	fmt.Println(math.Round(total))

	circumference := 15.5 * 2 * math.Pi

	fmt.Printf("circumference: %.2f\n", circumference)
}

func Learn02_04() {

	t := time.Date(2000, time.November, 10, 23, 0, 0, 0, time.UTC)
	fmt.Printf("Go was launched: %s\n", t)

	t2 := time.Now()
	fmt.Printf("Current time: %s\n", t2)
	fmt.Printf("Current type: %T\n", t2)

	fmt.Printf(t2.Format(time.ANSIC)+"\n")
	fmt.Printf(t2.Format(time.RFC822Z)+"\n")

	tomorrow := t2.AddDate(0, 0, 1)
	fmt.Printf(tomorrow.Format(time.RFC822Z) + "\n")

	format := "Mon, 02 Jan 2006 15:04:05 -0700"
	fmt.Printf(tomorrow.Format(format) + "\n")

	format2 := "Mon 01 02 03 04 05 06 07 08 09 10 11 12 13 14 15"
	fmt.Printf(tomorrow.Format(format2) + "\n")
}