package string

func Head(str string, n int) string {
	runes := []rune(str)
	if n == 0 {
		return ""
	} else if n < 0 {
		return Tail(str, -n)
	} else if n >= len(runes) {
		return str
	} else {
		return string(runes[:n])
	}
}

func Tail(str string, n int) string {
	runes := []rune(str)
	if n == 0 {
		return ""
	} else if n < 0 {
		return Head(str, -n)
	} else if n >= len(runes) {
		return str
	} else {
		return string(runes[n:])
	}
}
