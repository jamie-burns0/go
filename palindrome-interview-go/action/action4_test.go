package action_test

import (
	"testing"

	"jamieburns.me/palidrome-interview/action"
)

func TestANilOrEmptyMapReturnsAnEmptyString( t *testing.T ) {

	var m map[string]int

	p := action.MakePalindromeFromFreqCharacterMap( m )

	if len( p ) != 0 {
		t.Errorf( "Expected an empty string for an empty map. Got %v", p )
	}

	m2 := map[string]int {}

	p2 := action.MakePalindromeFromFreqCharacterMap( m2 )

	if len( p2 ) != 0 {
		t.Errorf( "Expected an empty string for an empty map. Got %v", p2 )
	}
}

func TestAValidMapReturnsAPalindrome( t *testing.T ) {

	m := map[string]int {
		"a": 1,
	}

	p := action.MakePalindromeFromFreqCharacterMap( m )

	if p != "a" {
		t.Errorf( "Expected palindrome 'a'. Got %v", p )
	}

	m2 := map[string]int {
		"a": 2,
		"b": 1,
	}

	p2 := action.MakePalindromeFromFreqCharacterMap( m2 )

	if p2 != "aba" {
		t.Errorf( "Expected palindrome 'aba'. Got %v", p2 )
	}

	m3 := map[string]int {
		"a": 2,
		"b": 1,
		"c": 4,
	}

	p3 := action.MakePalindromeFromFreqCharacterMap( m3 )

	if p3 != "accbcca" {
		t.Errorf( "Expected palindrome 'accbcca'. Got %v", p3 )
	}
}

func TestAMapWithAccentedLatinCharactersReturnsAPalindromeWithTheAccentedCharacters( t *testing.T ) {

	m := map[string]int {
		"á": 1,
		"â": 2,
		"ç": 2,
	}

	p := action.MakePalindromeFromFreqCharacterMap( m )

	if p != "âçáçâ" {
		t.Errorf( "Expected palindrome 'âçáçâ'. Got %v", p )
	}
}