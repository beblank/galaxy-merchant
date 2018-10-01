package converter

import "strings"

var (
	glob = "I"
	prok = "V"
	pish = "X"
	tegj = "L"
)

func ConvertGalaxy(currency string) string {
	var result []string
	list := strings.Split(currency, " ")
	list = delete_empty(list)
	if len(list) == 0 {
		return "No Input"
	}
	for _, str := range list {
		switch str {
		case "glob":
			result = append(result, glob)
			break
		case "prok":
			result = append(result, prok)
			break
		case "pish":
			result = append(result, pish)
			break
		case "tegj":
			result = append(result, tegj)
			break
		}
	}
	return strings.Join(result, "")
}

func delete_empty(s []string) []string {
	var r []string
	for _, str := range s {
		if str != "" {
			r = append(r, str)
		}
	}
	return r
}
