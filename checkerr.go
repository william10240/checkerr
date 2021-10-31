package checkerr

import (
	"log"
	"runtime"
)

func Checkerr(err error) {
	if err != nil {
		_, file, line, ok := runtime.Caller(1)
		if ok {
			log.Println(file, line, err.Error())
		} else {
			log.Println(err.Error())
		}
	}
}
