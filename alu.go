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

func NewLogger(name string) (*log.Logger) {
	if len(name) == 0 {
		_, file := filepath.Split(os.Args[0])
		name = file + ".log"
	}

	f, err := os.OpenFile(name, os.O_APPEND | os.O_WRONLY | os.O_CREATE, 0600)
	if err != nil {
		log.Panicf("%s has error, %s.\n", Caller(), err.Error())
	}

	return log.New(f, "", log.LstdFlags | log.Lshortfile | log.Lmicroseconds)
}

func ToDateString(date *time.Time) string {
	if date == nil {
		return ""
	}
    return fmt.Sprintf("%d-%02d-%02d", date.Year(), date.Month(), date.Day())
}

func ToDateTimeString(date *time.Time) string {
	if date == nil {
		return ""
	}
    return fmt.Sprintf("%d-%02d-%02d %02d:%02d:%02d", date.Year(), date.Month(), date.Day(), date.Hour(), date.Minute(), date.Second())
}

func EasyDate(d time.Time) (string) {
	twZone, err := time.LoadLocation("Asia/Taipei")
	if err != nil {
		errLog.Print(err.Error())
		return ""
	}

	now := time.Now().In(twZone)
	today := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, twZone)
	this_year := time.Date(now.Year(), time.January, 1, 0, 0, 0, 0, twZone)
	yesterday := today.Add(-24 * time.Hour)
	twoDaysBefore := yesterday.Add(-24 * time.Hour)

	str := fmt.Sprintf("%02d:%02d:%02d", d.Hour(), d.Minute(), d.Second())
	if d.After(yesterday) && d.Before(today) {
		str = "昨天 " + str
	} else if d.After(twoDaysBefore) && d.Before(yesterday) {
		str = "前天 " + str
	} else if d.After(this_year) && d.Before(twoDaysBefore) {
		str = fmt.Sprintf("%d 月 %d 日 ", d.Month(), d.Day()) + str
	} else if d.Before(this_year) {
		str = fmt.Sprintf("%02d-%02d-%02d ", d.Year() - 1911, d.Month(), d.Day()) + str
	}

	return str
}
