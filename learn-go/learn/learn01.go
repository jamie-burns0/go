package learn

import "fmt"

func Learn01() {

	const explicitConst int = 1
	const implicitConst = 1

	var explicitInt int = explicitConst
	var explicitString string = "explicitly hello"

	implicitInt := implicitConst
	implicitString := "implicitly hello"

	fmt.Println(explicitInt, explicitString, implicitInt, implicitString)
}