package log

import "fmt"

func Info(format string, a ...any) {
	format = "[Info] " + format
	fmt.Printf(format, a...)
}

func Warn(format string, a ...any) {
	format = "[Warn] " + format
	fmt.Printf(format, a...)
}

func Error(format string, a ...any) {
	format = "[Error] " + format
	fmt.Printf(format, a...)
}
