package validator

import (
	"testing"
)

func TestValid(t *testing.T) {
	ss := []string{
		"123",
		"asd123",
		"ASD",
		"a_a",
		"a-a",
		"a*a",
		"中文",
		"a@a",
		"a,a",
		"a，、/a",
		"",
	}

	if err := ValidateStrings(ss...); err != nil {
		t.Error(err.Error())
	}
}
