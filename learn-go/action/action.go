package action

import "rsc.io/quote"

func SaySomething() string {
	return quote.Go()
}

func SaySomethingElse() string {
	return quote.Opt()
}