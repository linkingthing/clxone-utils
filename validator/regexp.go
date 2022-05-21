package validator

import (
	"fmt"
	"regexp"
)

type StringRegexp struct {
	Regexp       *regexp.Regexp
	ErrMsg       string
	ExpectResult bool
}

var StringRegexps = []*StringRegexp{
	{
		Regexp:       regexp.MustCompile(`^[0-9a-zA-Z-_/,，、\p{Han}@*]+$`),
		ErrMsg:       "is not legal",
		ExpectResult: true,
	},
	{
		Regexp:       regexp.MustCompile(`(^-)|(^\.)|(^/)|(^,)|(^，)|(^、)`),
		ErrMsg:       "is not legal",
		ExpectResult: false,
	},
	{
		Regexp:       regexp.MustCompile(`-$|_$|/$|，$|、$|,$|\.$`),
		ErrMsg:       "is not legal",
		ExpectResult: false,
	},
}

func ValidateString(s string) error {
	if s != "" {
		for _, reg := range StringRegexps {
			if ret := reg.Regexp.MatchString(s); ret != reg.ExpectResult {
				return fmt.Errorf("%s %s", s, reg.ErrMsg)
			}
		}
	}

	return nil
}
