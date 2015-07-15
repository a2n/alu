package alu

import (
    "log"
    "runtime"
    "os"
    "path/filepath"
    "time"
    "fmt"
)

func Caller() string {
    pc, _, _, ok := runtime.Caller(1)
    if !ok {
	    return "Unknown"
    } else {
	    return runtime.FuncForPC(pc).Name()
    }
}

func InitLog(name string) {
    log.SetFlags(log.LstdFlags | log.Lshortfile | log.Lmicroseconds)

	if len(name) == 0 {
		_, file := filepath.Split(os.Args[0])
		name = file + ".log"
	}

	f, err := os.OpenFile(name, os.O_APPEND | os.O_WRONLY | os.O_CREATE, 0600)
	if err != nil {
		log.Printf("%s has error, %s", Caller(), err.Error())
		return
	}

	log.SetOutput(f)
}

func ResetLog() {
	log.SetOutput(os.Stdout)
    log.SetFlags(log.LstdFlags)
}

func ToDateString(date *time.Time) string {
    return fmt.Sprintf("%d-%02d-%02d", date.Year(), date.Month(), date.Day())
}

func ToDateTimeString(date *time.Time) string {
    return fmt.Sprintf("%d-%02d-%02d %02d:%02d:%02d", date.Year(), date.Month(), date.Day(), date.Hour(), date.Minute(), date.Second())
}
