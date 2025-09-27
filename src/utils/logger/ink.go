package logger

import "fmt"

const (
	reset         = "\033[0m"
	redColour     = "\033[31m"
	greenColour   = "\033[32m"
	yellowColour  = "\033[33m"
	blueColour    = "\033[34m"
	magentaColour = "\033[35m"
	cyanColour    = "\033[36m"
)

func paint(colour string) func(string) string {
	return func(s string) string {
		return fmt.Sprintf("%s%s%s", colour, s, reset)
	}
}

func Green(str string) string {
	return paint(greenColour)(str)
}

func Red(str string) string {
	return paint(redColour)(str)
}

func Yellow(str string) string {
	return paint(yellowColour)(str)
}

func Blue(str string) string {
	return paint(blueColour)(str)
}

func Magenta(str string) string {
	return paint(magentaColour)(str)
}

func Cyan(str string) string {
	return paint(cyanColour)(str)
}
