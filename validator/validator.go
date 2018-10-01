package validator

import "regexp"

var romanNumberRegex = "^M{0,4}(CM|CD|D?C{0,3})(XC|XL|L?X{0,3})(IX|IV|V?I{0,3})$"

func IsValid(currency string) bool {
	cur := []byte(currency)
	ret, err := regexp.Match(romanNumberRegex, cur)
	if err != nil {
		return false
	}
	return ret
}
