package app_errors

import (
	"os"
	"runtime"
	"strconv"
)

var stackTraceEnabled bool

func init() {
	stackTraceEnabled, _ = strconv.ParseBool(os.Getenv("STACK_TRACE_ENABLED"))
}

func newStackTrace() []byte {
	buf := make([]byte, 1024)
	for {
		n := runtime.Stack(buf, true)
		if n < len(buf) {
			break
		}
		buf = make([]byte, 2*len(buf))
	}
	return buf
}

func errorWithStack(err string, trace []byte) string {
	if stackTraceEnabled {
		return err + string(trace)
	}
	return err
}
