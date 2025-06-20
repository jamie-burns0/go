package action_test

import (
	"testing"

	"me.jamieburns/go-demo-1/action"
)

func TestAction( t *testing.T) {
	if action.SaySomething() == action.SaySomethingElse() {
		t.Fatal( "expected SaySomething and SaySomethingElse to be different" )
	}
}