package strings

import (
	"fmt"
	"regexp"
	"testing"
)

// TestRandNumCapitals tests RandNumCapitals function.
func TestRandNumCapitals(t *testing.T) {
	l := 20
	r := regexp.MustCompile(fmt.Sprintf("[A-Z0-9]{%d}", l))

	s := RandNumCapitals(l)
	if len(s) != l {
		t.Fatal("RandNumCapitals generated string length: ", len(s))
	}

	if !r.MatchString(s) {
		t.Errorf("Generated string %s contains characters out of capitals and numbers.")
	}
}

// TestRandString tests RandString function.
func TestRandString(t *testing.T) {
	l := 30
	r := regexp.MustCompile(fmt.Sprintf("[[:alnum:]]{%d}", l))

	s := RandString(l)
	if len(s) != l {
		t.Fatal("RandString generated string length: ", len(s))
	}

	if !r.MatchString(s) {
		t.Errorf("Generated string %s contains characters out of capitals and numbers.")
	}
}

// TestRandLetters tests RandLetters function.
func TestRandLetters(t *testing.T) {
	l := 30
	r := regexp.MustCompile(fmt.Sprintf("[a-zA-Z]{%d}", l))

	s := RandLetters(l)
	if len(s) != l {
		t.Fatal("RandLetters generated string length: ", len(s))
	}

	if !r.MatchString(s) {
		t.Errorf("RandLetters string %s contains characters out of letters.")
	}
}

// TestRandHex tests RandHex function.
func TestRandHex(t *testing.T) {
	l := 30
	r := regexp.MustCompile(fmt.Sprintf("[0-9A-F]{%d}", l))

	s := RandHex(l)
	if len(s) != l {
		t.Fatal("RandHex generated string length: ", len(s))
	}

	if !r.MatchString(s) {
		t.Errorf("RandHex string %s contains characters out of hex")
	}
}

// TestRandNumbers tests RandNumbers function.
func TestRandNumbers(t *testing.T) {
	l := 30
	r := regexp.MustCompile(fmt.Sprintf("[0-9]{%d}", l))

	s := RandNumbers(l)
	if len(s) != l {
		t.Fatal("RandNumbers generated string length: ", len(s))
	}

	if !r.MatchString(s) {
		t.Errorf("RandNumbers string %s contains characters out of numbers")
	}
}
