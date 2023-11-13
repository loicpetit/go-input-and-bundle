package log

import (
	"fmt"
)

func Debug(s ...any) {
	fmt.Println(s...)
}

func Error(s ...any) {
	fmt.Println(s...)
}
