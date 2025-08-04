package logger

import (
	"fmt"
	"time"
)

type printer struct {
}

type printMessage struct {
	colour byte
	text   *string
}

func getMoment(t time.Time) string {
	return t.UTC().Format("2006-01-02T15:04:05.000Z")
}

func (p *printer) init() chan<- printMessage {
	colorMap := map[byte]func(string) string{
		green:   Green,
		red:     Red,
		yellow:  Yellow,
		blue:    Blue,
		cyan:    Cyan,
		magenta: Magenta,
	}
	ch := make(chan printMessage, 1_000)
	go func() {
		defer close(ch)
		for msg := range ch {
			toLog := fmt.Sprintf("%s: %s", getMoment(time.Now()), *msg.text)
			painter, ok := colorMap[msg.colour]
			if ok {
				toLog = painter(toLog)
			}
			fmt.Println(toLog)
		}
	}()
	return ch
}
