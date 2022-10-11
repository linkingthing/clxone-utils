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
		"中国>四川",
		"https://127.0.0.1:22",
		"https://[::1]:22",
		"https://www.baidu.com",
		"127.0.0.1",
		"你好(2001)",
		"你（好）吗",
		"你 好",
	}

	if err := ValidateStrings(ss...); err != nil {
		t.Error(err.Error())
	}
}
