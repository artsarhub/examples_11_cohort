package functions

import "errors"

func ApplyOperation(a, b int, f func(a, b int) int) int {
	fuzz := func() int {
		return a * b
	}
	return fuzz()
}

func IsEqual(a, b int) float64 {
	var res float64 = float64(a) / float64(b)
	return res
}

func ValidateLogin(v string) error {
	if v == "" {
		return errors.New("invalid login")
	}
	return nil
}

func ValidatePassword(v string) error {
	if len(v) <= 8 {
		return errors.New("invalid password")
	}
	return nil
}

func Validate(v string, validator func(string) error) error {
	return validator(v)
}
