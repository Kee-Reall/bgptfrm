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

func process(ch chan printMessage, cm map[byte]func(string) string) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println(Red("Recovered from panic:"), r)
		}
		process(ch, cm)
	}()
	for msg := range ch {
		toLog := fmt.Sprintf("%s: %s", getMoment(time.Now()), *msg.text)
		painter, ok := cm[msg.colour]
		if ok {
			toLog = painter(toLog)
		}
		fmt.Println(toLog)
	}
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
	go process(ch, colorMap)
	return ch
}
