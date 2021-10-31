package checkerr

import (
	"runtime"
	"log"
)

func Checkerr(err error){
	if err != nil {
		_, file, line, ok := runtime.Caller(1)
		if ok {
			log.SetFlags(log.LstdFlags)
			log.Println(file, line, err.Error())
			log.SetFlags(log.LstdFlags | log.Llongfile)
		} else {
			log.SetFlags(log.LstdFlags | log.Llongfile)
			log.Println(file, line, err.Error())

		}
	}
}
