package logging

import (
	"log"
)

var logger *log.Logger = nil

func Init() {
    if logger == nil {
        logger = log.Default()
    }
}

func Log(v ...interface{}) {
    Init()
    logger.Println(v...)
}

func Fatal(v ...interface{}) {
    Init()
    logger.Fatalln(v...)
}
