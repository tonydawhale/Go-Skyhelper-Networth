package utils

import "strings"

func TitleCast(str string) string {
	splitStr := strings.Split(strings.ReplaceAll(strings.ToLower(str), "_", " "), " ")
	for i := 0; i < len(splitStr); i++ {
		if splitStr[i] == "" {
			continue
		}
		splitStr[i] = strings.ToUpper(string(splitStr[i][0])) + splitStr[i][1:]
	}
	return strings.Join(splitStr, " ")
}
