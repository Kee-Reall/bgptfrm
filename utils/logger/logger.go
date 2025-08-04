package logger

type Logger struct {
	printer *printer
	logBus  chan<- printMessage
}

const (
	green   byte = 0
	red     byte = 1
	yellow  byte = 2
	blue    byte = 3
	cyan    byte = 4
	magenta byte = 5
)

var GlobalLogger *Logger = nil

func (l *Logger) provide(colour byte, text *string) {
	select {
	case l.logBus <- printMessage{colour, text}:
	default:
		return
	}
}

func Info(info string) {
	GetLogger().provide(green, &info)
}

func Error(str string) {
	GetLogger().provide(red, &str)
}

func ErrorException(err error) {
	Error(err.Error())
}

func Warn(str string) {
	GetLogger().provide(yellow, &str)
}

func WarnException(err error) {
	Warn(err.Error())
}

func Debug(str string) {
	GetLogger().provide(magenta, &str)
}

func DebugException(err error) {
	Debug(err.Error())
}

func Log(str string) {
	GetLogger().provide(cyan, &str)
}

func Trace(str string) {
	GetLogger().provide(blue, &str)
}

func GetLogger() *Logger {
	if GlobalLogger == nil {
		printer := &printer{}
		logBus := printer.init()
		GlobalLogger = &Logger{printer, logBus}
	}
	return GlobalLogger
}
