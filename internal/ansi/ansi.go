package ansi

const colorReset = "\033[0m"
const colorRed = "\033[31;1m"

func Error(s string) string {
	return colorRed + s + colorReset
}
