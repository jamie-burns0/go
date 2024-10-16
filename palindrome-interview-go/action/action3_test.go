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

func TestAMapWithAccentedLatinCharactersIsAPalindromeCandidate( t *testing.T ) {

	m := map[string]int {
		"à": 2,
		"á": 2,
		"â": 2,
		"ã": 2,
		"ä": 2,
		"å": 2,
		"ç": 2,
		"è": 2,
		"é": 2,
		"ê": 2,
		"ë": 2,
		"ì": 2,
		"í": 2,
		"î": 2,
		"ï": 2,
		"ñ": 2,
		"ò": 2,
		"ó": 2,
		"ô": 2,
		"õ": 2,
		"ö": 2,
		"ù": 2,
		"ú": 2,
		"û": 2,
		"ü": 2,
		"ý": 2,
		"ÿ": 2,
	}

	if ! action.IsFreqCharacterMapAPalindromeCandidate( m ) {
		t.Errorf( "Expected true. Got false" )
	}
}