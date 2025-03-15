package iteration

import "strings"

// func Repeat(character string, times int) string {
// 	var repeated string
// 	for i := 0; i < times; i++ {
// 		repeated += character
// 	}
// 	return repeated
// }

// more memory efficient with strings.Builder

func Repeat(character string, times int) string {
	var repeated strings.Builder
	for i := 0; i < times; i++ {
		repeated.WriteString(character)
	}
	return repeated.String()
}
