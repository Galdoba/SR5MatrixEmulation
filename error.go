package main

import (
	"fmt"
	"runtime"
)

func assert(ok bool, s string) {
	if !ok {
		_, file, line, _ := runtime.Caller(1)
		panic(fmt.Sprintf("\033[31m\n%s:%d\n[error] %s\033[0m", file, line, s))
	}
}
