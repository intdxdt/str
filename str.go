package str

import "strings"

//Capitalize first letter in a word boundaries
func Capitalize(s string) string {
	return strings.Title(s)
}

//Repeat a given string
func Repeat(s string, count int) string {
	return strings.Repeat(s, count)
}

//Ljust justify string based on padding
func Ljust(s string, width int, pad string) string {
	if len(s) < width {
		return s + Repeat(pad, width - len(s))
	}
	return s
}

//Rjust justify string based on padding
func Rjust(s string, width int, pad string) string {
	if len(s) < width {
		return Repeat(pad, width - len(s)) + s
	}
	return s

}

//Center a string based on padding
func Center(s string, width int, pad string) string {
	if len(s) < width {
		length := width - len(s)
		remain := ""
		if length % 2 != 0 {
			remain = pad
		}
		var pads = Repeat(pad, length / 2)
		return pads + s + pads + remain
	}
	return s
}
