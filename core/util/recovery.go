package util

import (
	"fmt"
	"runtime/debug"
)

func Recovery() {
	if err := recover(); err != nil {
		fmt.Println(err)
		debug.PrintStack()
	}
}
