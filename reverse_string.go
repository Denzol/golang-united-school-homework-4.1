package reverse_string

import "strings"

func ReverseString(input string) (output string) {
	slice := strings.Split(input, "")
	for i := 0; i < len(slice)/2; i++ {
		c := slice[i]
		slice[i] = slice[len(slice)-1-i]
		slice[len(slice)-1-i] = c
	}
	output = strings.Join(slice, "")
	return output
}
