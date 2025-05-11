package debug

import "runtime"

// if everything is ok, returns file, line number, and function name
// otherwise returns "?", 0, "?"
// https://stackoverflow.com/questions/25927660/how-to-get-the-current-function-name
func trace() (string, int, string) {
	pc := make([]uintptr, 15)
	n := runtime.Callers(4, pc)
	frames := runtime.CallersFrames(pc[:n])
	frame, _ := frames.Next()
	return frame.File, frame.Line, frame.Function
}
