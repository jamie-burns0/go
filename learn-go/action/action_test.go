package action_test

import (
	"testing"

	"jamieburns.me/go-demo-1/action"
)

func TestAction( t *testing.T) {
	if action.SaySomething() == action.SaySomethingElse() {
		t.Fatal( "expected SaySomething and SaySomethingElse to be different" )
	}
}