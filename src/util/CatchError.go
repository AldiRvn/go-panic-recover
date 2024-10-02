package util

import (
	"fmt"
	"path/filepath"
	"runtime"
	"time"

	"github.com/logrusorgru/aurora"
)

func CatchError() {
	if err := recover(); err != nil {
		_, file, no, ok := runtime.Caller(3)
		if ok {
			file = fmt.Sprintf("%s:%d", file, no)
			fmt.Printf("[%s] %s in file %s or %s with error:\n\t%v\n",
				time.Now().Format(time.DateTime), aurora.Red("PANIC !"), filepath.Base(file), file,
				aurora.Red(err))
		}
	}
}
