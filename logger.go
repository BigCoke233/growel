package growel

import (
    "log"
)

type Logger struct{}

func (Logger) Info(msg string, v ...any) {
    log.Printf("[INFO] "+msg, v...)
}

func (Logger) Error(err error, msg string, v ...any) {
    log.Printf("[ERROR] "+msg+": %v", append(v, err)...)
}

var L = Logger{}
