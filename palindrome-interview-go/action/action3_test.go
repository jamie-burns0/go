package action_test

import (
	"testing"

	"jamieburns.me/palidrome-interview/action"
)

func TestANilOrEmptyMapReturnsFalse( t *testing.T ) {

	m := map[string]int {}

	if action.IsFreqCharacterMapAPalindromeCandidate( m ) {
		t.Errorf( "Expected false. Got true" )
	}
}

func TestAnyMapWithOnlyOneEntryAndEvenFrequencyReturnsTrue( t *testing.T ) {

	m := map[string]int {
		"a": 2,
	}

	if ! action.IsFreqCharacterMapAPalindromeCandidate( m ) {
		t.Errorf( "Expected true. Got false" )
	}
}

func TestAnyMapWithOnlyOneEntryAndOddFrequencyReturnsTrue( t *testing.T ) {

	m := map[string]int {
		"a": 3,
	}

	if ! action.IsFreqCharacterMapAPalindromeCandidate( m ) {
		t.Errorf( "Expected true. Got false" )
	}
}

func TestAValidMapWithMultipleEntriesReturnsTrue( t *testing.T ) {

	m := map[string]int {
		"a": 2,
		"b": 4,
	}

	if ! action.IsFreqCharacterMapAPalindromeCandidate( m ) {
		t.Errorf( "Expected true. Got false" )
	}
}

func TestAnInvalidMapWithMultipleEntriesReturnsFalse( t *testing.T ) {

	m := map[string]int {
		"a": 2,
		"b": 1,
		"c": 3,
	}

	if action.IsFreqCharacterMapAPalindromeCandidate( m ) {
		t.Errorf( "Expected false. Got true" )
	}
}