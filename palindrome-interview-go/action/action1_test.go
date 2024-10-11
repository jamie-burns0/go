package action_test

import (
	"testing"

	"jamieburns.me/palidrome-interview/action"
)

func TestIsAPalindromeCandidate(t *testing.T) {

	if action.IsStringAPalindromeCandidate("abcdef") != true {
		t.Error("expected true.")
	}

	if action.IsStringAPalindromeCandidate("ab4cdef") != false {
		t.Error("expected false.")
	}

	if action.IsStringAPalindromeCandidate("") != true {
		t.Error("expected false.")
	}
}
