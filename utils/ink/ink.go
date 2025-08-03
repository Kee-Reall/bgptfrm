package ink

import "fmt"

const (
	reset   = "\033[0m"
	red     = "\033[31m"
	green   = "\033[32m"
	yellow  = "\033[33m"
	blue    = "\033[34m"
	magenta = "\033[35m"
	cyan    = "\033[36m"
)

func Red(str string) string {
	return fmt.Sprintf("%s%s%s", red, str, reset)
}

func Green(str string) string {
	return fmt.Sprintf("%s%s%s", green, str, reset)
}

func Yellow(str string) string {
	return fmt.Sprintf("%s%s%s", yellow, str, reset)
}

func Blue(str string) string {
	return fmt.Sprintf("%s%s%s", blue, str, reset)
}

func Magenta(str string) string {
	return fmt.Sprintf("%s%s%s", magenta, str, reset)
}

func Cyan(str string) string {
	return fmt.Sprintf("%s%s%s", cyan, str, reset)
}
