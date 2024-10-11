package action_test

import (
	"testing"

	"jamieburns.me/palidrome-interview/action"
)

func TestAnEmptyStringReturnsAnEmptyMap( t *testing.T ) {

	if( len(action.FrequencyOfCharacterMap( "" )) != 0 ) {
		t.Error( "expected true." )
	}
}

func TestAStringWithOnlyOneOccuranceOfOneCharacterReturnsAMapWithOnlyOneEntry( t *testing.T ) {

	s := "a"

	m := action.FrequencyOfCharacterMap( s )

	if len(m) != 1 {
		t.Errorf( "Expected only one entry. Got %v", len(m) )
	}

	if v, ok := m["a"]; ok {
		if v != 1 {
			t.Errorf( "Expected frequency of 'a' to be 1. Got %d", v)
		}
	} else {
		t.Errorf( "Expected frequency map to contain 'a'")
	}
}

func TestAStringWithMultipleOccurancesOfOnlyOneCharacterReturnsAMapWithOnlyOneEntry( t *testing.T ) {

	s := "aa"

	m := action.FrequencyOfCharacterMap( s )

	if len(m) != 1 {
		t.Errorf( "Expected only one entry. Got %v", len(m) )
	}

	if v, ok := m["a"]; ok {
		if v != 2 {
			t.Errorf( "Expected frequency of 'a' to be 2. Got %d", v)
		}
	} else {
		t.Errorf( "Expected frequency map to contain 'a'")
	}
}

func TestAStringWithOneOccuranceEachOfMultipleCharacterReturnsAMapWithMultipleEntries( t *testing.T ) {

	s := "ab"

	m := action.FrequencyOfCharacterMap( s )

	if len(m) != 2 {
		t.Errorf( "Expected two entries. Got %v", len(m) )
	}

	if v, ok := m["a"]; ok {
		if v != 1 {
			t.Errorf( "Expected frequency of 'a' to be 1. Got %d", v)
		}
	} else {
		t.Errorf( "Expected frequency map to contain 'a'")
	}

	if v, ok := m["b"]; ok {
		if v != 1 {
			t.Errorf( "Expected frequency of 'b' to be 1. Got %d", v)
		}
	} else {
		t.Errorf( "Expected frequency map to contain 'b'")
	}
}

func TestAStringWithMultipleOccurancesEachOfMultipleCharactersReturnsAMapWithMultipleEntries( t *testing.T ) {

	s := "abaab"

	m := action.FrequencyOfCharacterMap( s )

	if len(m) != 2 {
		t.Errorf( "Expected two entries. Got %v", len(m) )
	}

	if v, ok := m["a"]; ok {
		if v != 3 {
			t.Errorf( "Expected frequency of 'a' to be 3. Got %d", v)
		}
	} else {
		t.Errorf( "Expected frequency map to contain 'a'")
	}

	if v, ok := m["b"]; ok {
		if v != 2 {
			t.Errorf( "Expected frequency of 'b' to be 2. Got %d", v)
		}
	} else {
		t.Errorf( "Expected frequency map to contain 'b'")
	}
}