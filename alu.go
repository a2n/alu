package alu

import (
    "log"
    "runtime"
)

func Caller() string {
    pc, _, _, ok := runtime.Caller(1)
    if !ok {
	    return "Unknown"
    } else {
	    return runtime.FuncForPC(pc).Name()
    }
}

func InitLog() {
    log.SetFlags(log.Lshortfile | log.LstdFlags)
}

func ResetLog() {
    log.SetFlags(log.LstdFlags)
}
