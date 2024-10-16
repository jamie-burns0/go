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

func TestStringWithAccentedLatinCharactersIsAPalindromeCandidate( t *testing.T) {

	if action.IsStringAPalindromeCandidate("ààááââããääåå") != true {
		t.Error("expected true.")
	}

	if action.IsStringAPalindromeCandidate("çç") != true {
		t.Error("expected true")
	}

	if action.IsStringAPalindromeCandidate("èèééêêëë") != true {
		t.Error("expected true")
	}

	if action.IsStringAPalindromeCandidate("ììííîîïï") != true {
		t.Error("expected true")
	}

	if action.IsStringAPalindromeCandidate("ññ") != true {
		t.Error("expected true")
	}

	if action.IsStringAPalindromeCandidate("òòóóôôõõöö") != true {
		t.Error("expected true")
	}

	if action.IsStringAPalindromeCandidate("ùùúúûûüü") != true {
		t.Error("expected true")
	}

	if action.IsStringAPalindromeCandidate("ýýÿÿ") != true {
		t.Error("expected true")
	}
}
