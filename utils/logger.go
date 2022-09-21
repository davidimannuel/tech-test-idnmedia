package utils

import (
	"context"
	"fmt"
	"log"
	"runtime"
)

func FnTrace() string {
	pc := make([]uintptr, 10) // at least 1 entry needed
	runtime.Callers(2, pc)
	f := runtime.FuncForPC(pc[0])
	file, line := f.FileLine(pc[0])
	return fmt.Sprintf("%s:%d %s", file, line, f.Name())
}

func LogInfo(ctx context.Context, logTrace, msg string, args ...interface{}) {
	log.Println("[INFO]", fmt.Sprintf(msg, args...), logTrace)
}

func LogError(ctx context.Context, logTrace, msg string, args ...interface{}) {
	log.Println("[ERROR]", fmt.Sprintf(msg, args...), logTrace)
}
